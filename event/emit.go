package event

import (
	"bitbucket.org/gdamore/mangos"
	"bitbucket.org/gdamore/mangos/protocol/pub"
	"bitbucket.org/gdamore/mangos/transport/ipc"
	"bitbucket.org/gdamore/mangos/transport/tcp"
	"github.com/dustinrc/bolted"
)

type Emitter interface {
	Emit(string)
}

type emitConnection struct {
	sock mangos.Socket
	cfg  bolted.Configurer
}

func NewEmitConnection(cfg bolted.Configurer) *emitConnection {
	var sock mangos.Socket
	var err error
	var url string
	var ok bool

	if sock, err = pub.NewSocket(); err != nil {
		bolted.Die("Cannot initiate socket for emit: %s", err.Error())
	}

	sock.AddTransport(ipc.NewTransport())
	sock.AddTransport(tcp.NewTransport())

	if url, ok = cfg.Get("url"); !ok {
		bolted.Die("No URL provided in the configuration")
	}

	if err = sock.Listen(url); err != nil {
		bolted.Die("Cannot listen on socket for emit: %s", err.Error())
	}

	return &emitConnection{sock, cfg}
}

func (ec *emitConnection) Emit(msg string) {
	if err := ec.sock.Send([]byte(msg)); err != nil {
		bolted.Die("Cannot emit: %s", err.Error())
	}
}
