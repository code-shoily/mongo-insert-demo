package main

import (
	"log"
	"net"
	"strconv"

	"gopkg.in/mgo.v2"

	"github.com/code-shoily/insert_server/srv"
)

func main() {
	srv.SetLogger()

	server, err := net.Listen("tcp", ":"+strconv.Itoa(srv.NetPort))
	if err != nil {
		log.Fatal("[NET ERROR] - " + err.Error())
	}

	session, err := mgo.Dial(srv.MongoHost)
	if err != nil {
		log.Fatal("[MONGO ERROR] - " + err.Error())
	}

	// Setting up TCP connections
	connections := srv.ClientConnections(server)

	// Setting up Mongo parameters
	session.SetMode(mgo.Monotonic, true)

	log.Println("Server listening at port: " + strconv.Itoa(srv.NetPort))
	for {
		go srv.HandleConnection(<-connections, session)
	}
}
