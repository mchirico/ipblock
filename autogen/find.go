package autogen

import (
	"fmt"
	"strings"
)

type FT struct {
	m map[string]*MT
}

func NewFT() (*FT, error) {
	m, err := ConvertJSON()
	if err != nil {
		return nil, err
	}
	return &FT{m}, err
}

func (ft *FT) Find(ip string) (string, string, bool) {

	s := strings.Split(ip, ".")

	try := [3]string{}

	try[0] = fmt.Sprintf("%s.%s.%s.0", s[0], s[1], s[2])
	try[1] = fmt.Sprintf("%s.%s.0.0", s[0], s[1])
	try[2] = fmt.Sprintf("%s.0.0.0", s[0])

	for i := 0; i < 3; i++ {
		if r, ok := ft.m[try[i]]; ok {
			return r.Cider, r.File, ok
		}
	}

	return "", "", true
}
