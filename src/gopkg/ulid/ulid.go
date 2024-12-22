package ulid

import (
	"crypto/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

var randPool = sync.Pool{
	New: func() any {
		return ulid.Monotonic(rand.Reader, 0)
	},
}

func New() string {
	entrophy := randPool.Get().(*ulid.MonotonicEntropy)
	randPool.Put(entrophy)

	return ulid.MustNew(ulid.Timestamp(time.Now()), entrophy).String()
}
