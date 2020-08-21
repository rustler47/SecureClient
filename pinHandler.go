package SecureClient

/*
	SSL Pin Handler
		- Store DialerConfig for reuse

	Todo:
		- Store pins to file, only generate new pins like once a week instead of every bot startup
			- Encrypt file to keep snoopers out
			
	Note: Logging is at a minimum in here. We want to be somewhat secrative
*/

import (
	"fmt"
	"github.com/tam7t/hpkp"
)

type SSLPinner struct {
	RequireAll     bool     //--- Require all pins to be generated properly
	Hosts          []string //--- List of hosts pinned/to pin
	
	DialerConfig   *hpkp.DialerConfig
	BadPinDetected func(string)
}
	
func New(hosts []string, requireAll bool, badPinDetected func(p string)) (*SSLPinner, error) {
	pinner := &SSLPinner{
		Hosts:          hosts,
		RequireAll:     requireAll,
		BadPinDetected: badPinDetected,
	}
	
	err := pinner.CreateDialer()
	
	return pinner, err
}

//--- Most useful function in go
func Crash(parent string) {
	if r := recover(); r != nil {//---Handle crashes
		fmt.Printf("%v Crashed!!!!\nError: %v", parent, r)
	}
}
