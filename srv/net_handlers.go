package srv

import (
	"bufio"
	"fmt"
	"net"
	"sync"

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
				fmt.Println(err.Error())
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()

	return ch
}

// HandleConnection handles the connection. In this case, inserts the received
// string into the store by first converting it to JSON format.
func HandleConnection(client net.Conn, waitGroup *sync.WaitGroup, session *mgo.Session) {
	defer waitGroup.Done()
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	b := bufio.NewReader(client)

	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}

		// TODO Convert this into SerializeLocation for real data
		data, err := SerializeTodo(line)
		if err != nil {
			fmt.Println(err.Error())
			client.Write([]byte(err.Error()))
			client.Close()
		}

		data.Save(sessionCopy)
		client.Close()
	}
}
