//TODO:
/*Логи
Натс
Подключение к БД
*/

package lib

import (
	"io/ioutil"

	"Tree/lib/log"

	"github.com/streadway/amqp"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	Properties
	Log log.Logger
}

func (srv *Service) SetName(name string) {
	srv.Name = name
}

func (srv *Service) Start() {

	srv.Log.Infof("Start of %s", srv.Name)
	defer srv.Log.Infof("Finish of %s", srv.Name)

	_, err := srv.RabbitChannel.QueueDeclare(
		srv.Name+"-queue",
		false,
		false,
		false,
		false,
		nil,
	)
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
		srv.Log.Fatal(err)
	}

	forever := make(chan bool)
	go func() {
		for req := range msgs {

			//обработчик сообщений, котоырй вызывает соответствующий запрос из апи
			srv.Log.Info(req)
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
	srv.Log.Init(srv.Name)

	//загружаем настройки
	yfile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		srv.Log.Fatal(err)
	}

	cfg := ServiceConfig{}
	err = yaml.Unmarshal(yfile, &cfg)
	if err != nil {
		srv.Log.Fatalf("error: %v", err)
	}

	//подключаемся к БД

	postgresConfig := "host="
	postgresConfig += cfg.Gorm_config.Host +
		" user=" + cfg.Gorm_config.User +
		" password=" + cfg.Gorm_config.Password +
		" dbname=" + cfg.Gorm_config.DB +
		" port=" + cfg.Gorm_config.Port
	DB, err := gorm.Open(postgres.Open(postgresConfig), &gorm.Config{})
	if err != nil {
		srv.Log.Fatal(err)
	}

	DB.Debug() //FIXME

	//Создаем очередь и подписываемся для просулшивания сообщений
	rabbitConfig := "amqp://" +
		cfg.Rabbit_config.User +
		":" + cfg.Rabbit_config.Password +
		"@" + cfg.Rabbit_config.Host +
		":" + cfg.Rabbit_config.Port + "/"

	con, err := amqp.Dial(rabbitConfig)
	if err != nil {
		srv.Log.Fatal(err)
	}

	//defer con.Close()

	srv.RabbitChannel, err = con.Channel()
	if err != nil {
		srv.Log.Fatal(err)
	}
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
