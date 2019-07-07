package fileAnalysis

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Regex(t *testing.T) {

	array := [][]byte{}

	file := "../test-fixtures/mail.log"
	r := ReadFile(file)
	split := strings.Split(r,"\n")
	for _,v := range split {
		ip := ExtractIP4(v)
		if len(ip) > 0 {
			for _,v2 := range ip {
				array = append(array, v2)
			}
		}

	}

	if len(array) != 1884 {
		t.Fatalf("Length not correct:%d\n",len(array))
	}

}


func TestStats(t *testing.T) {
	file := "../test-fixtures/mail.log"
	array := BuildArray(file)
	if len(array) != 1884 {
		t.Fatalf("Length not correct:%d\n",len(array))
	}
	m := Stats(array)
	for _,v := range m {
		fmt.Printf("zone: %s\n",v.Zone)
		fmt.Printf("  block: %s\n",v.Block)
		fmt.Printf("  count: %d\n",v.Count)
	}

	fmt.Printf("\nNon us.zone rules:\n\n")
	for _,v := range m {
		if v.Zone != "us.zone" {
			if len(v.Block) > 4 {
				fmt.Printf("iptables -A INPUT -s %s -j DROP\n", v.Block)
			}
		}

	}
	fmt.Printf("iptables-save | awk '!seen[$0]++'|iptables-restore\n")
}

