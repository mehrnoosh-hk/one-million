package service

import (
	"math/rand"
	"strings"
	"time"
)

type PayLoad struct {
	Text string `json:"text"`
}

func (p PayLoad) Process() string {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	return strings.ToUpper(p.Text)
}
