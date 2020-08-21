package SecureClient

/*
	Create new SSL-pinned client
	
	- The dialerConfig contains all of the generated pins
	- dialerConfig only needs to be made once, it should be shared and used on creation of new clients

*/

import (
	"fmt"
	"strings"
	"errors"
	"net/url"
	"net/http"
)


//--- Wrapper for go's standard implementation of the http.RoundTripper interface
type SSLPinnerTransport struct{
	DefaultTransport *http.Transport
	BadPinDetected   func(proxy string)
	Proxy            string
}

//--- Implementing http.RoundTripper
func (t *SSLPinnerTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := t.DefaultTransport.RoundTrip(r)
	
	if err != nil {
		if strings.Contains(fmt.Sprintf("%v", err), "x509: certificate signed by unknown authority") {
			t.BadPinDetected(t.Proxy)
			return resp, errors.New("Security Error - Unsafe proxy")
		}
	}
	
	return resp, err
}

func (p *SSLPinner) NewClient(proxy string) (*http.Client, error) {
	client, defaultTransport := &http.Client{}, &http.Transport{}

	if proxy != "" {	
		u, err := url.Parse(proxy)
		if err != nil { return client, err }
		defaultTransport = &http.Transport{
			DialTLS: p.DialerConfig.NewDialer(),
			Proxy:   http.ProxyURL(u),
		}
	} else {
		defaultTransport = &http.Transport{
			DialTLS: p.DialerConfig.NewDialer(),
			//--- No proxy
		}
	}
	
	client.Transport = &SSLPinnerTransport{
		//--- Standard go transport - If you have a custom transport, plug it in here
		//---   uTLS, etc.
		DefaultTransport: defaultTransport,
		
		//--- Triggered when SSL pin is not matched, ie mitm proxy passes an invalid cert
		BadPinDetected:   p.BadPinDetected,
		
		//--- Store proxy here for logging in BadPinDetected
		//--- It would be easy to see if the proxy seems like a MITM proxy
		//--- http://localhost:5555,  http://localhost:8080 etc.  
		Proxy: proxy,
	}
	
	return client, nil
}