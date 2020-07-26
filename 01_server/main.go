package main

import (
	"log"
	"net/http"
)

func main(){


}

func Server(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, err := NewHandler(w, r)
		if err != nil {
			log.Println(err)
		}
		if err = ws.Handshake(); err != nil {
			log.Println(err)
		}
	}