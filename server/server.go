package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

type Move struct {
	Board     []int
	Priority  int
	Heuristic int
}

var moves []Move

func getMoves(state *State) {
	if state != nil {
		getMoves(state.parent)
		m := Move{
			Board:     state.board,
			Priority:  state.priority,
			Heuristic: state.heuristic,
		}
		moves = append(moves, m)
		printState(state)
		e.moves++
	}
}

func launchServer() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	moves = make([]Move, 0)

	play()

	i := 0
	server.On("connection", func(socket socketio.Socket) {
		log.Println("CONNECTED")

		rawMove, err := json.Marshal(moves[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		// emit the first State on connection
		socket.Emit("nextState", rawMove)

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
	log.Println("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
