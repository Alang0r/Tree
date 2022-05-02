//TODO:
/*Логи
Натс
Подключение к БД
*/

package lib

import (
	"fmt"
	"log"
	"net/http"
)

type Service struct {
	Settings ServiceSettings
}

func (srv *Service) SetName(name string) {
	srv.Settings.Name = name
}

func (srv *Service) Listen() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func (srv *Service) Start() {
	fmt.Println(srv.Settings.Name)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

type ServiceSettings struct {
	Name string

}
