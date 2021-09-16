package main

import (
	"errors"
	"regexp"
)

type Sparkle struct {
	Sparklee string `json:"sparklee"`
	Reason   string `json:"reason"`
}

func NewSparkle(body string) (*Sparkle, error) {
	username, reason, err := parse(body)
	if err != nil {
		return nil, err
	}

	return &Sparkle{
		Sparklee: username,
		Reason:   reason,
	}, nil
}

func parse(body string) (string, string, error) {
	if len(body) == 0 {
		return "", "", errors.New("sparkle is empty")
	}

	regex := regexp.MustCompile(`\A\s*(?P<username>@\w+)\s+(?P<reason>.+)\z`)
	matches := regex.FindStringSubmatch(body)

	if len(matches) == 0 {
		return "", "", errors.New("sparkle is invalid")
	}

	username := matches[regex.SubexpIndex("username")]
	reason := matches[regex.SubexpIndex("reason")]
	return username, reason, nil
}
