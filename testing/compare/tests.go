package main

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	sc "AI/Common/SecureClient"
)


func TestGoodProxy(host, proxy string, pinner *sc.SSLPinner) {
	fmt.Println("\n================== Good Connection Test ==================")
	fmt.Println("Testing pinned request thru legit proxy\n")
	
	client, err := pinner.NewClient(proxy)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	client.Timeout = time.Second * 10
	
	if err := TestPin(host, client); err != nil { fmt.Println(err) }	
}

func TestMITMProxy(host, proxy string, pinner *sc.SSLPinner) {
	fmt.Println("\n================== Bad Connection Test ==================")
	fmt.Println("Testing pinned request thru man-in-the-middle proxy\n")
	
	client, err := pinner.NewClient(proxy)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	client.Timeout = time.Second * 10
	
	if err := TestPin(host, client); err != nil { fmt.Println(err) }
}

func TestPin(host string, client *http.Client) error {
	resp, err := client.Get("https://"+host)
	if err != nil {
		if strings.Contains(fmt.Sprintf("%v",err), "Unsafe proxy") {
			return err
		}
		fmt.Println(fmt.Sprintf("Request sent - Proxy/Connection is secure!\n(Error with request: %v)", err))
		return nil
	}
	fmt.Println(fmt.Sprintf("Request sent - Proxy/Connection is secure! (%v says: %v)", host, resp.Status))
	return nil
}

func pause() {
	p := make(chan bool, 1)
	<-p
}
