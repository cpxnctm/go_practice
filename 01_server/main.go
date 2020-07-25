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
		lg, err := os.Create("logging.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer lg.Close()
		w1, err := lg.WriteString(ln)
		if err != nil {
			fmt.Println(err)
		fmt.Printf("%b", w1)

		}



		fmt.Println(ln)
		fmt.Fprintf(conn, "You sent: %s\n", ln)

	}
	defer conn.Close()
}
