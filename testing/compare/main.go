package main

import (
	"fmt"
	sc "github.com/rustler47/SecureClient"
)

func main() {
	defer pause()
	fmt.Println("SSL Pinning test\n\n")
	
	//--- (Control) Good proxy/local ip to test
	proxy := ""    //"http://user:pass@ip:port"
	
	//--- (Test trigger) Postman default proxy capture
	MITMProxy  :="http://localhost:5555"
	
	// --- fail if any site is not able to be pinned on creation
	requireAll := true 
	
	hosts := []string{
		"kith.com",
		"undefeated.com",
		"www.cncpts.com",
		"www.footlocker.com",
		"www.sneakersnstuff.com",
	}
	
	//--- Host from list to test pins w/ the two proxies above
	testHost := "undefeated.com"
	
	//--- BadPinDetected fires on SSL Pin failure
	//--- Should only really happen if they are using a Man-In-The-Middle proxy to sniff your bum
	BadPinDetected := func(proxy string){
		fmt.Println("WARNING! Failed SSL pinning - Invalid cert detected\n", "Proxy:", proxy)
		
		//--- Example usage: Flag the user's key and report the proxy that triggerd this func
		//KeyAPI.FlagSketchyUser(proxy, "Failed SSL Pin", user.Key) 
	}
	
	pinner, err := sc.New(hosts, requireAll, BadPinDetected)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Testing on "+testHost+"\n")
	
	TestGoodProxy(testHost, proxy,     pinner) 
	TestMITMProxy(testHost, MITMProxy, pinner) 
}