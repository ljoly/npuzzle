package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

// func handler(w http.ResponseWriter, req *http.Request) {
// 	// ...
// 	enableCors(&w)
// 	// ...
// }

func launchServer(e Env) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	play(&e)
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

	// http.HandleFunc("/socket.io/", handler)

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
