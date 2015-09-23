# slack-incoming-webhooks [![Build Status](https://travis-ci.org/monochromegane/slack-incoming-webhooks.svg?branch=master)](https://travis-ci.org/monochromegane/slack-incoming-webhooks)

A Slack Incoming Webhooks client in Go.

## Usage

### Incoming Webhooks

```sh
$ slack-incoming-webhooks \
    -icon-emoji ":ghost:" \
    -text "This is posted to #general and comes from a bot named webhookbot." \
    YOUR_WEBHOOK_URL
```

### Richly-formatted messages

You can use attachment and fields parameters.

```sh
$ slack-incoming-webhooks \
    -attachment-color "good" \
    -attachment-text "This is posted to #general and comes from a bot named webhookbot." \
    YOUR_WEBHOOK_URL
```

```sh
$ slack-incoming-webhooks \
    -attachment-color "good" \
    # First field
    -attachment-field-title "Project" \
    -attachment-field-value "Awesome Project" \
    -attachment-field-short "true" \
    # Second field
    -attachment-field-title "Environment" \
    -attachment-field-value "production" \
    -attachment-field-short "true" \
    YOUR_WEBHOOK_URL
```

### As a Go library

You can use as a Go library.

```go
import slack "github.com/monochromegane/slack-incoming-webhooks"

func main() {
	slack.Client{
		WebhookURL: YOUR_WEBHOOK_URL,
	}.Post(&slack.Payload{
		Text: "Text",
	})
}
```

## Support parameters

**All** Incoming Webhooks and Attachments parameters is supported :beer: !

See Slack API doc.

- [Incoming Webhooks](https://api.slack.com/incoming-webhooks)
- [Attachments and field](https://api.slack.com/docs/attachments)


## Installation

```sh
$ go get github.com/monochromegane/slack-incoming-webhooks/...
```

## Contribution

1. Fork it
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

## License

[MIT](https://github.com/monochromegane/slack-incoming-webhooks/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)

