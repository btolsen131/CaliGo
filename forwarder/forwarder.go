package forwarder

import(
	"fmt"
	"net"
)

func ForwardMessage(destination string, message string) {
    conn, err := net.Dial("tcp", destination)
    if err != nil {
        fmt.Printf("Error connecting to destination %s: %v\n", destination, err)
        return
    }
    defer conn.Close()

    // Send the message to the destination
    _, err = fmt.Fprintf(conn, message)
    if err != nil {
        fmt.Printf("Error sending message to destination %s: %v\n", destination, err)
        return
    }

    fmt.Printf("Message sent to destination %s\n", destination)
}