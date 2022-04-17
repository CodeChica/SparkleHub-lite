package domain

import (
	"errors"
	"regexp"

	"github.com/codechica/SparkleHub-lite/pkg/pls"
)

type Sparkle struct {
	ID       string `json:"id" jsonapi:"primary,sparkles"`
	Sparklee string `json:"sparklee" jsonapi:"attr,sparklee"`
	Reason   string `json:"reason" jsonapi:"attr,reason"`
}

var SparkleRegex = regexp.MustCompile(`\A\s*(?P<sparklee>@\w+)\s+(?P<reason>.+)\z`)
var SparkleeIndex = SparkleRegex.SubexpIndex("sparklee")
var ReasonIndex = SparkleRegex.SubexpIndex("reason")

var ReasonIsRequired = errors.New("Reason is required")
var SparkleIsEmpty = errors.New("Sparkle is empty")
var SparkleIsInvalid = errors.New("Sparkle is invalid")
var SparkleeIsRequired = errors.New("Sparklee is required")

func NewSparkle(text string) (*Sparkle, error) {
	if len(text) == 0 {
		return nil, SparkleIsEmpty
	}

	matches := SparkleRegex.FindStringSubmatch(text)
	if len(matches) == 0 {
		return nil, SparkleIsInvalid
	}

	return &Sparkle{
		ID:       pls.GenerateULID(),
		Sparklee: matches[SparkleeIndex],
		Reason:   matches[ReasonIndex],
	}, nil
}

func (s *Sparkle) Validate() error {
	if s.Sparklee == "" {
		return SparkleeIsRequired
	}
	if s.Reason == "" {
		return ReasonIsRequired
	}
	return nil
}
