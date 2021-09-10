package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/yuto51942/go-websocket-sample/socket"
	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", websocket.Handler(socket.Socket))
	logrus.Info("start server to :8080")
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
