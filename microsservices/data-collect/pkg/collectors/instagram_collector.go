package collectors

import (
	"fmt"
	"io"
	"net/http"
)

type InstagramCollector struct {
	ClientID    string
	RedirectUri string
}

func NewInstagramCollector(ClientID, RedirectUri string) *InstagramCollector {
	return &InstagramCollector{ClientID: ClientID, RedirectUri: RedirectUri}
}

func (i *InstagramCollector) Auth() (token string, err error) {
	req, err := http.NewRequest("POST", "https://api.instagram.com/oauth/authorize", nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(body)

	return token, nil

}
func (i *InstagramCollector) Collect() error {
	return nil
}
