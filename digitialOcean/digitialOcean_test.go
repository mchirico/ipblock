package digitialOcean

import (
	"os"
	"testing"
)

func TestDev(t *testing.T) {

	if os.Getenv("DOTOKEN") == "" {
		println("No access token")
		return
	}

	if os.Getenv("FIREWALLID") == "" {
		println("No firewall token")
		return
	}

	Dev()

}
