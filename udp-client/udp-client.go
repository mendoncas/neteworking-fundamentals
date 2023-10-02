package main

import (
	"net"
	"os"
	"time"
)

func main(){
  server, err := net.ResolveUDPAddr("udp", ":3000")
  if err != nil {
    println("failed resolving address: ", err.Error())
    os.Exit(1)
  }

  conn, err := net.DialUDP("udp", nil, server)
  if err != nil {
    println("failed starting the connection: ", err.Error())
    os.Exit(1)
  }
  defer conn.Close()

  for {
    _, err = conn.Write([]byte("This is a UDP message!!"))
    if err != nil {
      println("failed write data: ", err.Error())
    }
    received := make([]byte, 1024)
    _, err = conn.Read(received)
    if err != nil {
      println("read data failed:", err.Error())
    }
    println(string(received))
    time.Sleep(7 * time.Second)
  }
}
