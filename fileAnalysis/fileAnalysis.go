package fileAnalysis

import (
	"bytes"
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
				m[s] = &STAT{1, file, cider}
			}

		}
	}
	return m

}

func Display(file string) (string, error) {

	var b bytes.Buffer
	nonUSzones := 0

	array := BuildArray(file)
	zsummary := map[string]string{}
	if len(array) < 1 {
		return "", fmt.Errorf("No IPs in file?")
	}

	m := Stats(array)
	for _, v := range m {
		fmt.Fprintf(&b, "zone: %s\n", v.Zone)
		fmt.Fprintf(&b, "  block: %s\n", v.Block)
		fmt.Fprintf(&b, "  count: %d\n", v.Count)
		if v.Zone != "us.zone" {
			zsummary[v.Zone] = "yes"
			nonUSzones += 1
		}
	}

	fmt.Fprintf(&b, "\nConsider blocking the following:\n\n")
	for k, _ := range zsummary {
		fmt.Fprintf(&b, "%s\n", k)
	}

	fmt.Fprintf(&b, "\n\n")
	script := `

--- SCRIPT ---

#!/bin/bash
mkdir rules
cd rules
curl http://www.ipdeny.com/ipblocks/data/countries/all-zones.tar.gz -o all-zones.tar.gz
tar -xzf all-zones.tar.gz


for FILE in %s
do
  echo -e '#!/bin/bash\n' > ${FILE}.sh
  awk '{printf("iptables -I INPUT 50 -s %%%%s -j DROP\n",$1)}' "${FILE}.zone"  >> "${FILE}.sh"
  echo -e 'iptables-save | awk '"'"'!seen[$0]++'"'"'|iptables-restore\n' >> "${FILE}.sh"
  chmod 700 "${FILE}.sh"
done


--- ALL IN TEXT ---
--- Manually add to iptables file --

mkdir rules
cd rules
curl http://www.ipdeny.com/ipblocks/data/countries/all-zones.tar.gz -o all-zones.tar.gz
tar -xzf all-zones.tar.gz


echo -e '#\n' > all.txt
for FILE in %s
do
  awk '{printf("-A INPUT -s %%%%s -j DROP\n",$1)}' "${FILE}.zone"  >> all.txt
done

`
	zoneList := ""
	for k, _ := range zsummary {
		if !strings.Contains(k, ".") {
			continue
		}
		r := strings.Split(k, ".")
		zoneList = zoneList + " " + r[0]

	}

	if nonUSzones > 1 {
		fmt.Fprintf(&b, script, zoneList,zoneList)
	}
	return b.String(), nil

}
