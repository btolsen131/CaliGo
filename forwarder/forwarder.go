package forwarder

import (
	"fmt"
	"net"
)

var DestinationsMap map[string]net.Conn

func init() {
	DestinationsMap = make(map[string]net.Conn)
}

func ForwardMessage(destination string, message string) error{

	//Check to see if we already have a connection for this destination
	conn, ok := DestinationsMap[destination]

	//Make connection if we dont have the connection in our dict
	if !ok {
		newConn, err := net.Dial("tcp", destination)
		if err != nil {
			return fmt.Errorf("error connection to destination %s: %v", destination, err)
		}
		DestinationsMap[destination] = newConn
		conn = newConn
	}
	// Send the message to the destination
	_, err := fmt.Fprintf(conn, message)
	if err != nil {
		return fmt.Errorf("error sending message to destination %s: %v", destination, err)
	}

	fmt.Printf("Message sent to destination %s\n", destination)
	return nil
}
