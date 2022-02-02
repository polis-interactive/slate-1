package main

import (
	"fmt"
	"github.com/pborges/huejack"
	"log"
	"net"
	"os"
)

func main() {
	huejack.SetLogger(os.Stdout)
	huejack.Handle("test", func(req huejack.Request, res *huejack.Response) {
		fmt.Println("im handling test from", req.RemoteAddr, req.RequestedOnState)
		res.OnState = req.RequestedOnState
		// res.ErrorState = true //set ErrorState to true to have the echo respond with "unable to reach device"
		return
	})

	// it is very important to use a full IP here or the UPNP does not work correctly.
	// one day ill fix this

	address := fmt.Sprintf("%s:%d", GetOutboundIP().String(), 5000)
	log.Println(address)
	panic(huejack.ListenAndServe(address))
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
