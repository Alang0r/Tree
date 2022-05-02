//TODO:
/*Логи
Натс
Подключение к БД
*/

package lib

import (
	"io/ioutil"

	"Tree/lib/log"

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

	//TODO: 3)подписываемся на очередь запросов и слушаем

	srv.Log.Infof("Тест лога обернутого")

}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

type Properties struct {
	Name string
	DB   *gorm.DB
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
	srv.Log.Info(cfg)

	//подключаемся к БД

	connectionString := "host="
	connectionString += cfg.Gorm_config.Host +
		" user=" + cfg.Gorm_config.User +
		" password=" + cfg.Gorm_config.Password +
		" dbname=" + cfg.Gorm_config.DBName +
		" port=" + cfg.Gorm_config.Port
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		srv.Log.Fatal(err)
	}
	DB.Debug()
}

type ServiceConfig struct {
	Gorm_config GormConfig
}
type GormConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}
