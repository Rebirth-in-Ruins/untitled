// Package main implements a bot that can be controlled manually
package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/rebirth-in-ruins/torpedodge/server/game"
)

const (
	urlStr = "ws://localhost:8080/play"
	// urlStr = "wss://gameserver.resamvi.io/play"
)

func main() {
	conn, _, err := websocket.Dial(context.Background(), urlStr, nil)
	if err != nil {
		panic(err)
	}
	defer conn.CloseNow()

	// Join as player
	err = wsjson.Write(context.Background(), conn, "JOIN JulienBot.go")
	if err != nil {
		panic(err)
	}

	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)

	for {
		// RECEIVE NEXT STATE
		var state game.GameStateResponse
		err = wsjson.Read(context.Background(), conn, &state)
		if err != nil {
			panic(err)
		}

		// Wait for input
		os.Stdin.Read(b)

		action := ""
		switch string(b) {
		case "w":
			action = "UP"
		case "s":
			action = "DOWN"
		case "a":
			action = "LEFT"
		case "d":
			action = "RIGHT"
		case " ":
			action = "BOMB"
		case "q":
			action = "LASER"
		case "j":
			action = "BOMBDOWN"
		case "k":
			action = "BOMBUP"
		case "h":
			action = "BOMBLEFT"
		case "l":
			action = "BOMBRIGHT"
		}
		fmt.Println(action)

		// SEND ACTION
		err = wsjson.Write(context.Background(), conn, action)
		if err != nil {
			panic(err)
		}
	}
}
