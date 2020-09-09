package main

import (
	server "Receiptify/server"
)

func main() {
	always := make(chan int)
	go server.Router()
	<-always
}
