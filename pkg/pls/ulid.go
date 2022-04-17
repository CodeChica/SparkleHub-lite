package pls

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateULID() string {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	entropy := rand.New(source)
	id, _ := ulid.New(ulid.Timestamp(time.Now()), entropy)
	return id.String()
}
