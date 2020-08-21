package SecureClient

/*
	Create SSL Pinned TLS dialer
	
	- This is an INIT function, should only be called once per run max
	- New clients use the DialTLS created by the *hpkp.DialerConfig
*/

import (
	"errors"
	"github.com/tam7t/hpkp"
)

func (p *SSLPinner) CreateDialer() error {
	pinChannel := make(chan PinnedSite, 1000)
	
	//fmt.Println("Generating pins!")
	
	go p.GeneratePins(pinChannel)
	
	s := hpkp.NewMemStorage()
	
	for pinned := range pinChannel {
		if pinned.Failed && p.RequireAll { return errors.New("Error creating secure client!") }
		
		//fmt.Println(pinned.Host, "Secured -", len(pinned.Pins), "generated")
		
		s.Add(pinned.Host, &hpkp.Header{
			Permanent: true,
			Sha256Pins: pinned.Pins,
		})
	}
	
	p.DialerConfig = &hpkp.DialerConfig{
		Storage:   s,
		PinOnly:   true,
		TLSConfig: nil,
		Reporter: func(p *hpkp.PinFailure, reportUri string) {
			//fmt.Println("Pin failure: ", p)
		},
	}
	
	return nil
}
