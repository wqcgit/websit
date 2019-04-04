package main

import (
	"net/http"
	"time"
	"log"
	"io"
	"privMux"
)

func tmpfunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "please login")
}

func main() {
	svr := http.Server{
		Addr: ":8080",
		Handler: &privMux.PrivHandler{},
		ReadTimeout: 5*time.Second,
	}
	privMux.InitUrlMux()
	privMux.SetUrlHandler("/login", tmpfunc)
	err := svr.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}



   