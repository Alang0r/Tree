//TODO:
/*Логи
Натс
Подключение к БД
*/

package lib

import (
	// "fmt"
	// "net/http"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	l2 "Tree/lib/log"
)

type Service struct {
	Properties
}

func (srv *Service) SetName(name string) {
	srv.Name = name
}

func (srv *Service) Start() {
	
	//TODO: 3)подписываемся на очередь запросов и слушаем
	logger := l2.Logger{}
	logger.Init()
	logger.Infof("Тест лога обернутого")


}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

type Properties struct {
	Name string
	DB   *gorm.DB
}

func( srv *Service) Configure() {
	//TODO: 1)подтягиваем конфиг из файла
	//TODO: 2)настраиваем всё необходимое- логи, подключение к БД, ...
	//config := ServiceConfig{}

	 yfile, err := ioutil.ReadFile("config.yaml")
     if err != nil {
          log.Fatal(err)
     }
     data := make(map[interface{}]interface{})
     err2 := yaml.Unmarshal(yfile, &data)
     if err2 != nil {
          log.Fatal(err2)
     }
     for k, v := range data {
          fmt.Printf("%s -> %d\n", k, v)
     }
	 dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

	}
	DB.Debug()
}



type ServiceConfig struct {
	GormConfig
}
type GormConfig struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
	User string `ini:"user"`
	Password string `ini:"password"`
	DBName string `ini:"dbname"`
	Sslmode bool `ini:"sslmode"`
	TimeZone string `ini:"timezone"`
}
