package gaea

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

var t = time.Unix(1000000, 0)
var entropy = ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)

func newULID() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}
