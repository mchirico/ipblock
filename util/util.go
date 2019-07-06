package util

import (
	"fmt"
	"github.com/mchirico/ipblock/request"
	"strings"
)

func BuildBlock(url string) (string, error) {
	rules := ""

	r, err := request.Get(url)
	if err != nil {
		return rules, err
	}
	list := strings.Split(r, "\n")

	for _, v := range list {
		if strings.Contains(v, "/") && strings.Contains(v, ".") {
			rule := fmt.Sprintf("ufw insert 12 deny from %s\n", v)
			rules += rule
		}
	}

	return rules, err
}
