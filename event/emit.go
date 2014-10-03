package event

import (
	"bitbucket.org/gdamore/mangos"
	"bitbucket.org/gdamore/mangos/protocol/pub"
	"bitbucket.org/gdamore/mangos/transport/ipc"
	"bitbucket.org/gdamore/mangos/transport/tcp"
	"github.com/dustinrc/bolted"
)

type Emitter interface {
	bolted.Configurer
	Emit()
}

type emitConnection struct {
	sock mangos.Socket
}

func NewEmitConnection(url string) *emitConnection {
	var sock mangos.Socket
	var err error

	if sock, err = pub.NewSocket(); err != nil {
		bolted.Die("Cannot initiate socket for emit: %s", err.Error())
	}

	sock.AddTransport(ipc.NewTransport())
	sock.AddTransport(tcp.NewTransport())

	if err = sock.Listen(url); err != nil {
		bolted.Die("Cannot listen on socket for emit: %s", err.Error())
	}

	return &emitConnection{sock}
}

func (ec *emitConnection) Emit(msg string) {
	if err := ec.sock.Send([]byte(msg)); err != nil {
		bolted.Die("Cannot emit: %s", err.Error())
	}
}
