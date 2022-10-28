package main

import (
	"fmt"

	"api-go/server"
)

func statusServe(status string, sta chan string) {
	fmt.Println("status: ", status)
	sta <- "el estatus del servidor es: "
}

func main() {

	canal := make(chan string)
	go statusServe("anda cachondisimo ", canal)
	go statusServe("ahi va el server ", canal)
	go statusServe("anda mas prendido que un horno", canal)
	statusMensaje := <-canal
	fmt.Println(statusMensaje)

	srv := server.New(":8080")

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
