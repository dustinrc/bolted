package event

import (
	"bitbucket.org/gdamore/mangos"
	"bitbucket.org/gdamore/mangos/protocol/sub"
	"bitbucket.org/gdamore/mangos/transport/ipc"
	"bitbucket.org/gdamore/mangos/transport/tcp"
	"github.com/dustinrc/bolted"
)

type CallbackFunc func(string)

type Listener interface {
	bolted.Configurer
	Listen()
}

type listenConnection struct {
	sock     mangos.Socket
	callback CallbackFunc
}

func NewListenConnection(url string, cb CallbackFunc) *listenConnection {
	var sock mangos.Socket
	var err error

	if sock, err = sub.NewSocket(); err != nil {
		bolted.Die("Cannot initiate socket for listen: %s", err.Error())
	}

	sock.AddTransport(ipc.NewTransport())
	sock.AddTransport(tcp.NewTransport())

	if err = sock.Dial(url); err != nil {
		bolted.Die("Cannot dial on socket for listen: %s", err.Error())
	}

	return &listenConnection{sock, cb}
}

func (lc *listenConnection) Listen() {
	var msg []byte
	var err error

	for {
		if msg, err = lc.sock.Recv(); err != nil {
			bolted.Die("Cannot receive: %s", err.Error())
		}
		lc.callback(string(msg))
	}
}
