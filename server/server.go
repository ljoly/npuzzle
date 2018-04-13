package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

var (
	server *socketio.Server
	err    error
)

func launchServer() {
	server, err = socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	play()
	server.On("connection", func(socket socketio.Socket) {
		log.Println("CONNECTED")
		socket.On("hello", func(msg string) {
			log.Println("HELLO FROM FRONT !")
		})
		socket.On("disconnection", func() {
			log.Println("DISCONNECTED")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		server.ServeHTTP(w, r)
	})

	// http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
