package socket

import (
	"io"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func Socket(ws *websocket.Conn) {
	quit := make(chan bool)
	go Send(ws, quit)

	for {
		var response []byte
		if err := websocket.Message.Receive(ws, &response); err != nil {
			if err == io.EOF {
				quit <- true
				logrus.Info("Close...")
			}else {
				logrus.Errorf("receive err: %v", err)
			}
			return
		}
		logrus.Infof("Response to %v", string(response))
	}
}

func Send(ws *websocket.Conn, quit chan bool) {
	sendData := "hoge"
	for {
		select {
		case <- quit:
			return
		default:
			if err := websocket.Message.Send(ws, sendData); err != nil {
				logrus.Errorf("send err: %v", err)
				return
			}
			logrus.Info("send!")
			time.Sleep(3 * time.Second)
		}
	}
}
