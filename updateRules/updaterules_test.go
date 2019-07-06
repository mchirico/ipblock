package updateRules

import (
	"github.com/mchirico/ipblock/request"
	"github.com/mchirico/tlib/util"
	"strings"
	"testing"
)

func TestDownload(t *testing.T) {
	defer util.NewTlib().ConstructDir()()

	url := "http://www.ipdeny.com/ipblocks/data/countries/all-zones.tar.gz"
	urlMD5 := "http://www.ipdeny.com/ipblocks/data/countries/MD5SUM"
	file := "all-zones.tar.gz"

	downloadFile(file, url)

	r, err := request.Get(urlMD5)
	if err != nil {
		// We don't care to further test
		return
	}

	if !strings.Contains(r, FileMD5(file)) {
		t.Fatalf("No match")

	}

}
