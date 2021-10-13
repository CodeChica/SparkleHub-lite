package main

import (
	"errors"
	"regexp"
)

type Sparkle struct {
	Sparklee string `json:"sparklee"`
	Reason   string `json:"reason"`
}

var SparkleRegex = regexp.MustCompile(`\A\s*(?P<sparklee>@\w+)\s+(?P<reason>.+)\z`)
var SparkleeIndex = SparkleRegex.SubexpIndex("sparklee")
var ReasonIndex = SparkleRegex.SubexpIndex("reason")

var SparkleIsEmpty = errors.New("Sparkle is empty")
var SparkleIsInvalid = errors.New("Sparkle is invalid")

func NewSparkle(text string) (*Sparkle, error) {
	if len(text) == 0 {
		return nil, SparkleIsEmpty
	}

	matches := SparkleRegex.FindStringSubmatch(text)
	if len(matches) == 0 {
		return nil, SparkleIsInvalid
	}

	return &Sparkle{
		Sparklee: matches[SparkleeIndex],
		Reason:   matches[ReasonIndex],
	}, nil
}
