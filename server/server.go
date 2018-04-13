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

func raw(data Move) []byte {
	rawMove, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		printError("error json")
		return nil
	}
	return rawMove
}

func launchServer() {
	var index int

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	moves = make([]Move, 0)

	play()

	server.On("connection", func(socket socketio.Socket) {
		index = 0
		log.Println("CONNECTED")

		rawMove := raw(moves[index])
		// emit the first State on connection
		socket.Emit("state", rawMove)

		socket.On("prevState", func() {
			fmt.Println("PREV")
			if index > 0 {
				index--
				rawMove := raw(moves[index])
				socket.Emit("state", rawMove)
			}
		})

		socket.On("nextState", func() {
			fmt.Println("NEXT")
			if index < len(moves) {
				index++
				rawMove := raw(moves[index])
				socket.Emit("state", rawMove)
			}
		})

		socket.On("go", func() {
			fmt.Println("GO")
			index = 0
			for index < len(moves) {
				rawMove := raw(moves[index])
				socket.Emit("state", rawMove)
				index++
			}
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
	log.Println("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
