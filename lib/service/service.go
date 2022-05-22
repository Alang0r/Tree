//TODO: когда запускается сервис, в нем создается мапа 
//map["RequestName"] RequestType со всеми типами запросов
//когда приходит запрос, анмаршелим его в reqmap[request], т.к. запрос у нас в хедере
//if !ok reqmap["request"] return error запроса нет
//иначе дергаем из запроса execute

package lib

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"Tree/lib/log"
	"Tree/lib/request"

	"github.com/streadway/amqp"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	Properties
	Log log.Logger
}
const(
	tries = 3
	delay = 5
)

func (srv *Service) SetName(name string) {
	srv.Name = name
}



func (srv *Service) Start() {

		srv.Log.Infof("Start of %s", srv.Name)
		defer srv.Log.Infof("Finish of %s", srv.Name)

		//создаем очередь, если не было ранее
		_, err := srv.RabbitChannel.QueueDeclare(
			srv.Name+"-queue",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			srv.Log.Fatal(err.Error())
		}

		//Читаем сообщение и вызываем соответствующий запрос / возвращаем ошибку, если запрсоа нет
		msgs, err := srv.RabbitChannel.Consume(
			srv.Name+"-queue",
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			srv.Log.Fatal(err.Error())
		}

		forever := make(chan bool)
		go func() {
			for req := range msgs {

				//обработчик сообщений, который вызывает соответствующий запрос из апи
				srv.Log.Info("New request: " + string(req.Body))
				var data map[string]interface{}
				_ = json.Unmarshal(req.Body, &data)
			}
		}()
		<-forever
}

type Properties struct {
	Name          string
	DB            *gorm.DB //TODO: move out
	RabbitChannel *amqp.Channel
}

func (srv *Service) Configure() {
	//Инициализируем логи
	srv.Log.Init()

	srv.Log.Info("Configuration...")
	defer srv.Log.Info("Configuration complete.")
	//загружаем настройки
	srv.Log.Info("Loading config...")
	yfile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		srv.Log.Fatal(err.Error())
	}
	srv.Log.Info("Config loaded")

	cfg := ServiceConfig{}
	err = yaml.Unmarshal(yfile, &cfg)
	if err != nil {
		srv.Log.Fatalf("error: %v", err)
	}

	//подключаемся к БД
	srv.Log.Info("Connecting Postgres...")
	postgresConfig := "host="
	postgresConfig += cfg.Gorm_config.Host +
		" user=" + cfg.Gorm_config.User +
		" password=" + cfg.Gorm_config.Password +
		" dbname=" + cfg.Gorm_config.DB +
		" port=" + cfg.Gorm_config.Port
	DB, err := gorm.Open(postgres.Open(postgresConfig), &gorm.Config{})
	if err != nil {
		srv.Log.Fatal(err.Error())
	}
	srv.Log.Info("Postgres connected.")
	//FIXME: проверить, нужно ли джелать миграции

	DB.Debug() //FIXME

	//Создаем очередь и подписываемся для просулшивания сообщений
	srv.Log.Info("Connecting RabbitMQ...")
	rabbitConfig := "amqp://" +
		cfg.Rabbit_config.User +
		":" + cfg.Rabbit_config.Password +
		"@" + cfg.Rabbit_config.Host +
		":" + cfg.Rabbit_config.Port + "/"

	//пытаемся подключиться 3 раза с паузой 5 секунд
	for i:=0; i < tries ; i++ {
		con, err := amqp.Dial(rabbitConfig)
		
		if err != nil {
			if err == amqp.ErrClosed {
				time.Sleep(delay * time.Second)
				continue
			} else {
				srv.Log.Fatal(err.Error())
			}
		}

		srv.RabbitChannel, err = con.Channel()
		if err != nil {
			if err == amqp.ErrClosed {
				time.Sleep(delay * time.Second)
				continue
			} else {
				srv.Log.Fatal(err.Error())
			}
		}
	}
	srv.Log.Info("RabbitMQ connected")
	//defer srv.RabbitChannel.Close()

}

type ServiceConfig struct {
	Gorm_config   GormConfig
	Rabbit_config RabbitConfig
}
type GormConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type RabbitConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (srv *Service) Serve(req request.Request) {
	req.Execute()
}

