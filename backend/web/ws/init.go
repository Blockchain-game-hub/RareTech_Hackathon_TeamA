package ws

import (
	"fmt"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

func Connect(room string) (server *socketio.Server) {
	server = socketio.NewServer(nil)

	_, err := server.Adapter(&socketio.RedisAdapterOptions{
		Host:   "127.0.0.1",
		Port:   "6379",
		Prefix: "socket.io",
	})

	if err != nil {
		log.Fatal("error:", err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Join(room)
		return nil
	})

	server.OnEvent("/", "draw", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("res_draw", msg)
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	return server
}
