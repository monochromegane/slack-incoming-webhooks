package slack_incoming_webhooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	WebhookURL string
}

func (c Client) Post(p *Payload) error {

	json, err := json.Marshal(p)
	if err != nil {
		return err
	}

	resp, err := http.PostForm(c.WebhookURL, url.Values{"payload": []string{string(json)}})
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		return fmt.Errorf("status code: %d, response body: %s", resp.StatusCode, body)
	}
	return nil
}
