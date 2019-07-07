package fileAnalysis

import (
	"fmt"
	"github.com/mchirico/ipblock/autogen"
	"io/ioutil"
	"regexp"
	"strings"
)

func ReadFile(file string) string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func ExtractIP4(s string) [][]byte {

	re, _ := regexp.Compile(`(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])`)
	return re.FindAll([]byte(s), -1)

}

func BuildArray(file string) [][]byte {
	array := [][]byte{}

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
	return array
}

type STAT struct {
	Count int
	Zone  string
	Block string
}

func Stats(b [][]byte) map[string]*STAT {
	m := map[string]*STAT{}
	ft, err := autogen.NewFT()
	if err != nil {
		return nil
	}

	for _, v := range b {
		s := string(v)
		if val, okay := m[s]; okay {
			val.Count += 1
		} else {
			if cider, file, okay := ft.Find(s); okay {
				m[s] = &STAT{1,file,cider}
			}

		}
	}
	return m

}

func Display(file string) {
	array := BuildArray(file)
    if len(array) < 1 {
    	return
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
