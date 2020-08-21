# SecureClient
 ## Automatic SSL Pinning

 ## Secure the standard net/http client with **SSL pinning** to prevent users from sniffing requests with a **Man-In-The-Middle** proxy
 
This package takes in a list of hosts and provides a function to create net/http clients with SSL Pinning. 
For best practices
```
pinner := SecureClient.New(hosts, requireAll, BadPinDetected)
```
Should be called on startup (**typically in main()**), and whenever a client is needed you may call 
```
client := pinner.NewClient(proxy)
```
The SSL Pins only need to be generated once per program run _max_, which is done in **SecureClient.New()**
Future plans include storing SSL Pins to file and updating them once a week or so. I'm pretty sure the pins shouldnt change for a good bit of time.
 
 
# Example Usage

## Example 1 - See [tests](https://github.com/rustler47/SecureClient/blob/master/testing/compare/main.go)

## Example 2 
```
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
```

## References
[tam7t/hpkp](https://github.com/tam7t/hpkp/)
