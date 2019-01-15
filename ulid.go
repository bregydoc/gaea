package gaea

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

var t = time.Unix(1000000, 0)
var entropy = ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)

func newULID() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}
