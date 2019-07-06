package updateRules

import (
	"encoding/json"
	"fmt"
	"github.com/mchirico/ipblock/request"
	"github.com/mchirico/tlib/util"
	//"github.com/mchirico/ipblock/autogen"
	"os"
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

func downLoad() {
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
		panic("no download")

	}

}

//TODO: Download and create zone. FIX!
func Test_ReadZones(t *testing.T) {

	// Works but don't aways want too run this.
	return

	m := map[string]MT{}
	var mjson []byte
	var err error

	util.Mkdir("./junk")
	os.Chdir("./junk")
	downLoad()
	err = Unzip("./all-zones.tar.gz")
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	os.Chdir("./..")

	zones := ListZones("./junk")
	println(len(zones))

	for i, v := range zones {

		mjson, err = FileToJson(v, &m)
		if err != nil {
			t.FailNow()
		}
		if i < -3 {
			break
		}
	}

	json.Unmarshal(mjson, &m)
	tag := fmt.Sprintf("`%s`", mjson)
	CreateAutogen("../autogen/autogen.go", tag)
	fmt.Printf("%s\n", mjson)

	os.RemoveAll("./junk")
}
