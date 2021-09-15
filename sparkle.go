package main

import "strings"

type Sparkle struct {
	Sparklee string `json:"sparklee"`
	Reason   string `json:"reason"`
}

func NewSparkle(body string) *Sparkle {
	items := strings.SplitAfterN(body, " ", 2)
	if len(items) != 2 {
		return nil
	}
	username := strings.Trim(items[0], " ")
	reason := strings.Trim(items[1], " ")

	return &Sparkle{
		Sparklee: username,
		Reason:   reason,
	}
}
