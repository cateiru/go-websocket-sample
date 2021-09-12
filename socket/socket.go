package socket

import (
	"fmt"
	"io"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

var datas = make(map[string]*string)

func Socket(ws *websocket.Conn) {
	defer logrus.Info("Close socket!")

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
	datas["hoge"] = new(string)
	data := ""

	for {
		select {
		case <- quit:
			return
		default:
			if data != *datas["hoge"] {
				if err := websocket.Message.Send(ws, *datas["hoge"]); err != nil {
					logrus.Errorf("send err: %v", err)
					return
				}
				logrus.Info("send!")

				data = *datas["hoge"]
			}
			time.Sleep(3 * time.Second)
		}
	}
}


func Runner() {
	var index = 0

	for {
		for id, value := range datas {
			if id == "hoge" {
				*value = fmt.Sprint(index)
			}
		}
		time.Sleep(3 * time.Second)
		index++
	}
}
