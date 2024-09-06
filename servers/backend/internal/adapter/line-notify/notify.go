package linenotify

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Option struct {
	token string
}

func NewOption(conf *viper.Viper) Option {
	return Option{
		token: conf.GetString("linenotify.token"),
	}
}

// LineNotify line notify
type LineNotify struct {
	Token   string
	Message Message
}

// Message the notify message
type Message struct {
	UpdateTime time.Time
	ID         string
	Name       string
	Event      string
}

// NewMessage create a new message
func NewLineNotify(option Option) *LineNotify {
	return &LineNotify{
		Token: option.token,
	}
}

// Send send the message
func (ln *LineNotify) Send() error {
	messageStr := "\n"
	if ln.Message.UpdateTime.IsZero() {
		messageStr += fmt.Sprintf("時間 :\t%v \n", time.Now().Format("2006-01-02 15:04:05"))
	} else {
		messageStr += fmt.Sprintf("時間 :\t%v \n", ln.Message.UpdateTime.Format("2006-01-02 15:04:05"))
	}

	if ln.Message.ID != "" {
		messageStr += fmt.Sprintf("編號 :\t%v \n", ln.Message.ID)
	}

	if ln.Message.Name != "" {
		messageStr += fmt.Sprintf("名稱 :\t%v \n", ln.Message.Name)
	}

	if ln.Message.Event != "" {
		messageStr += fmt.Sprintf("事件 :\t%v \n", ln.Message.Event)
	}

	values := url.Values{}
	values.Add("message", messageStr)

	body := strings.NewReader(values.Encode())

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "https://notify-api.line.me/api/notify", body)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ln.Token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	res.Body.Close()
	return nil
}
