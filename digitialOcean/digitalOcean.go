package digitialOcean

import (
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
	"log"
	"os"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func Dev() {
	pat := os.Getenv("DOTOKEN")
	firewallID := os.Getenv("FIREWALLID")

	tokenSource := &TokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()

	firewall, _, err := client.Firewalls.Get(ctx, firewallID)

	//firewall, _, err := client.Firewalls.Get(ctx)
	if err != nil {
		log.Fatalf("bad token?")
	}
	fmt.Printf("account:%v\n", firewall)

}
