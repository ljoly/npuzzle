package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func initServer() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connect", func(socket socketio.Socket) {
		log.Println("CONNECTED")
		socket.On("hello from front", func(msg string) {
			log.Println("HELLO FROM FRONT !")
		})
		socket.On("disconnect", func() {
			log.Println("disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
