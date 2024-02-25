package receiver

import (
  "fmt"
  "net"
  "bufio"
)

func Listen(ports []int) {
  for _, port := range ports {
    go listenOnPort(port)
  }
}

func listenOnPort(port int){
  address := fmt.Sprintf(":%d", port)
  listener, err := net.Listen("tcp", address)
  if err != nil {
    fmt.Printf("Error listening on port %d: %v\n", port, err)
    return
  }

  defer listener.Close()

  fmt.Printf("Listening for messages on port %d\n", port)

  for {
    conn, err := listener.Accept()
    if err != nil{
      fmt.Printf("Error accepting connection : %v\n", err)
      continue
    }

    handleConnection(conn)

  }
}

func handleConnection(conn net.Conn){
  defer conn.Close()

  scanner := bufio.NewScanner(conn)
  for scanner.Scan(){
    line := scanner.Text()
    fmt.Println("Received:",line)
  }

  if err := scanner.Err()
  err != nil {
    fmt.Println("Error reading: ", err)
  }
}
