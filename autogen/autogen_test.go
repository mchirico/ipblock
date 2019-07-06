package autogen

import (
	"fmt"
	"testing"
)

func Test_ConvertJSON(t *testing.T) {
	m, err := ConvertJSON()
	if err != nil {
		t.FailNow()
	}

	if m["109.111.96.0"].File != "ad.zone" {
		t.FailNow()
	}

	fmt.Printf("m:=%v\n", m["109.111.96.0"].File)

}
