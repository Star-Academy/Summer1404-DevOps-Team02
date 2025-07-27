package services

import (
	"cmp"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"slices"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type Result struct {
	TTL     int
	IP      string
	Latency string
}

func createPacket(ttl int) *icmp.Message {

	return &icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   int(rand.Int32()),
			Seq:  ttl,
			Data: []byte("HELLO-RTR"),
		},
	}
}

func sendPacket(ttl int, dest *net.IPAddr) (*Result, bool) {
	var (
		res Result = Result{
			TTL: ttl,
		}
	)
	// log.Println(dest.IP.String())
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Printf("Failed to connect to ip in goroutine, %s\n", err)
		return nil, false
	}
	defer conn.Close()

	conn.IPv4PacketConn().SetTTL(ttl)
	req := createPacket(ttl)
	reqBytes, err := req.Marshal(nil)
	if err != nil {
		log.Println(err)
		return nil, false
	}

	if _, err := conn.WriteTo(reqBytes, dest); err != nil {
		log.Println(err)
		return nil, false
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	reply := make([]byte, 1024)
	s := time.Now()
	n, peer, err := conn.ReadFrom(reply)
	latency := time.Now().Sub(s)

	res.Latency = fmt.Sprintf("%dms", latency.Milliseconds())
	if err != nil {
		log.Println(err)
		res.IP = "*"
		return &res, false
	}

	rm, err := icmp.ParseMessage(1, reply[:n])
	res.IP = peer.String()
	switch rm.Type {
	case ipv4.ICMPTypeTimeExceeded:
	case ipv4.ICMPTypeEchoReply:
		return &res, true

	default:
		return nil, false
	}

	return &res, false
}

func Trace(dest string) []*Result {

	var (
		Maxhops = 30
	)

	ip, err := net.ResolveIPAddr("ip4", dest)
	if err != nil {
		log.Println(err)
	}
	results := make([]*Result, 1)
	for ttl := 1; ttl <= Maxhops; ttl++ {

		res, isLast := sendPacket(ttl, ip)
		if res != nil {
			results = append(results, res)
		}
		if isLast {
			break
		}

	}

	slices.SortFunc(results, func(a, b *Result) int {
		if a == nil || b == nil {
			return 0
		}
		return cmp.Compare(a.TTL, b.TTL)
	})
	for _, v := range results {
		if v != nil {
			fmt.Println(v.IP)
		}
	}

	return results
}
