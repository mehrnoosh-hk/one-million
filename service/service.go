package service

import (
	"strings"
	"time"
)

type PayLoad struct {
	Text string `json:"text"`
}

func (p PayLoad) Process() string {
	time.Sleep(2000 * time.Millisecond)
	return strings.ToUpper(p.Text)
}
