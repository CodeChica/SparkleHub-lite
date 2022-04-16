package domain

import (
	"errors"
	"math/rand"
	"regexp"
	"time"

	"github.com/oklog/ulid/v2"
)

type Sparkle struct {
	ID       string `jsonapi:"primary,sparkles"`
	Sparklee string `json:"sparklee" jsonapi:"attr,sparkle"`
	Reason   string `json:"reason" jsonapi:"attr,reason"`
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
		ID:       generateULID(),
		Sparklee: matches[SparkleeIndex],
		Reason:   matches[ReasonIndex],
	}, nil
}

func generateULID() string {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	entropy := rand.New(source)
	id, _ := ulid.New(ulid.Timestamp(time.Now()), entropy)
	return id.String()
}
