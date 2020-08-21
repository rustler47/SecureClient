package main

import (
	"fmt"
	"github.com/rustler47/SecureClient"
)

func main() {
	fmt.Println("SSL Pinning test\n\n")

	MITMProxy  :="http://localhost:5555"

	hosts := []string{ "kith.com" }

	BadPinDetected := func(proxy string){
		fmt.Println("WARNING! Failed SSL pinning - Invalid cert detected\n", "Proxy:", proxy)
	}
	
	pinner, err := SecureClient.New(hosts, true, BadPinDetected)
	if err != nil { return }

	client, err := pinner.NewClient(MITMProxy)
	if err != nil { return }
	
	client.Get("https://kith.com")
	
	pause := make(chan bool, 1)
	<-pause
}

