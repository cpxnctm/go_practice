package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("Websocket Open...")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the ws Homepage!")
}

//check incoming origin
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	//upgrade the http connection to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected Successfully!")
	reader(ws)
}

func reader(conn *websocket.Conn) {
	//create the logfile
	f, err := os.OpenFile("logging.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))
		//output data to the logfile
		lg := log.New(f, "logging:", log.LstdFlags)
		lg.Printf("%s\n",p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}

}
