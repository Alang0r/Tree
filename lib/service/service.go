//TODO:
/*Логи
Натс
Подключение к БД
*/

package lib

import (
	// "fmt"
	// "net/http"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	Properties
}

func (srv *Service) SetName(name string) {
	srv.Name = name
}

// func (srv *Service) Listen() {
// 	http.HandleFunc("/", handler) // each request calls handler
// 	//log.Fatal()
// 	// logger, _ := zap.NewProduction()
// 	// defer logger.Sync()
// 	// str := logger.Fatal(http.ListenAndServe("localhost:8000", nil))
// }
func (srv *Service) Start() {
	log, _ := zap.NewProduction()
	logger := log.Sugar()
	defer logger.Sync()
//	logger.
	logger.Infow("Start of", srv.Name)
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

type Properties struct {
	Name string
	DB   *gorm.DB
}
