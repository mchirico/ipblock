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
	split := strings.Split(r, "\n")
	for _, v := range split {
		ip := ExtractIP4(v)
		if len(ip) > 0 {
			for _, v2 := range ip {
				array = append(array, v2)
			}
		}

	}

	if len(array) != 1884 {
		t.Fatalf("Length not correct:%d\n", len(array))
	}

}

func TestStats(t *testing.T) {
	file := "../test-fixtures/mail.log"
	array := BuildArray(file)
	if len(array) != 1884 {
		t.Fatalf("Length not correct:%d\n", len(array))
	}
	m := Stats(array)
	for _, v := range m {
		fmt.Printf("zone: %s\n", v.Zone)
		fmt.Printf("  block: %s\n", v.Block)
		fmt.Printf("  count: %d\n", v.Count)
	}

	fmt.Printf("\nNon us.zone rules:\n\n")
	for _, v := range m {
		if v.Zone != "us.zone" {
			if len(v.Block) > 4 {
				fmt.Printf("iptables -I INPUT 50	-s %s -j DROP\n", v.Block)
			}
		}

	}
	fmt.Printf("iptables-save | awk '!seen[$0]++'|iptables-restore\n")
}

func TestDisplay(t *testing.T) {
	file := "../test-fixtures/mail.log"
	r, err := Display(file)
	if err != nil {
		t.Fatalf("Could not read file.")
	}
	if !strings.Contains(r, "for FILE in") {
		t.Logf("%s\n", r)
		t.Fatalf("Incorrect string generated")
	}
	if !strings.Contains(r, "block: 5.0.0.0/16") {
		t.Fatalf("Expecting: block: 5.0.0.0/16")
	}
	if !strings.Contains(r, "iptables -I INPUT 50 -s %%s -j") {
		t.Fatalf("Expecting: percent s ")
	}

	if strings.Contains(r,"MISSING") {
		t.Fatalf("Need input parameter?")
	}

}
