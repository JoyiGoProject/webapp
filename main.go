package main

import (
	"fmt"
	"net/http"
	"os"
	"tlog"
	"webapp/router"
)

func init() {
	router.Init()
}

func main() {
	initLog()
	tlog.Info("Https server Running on http://:80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("app start error", err)
		return
	}
}

func initLog() {
	if _, err := os.Stat("log"); err != nil {
		os.Mkdir("log", 0755)
	}
	tlog.SetLogger("file", `{"filename":"log/webapp.log"}`)
}
