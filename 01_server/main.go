package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main(){
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)

	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn){
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		ct:= fmt.Fprintf(conn, "You sent: %s\n", ln)

		os.Create("logs.txt")
		lg, err := os.OpenFile("logs.txt",os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
		}
		defer lg.Close()
		log.SetOutput(lg)


	}
	defer conn.Close()
}
