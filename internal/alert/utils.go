package alert

import (
	"fmt"

	"github.com/slack-go/slack"
)

// Slack & Mattermost compatible
func generateAttachments(color string, log map[string]string) []slack.Attachment {
	reason := slack.Attachment{
		AuthorName: ":warning: teler Alert",
		Title:      log["category"],
		Color:      color,
	}
	request := slack.Attachment{
		Title: "Request",
		Text: fmt.Sprintf(
			"%s %s %s",
			log["request_method"], log["request_uri"], log["request_protocol"],
		),
		Color: color,
	}
	fields := slack.Attachment{
		Color: color,
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: log["time_local"],
				Short: true,
			},
			{
				Title: "IP Address",
				Value: log["remote_addr"],
				Short: true,
			},
			{
				Title: "User-Agent",
				Value: log["http_user_agent"],
				Short: true,
			},
			{
				Title: "Referrer",
				Value: log["http_referer"],
				Short: true,
			},
			{
				Title: "Status code",
				Value: log["status"],
				Short: true,
			},
			{
				Title: "Bytes sent",
				Value: log["body_bytes_sent"],
				Short: true,
			},
		},
	}

	return []slack.Attachment{reason, request, fields}
}
