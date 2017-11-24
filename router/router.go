package router

import "net/http"

func Init() {
	router()
}

func router() {
	http.HandleFunc("/", controller)

}
