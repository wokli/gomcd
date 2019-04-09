package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/wokli/gomcd/pkg/storage"
)

func main() {
	st := storage.CreateStorage()
	log.Println("Storage initialized")
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	log.Println("Listening ", ":8080")

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn, st)
	}
}

func handle(conn net.Conn, st *storage.Storage) {

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cmd := scanner.Text()
		if cmd == "QUIT" {
			break
		}
		res, err := storage.ProcessCommand(st, cmd)

		if err != nil {
			fmt.Fprintln(conn, "ERR", err.Error())
		} else {
			fmt.Fprintln(conn, res)
		}
	}
}
