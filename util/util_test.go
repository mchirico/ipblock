package util

import (
	"github.com/mchirico/ipblock/request"
	"strings"
	"testing"
)

func TestBegin(t *testing.T) {
	url := "http://www.ipdeny.com/ipblocks/data/countries/br.zone"
	r, err := request.Get(url)
	if err != nil {
		t.Fatalf("Can't get data from url\n")
	}
	if len(r) < 15678 {
		t.Fatalf("Length tool low...\n")
	}
}

func TestBuildBlock(t *testing.T) {
	url := "http://www.ipdeny.com/ipblocks/data/countries/br.zone"
	rules, err := BuildBlock(url)
	if err != nil {
		t.Fatalf("err:%v\n", err)
	}
	s := strings.Split(rules, "\n")

	if len(s) < 1000 {
		t.FailNow()
	}

	for _, v := range s {
		if !strings.Contains(v, "/") {
			if v != "" {
				t.Fatalf("output bad:->%s<-\n", v)
			}
			println(v)
		}
	}

}
