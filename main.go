package main

import (
	"net"
	"strconv"
	"sync"

	"gopkg.in/mgo.v2"

	"github.com/code-shoily/insert_server/srv"
)

func main() {
	server, err := net.Listen("tcp", ":"+strconv.Itoa(srv.NetPort))
	session, err := mgo.Dial(srv.MongoHost)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error())
	}

	// Setting up TCP connections
	connections := srv.ClientConnections(server)

	// Setting up Mongo parameters
	session.SetMode(mgo.Monotonic, true)
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)

	for {
		go srv.HandleConnection(<-connections, &waitGroup, session)
	}
}
