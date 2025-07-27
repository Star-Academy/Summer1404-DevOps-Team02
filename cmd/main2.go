package main

import (
  "fmt"
  "net"
  "os"
  "time"

  "golang.org/x/net/icmp"
  "golang.org/x/net/ipv4"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Usage: go run main.go <destination>")
    return
  }
  dest := os.Args[1]
  maxHops := 30
  timeout := time.Second * 2

  destAddr, err := net.ResolveIPAddr("ip4", dest)
  if err != nil {
    fmt.Println("Resolve error:", err)
    return
  }

  conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
  if err != nil {
    fmt.Println("Listen error:", err)
    return
  }
  defer conn.Close()

  for ttl := 1; ttl <= maxHops; ttl++ {
    conn.IPv4PacketConn().SetTTL(ttl)

    msg := icmp.Message{
      Type: ipv4.ICMPTypeEcho,
      Code: 0,
      Body: &icmp.Echo{
        ID:   os.Getpid() & 0xffff,
        Seq:  ttl,
        Data: []byte("HELLO-RTR"),
      },
    }
    b, _ := msg.Marshal(nil)
    start := time.Now()
    _, err := conn.WriteTo(b, &net.IPAddr{IP: destAddr.IP})
    if err != nil {
      fmt.Println("WriteTo error:", err)
      return
    }

    conn.SetReadDeadline(time.Now().Add(timeout))
    reply := make([]byte, 1500)
    n, peer, err := conn.ReadFrom(reply)
    duration := time.Since(start)

    if err != nil {
      fmt.Printf("%2d  *\n", ttl)
      continue
    }

    rm, err := icmp.ParseMessage(1, reply[:n])
    if err != nil {
      fmt.Println("ParseMessage error:", err)
      return
    }

    switch rm.Type {
    case ipv4.ICMPTypeTimeExceeded:
      fmt.Printf("%2d  %v  %.2fms\n", ttl, peer, float64(duration.Microseconds())/1000)
    case ipv4.ICMPTypeEchoReply:
      fmt.Printf("%2d  %v  %.2fms (destination reached)\n", ttl, peer, float64(duration.Microseconds())/1000)
      return
    default:
      fmt.Printf("%2d  ?\n", ttl)
    }
  }
}
