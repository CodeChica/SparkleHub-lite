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
	Date     string `json:"date" jsonapi:"attr,date"`
}

var SparkleRegex = regexp.MustCompile(`\A\s*(?P<sparklee>@\w+)\s+(?P<reason>.+)\s+(?P<date>.+)\z`)
var SparkleeIndex = SparkleRegex.SubexpIndex("sparklee")
var ReasonIndex = SparkleRegex.SubexpIndex("reason")
var DateIndex = SparkleRegex.SubexpIndex("date")

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
		Date:     matches[DateIndex],
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
