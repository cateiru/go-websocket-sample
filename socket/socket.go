package socket

import (
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func Socket(ws *websocket.Conn) {
	go Send(ws)

	for {
		var response []byte
		if err := websocket.Message.Receive(ws, &response); err != nil {
			logrus.Errorf("receive err: %v", err)
			return
		}
		logrus.Infof("Response to %v", string(response))
	}
}

func Send(ws *websocket.Conn) {
	sendData := "hoge"
	for {
		if err := websocket.Message.Send(ws, sendData); err != nil {
			logrus.Errorf("send err: %v", err)
			return
		}
		logrus.Info("send!")
		time.Sleep(3 * time.Second)
	}
}
