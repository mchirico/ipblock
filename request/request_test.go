package request

import "testing"

func TestGetWithBrazil(t *testing.T) {

	url := "http://www.ipdeny.com/ipblocks/data/countries/br.zone"
	r, err := Get(url)
	if err != nil {
		t.Fatalf("Can't get data from url\n")
	}
	if len(r) < 15678 {
		t.Fatalf("Length tool low...\n")
	}

}
