package main

import (
	"bufio"
	"fmt"
	console "github.com/polis-mail-ru-golang-1/t2-invert-index-search-vadimrebrin/console"
	index "github.com/polis-mail-ru-golang-1/t2-invert-index-search-vadimrebrin/index"
	"net"
	"os"
	"strings"
)

var dict map[string]map[string]int

func main() {
	dict = make(map[string]map[string]int)
	files := os.Args[1:]

	index.BuildIndex(dict, console.ReadFiles(files))
	listener, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	name := conn.RemoteAddr().String()
	fmt.Println(name + " connected")
	conn.Write([]byte("Enter phrase or 'exit' to exit\n\r"))
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		text := scan.Text()
		text = strings.ToLower(text)
		if text == "exit" {
			fmt.Println(name + " disconnected")
			break
		}
		fmt.Println(name + " entered " + text)
		phrase := strings.Fields(text)
		res := index.FindPhrase(dict, phrase)
		conn.Write([]byte(console.PrintInfo(res)))
	}
}
