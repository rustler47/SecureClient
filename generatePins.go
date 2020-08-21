package SecureClient

/*
	Generate SSL pins
	
	- Concurrently dial all hosts to generate pins for client
	- Pins are stored in the *hpkp.DialerConfig, which is to be reused in each client made

*/

import (
	"sync"
	"crypto/tls"
	"github.com/tam7t/hpkp"
)

type PinnedSite struct {
	Host string
	Pins []string
	Failed bool
}

func (p *SSLPinner) GeneratePins(pinChannel chan PinnedSite){
	var wait sync.WaitGroup

	for _, h := range p.Hosts {
		wait.Add(1)
		//--- Concurrent because it's go, so that means its ez ( and we like speed )
		go func(w *sync.WaitGroup, host string) {
			defer Crash("GeneratePin")
			defer w.Done()
			if pins, err := GetSSLPins(host + ":443"); err == nil { 
				pinChannel <- PinnedSite{ Host: host, Pins: pins } 
			} else { 
				pinChannel <- PinnedSite{ Failed: true } 
			}
	
		}(&wait, h)
	}
	
	go func() {
		defer Crash("HandleDone")
		wait.Wait()
		close(pinChannel)
	}()
}

func GetSSLPins(server string) ([]string, error) {
	var pins []string
	conn, err := tls.Dial("tcp", server, &tls.Config{ InsecureSkipVerify: true })
	if err != nil { return pins, err }

	for _, cert := range conn.ConnectionState().PeerCertificates {
		pins = append(pins, hpkp.Fingerprint(cert))
	}

	return pins, nil
}