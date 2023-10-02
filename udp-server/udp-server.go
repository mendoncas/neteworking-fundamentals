package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
  server, err := net.ListenPacket("udp", ":3000")
  if err != nil {
    log.Fatal(err)
  }
  defer server.Close()

  for {
    buf := make([]byte, 1024)
    _, addr, err := server.ReadFrom(buf)
    if err != nil {
      continue
    }
    go response(server, addr, buf)
  }
}

func response(udpServer net.PacketConn, addr net.Addr, buf []byte){
  time := time.Now().Format(time.ANSIC)
  responseStr := fmt.Sprintf("packet received at: %v; \n message: %v", time, string(buf))
  udpServer.WriteTo([]byte(responseStr), addr)
}
