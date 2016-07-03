package srv

import (
	"bufio"
	"log"
	"net"

	"gopkg.in/mgo.v2"
)

// ClientConnections returns a channel of connections.
func ClientConnections(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0

	go func() {
		for {
			client, err := listener.Accept()
			if err != nil {
				log.Println(err.Error())
				continue
			}
			i++
			log.Printf("[INFO] - %d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()

	return ch
}

// HandleConnection handles the connection. In this case, inserts the received
// string into the store by first converting it to JSON format.
func HandleConnection(client net.Conn, session *mgo.Session) {
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	b := bufio.NewReader(client)

	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}

		data, err := SerializeLocation(line)
		if err != nil {
			log.Println("[SERIALIZER ERROR] - " + err.Error())
			client.Close()
			break
		}
		data.Save(sessionCopy)
		client.Close()
	}
}
