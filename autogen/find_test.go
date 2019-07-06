package autogen

import "testing"

func Test_NewFT(t *testing.T) {
	ft, err := NewFT()
	if err != nil {
		t.FailNow()
	}

	if ft.m["109.111.96.0"].File != "ad.zone" {
		t.FailNow()
	}
	c, z, ok := ft.Find("109.111.96.0")
	if !ok {
		t.Fatalf("can't find IP")
	}
	println(c)
	if z != "ad.zone" {
		t.Fatalf("expected: ad.zone")
	}

}
