package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	port string
)

func handleCtrlC(c chan os.Signal) {
	sig := <-c
	log.Fatal("fatal: recivied ", sig,"\ninfo: inducing graceful shutdown\ndone!")
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go handleCtrlC(c)
	log.Println("info: init")
	dir := os.Args[1]
	log.Println("info: serving from:", dir)
	http.Handle("/", http.FileServer(http.Dir(string(dir))))
	if os.Args[2] == "" {
		port = ":5000"
	} else {
		port = ":" + os.Args[2]
	}
	log.Println("info: running on port:", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
