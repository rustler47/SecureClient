# SecureClient
 Automatic SSL Pinning for golang net/http client
 Secure the standard net/http client with **SSL pinning** to prevent users from sniffing requests with a **Man-In-The-Middle** proxy
 
## Example Usage

Example 1 - See [tests](https://github.com/rustler47/SecureClient/blob/master/testing/main.go)

Example 2 
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
