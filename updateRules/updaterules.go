package updateRules

// http://www.ipdeny.com/ipblocks/data/countries/all-zones.tar.gz
// http://www.ipdeny.com/ipblocks/data/countries/MD5SUM

import (
	"crypto/md5"
	"fmt"
	"github.com/mchirico/ipblock/request"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func FileMD5(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func DownloadTest() (bool, error) {
	url := "http://www.ipdeny.com/ipblocks/data/countries/all-zones.tar.gz"
	urlMD5 := "http://www.ipdeny.com/ipblocks/data/countries/MD5SUM"
	file := "all-zones.tar.gz"

	downloadFile(file, url)

	r, err := request.Get(urlMD5)
	if err != nil {
		return false, err
	}

	if !strings.Contains(r, FileMD5(file)) {
		return false, err

	}

	return true, err
}
