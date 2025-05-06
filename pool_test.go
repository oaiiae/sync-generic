package sync_test

import (
	"slices"
	"testing"
	"time"

	"github.com/oaiiae/sync-generic"
	"github.com/stretchr/testify/require"
)

func TestPool(t *testing.T) {
	t.Run("returns zero value", func(t *testing.T) {
		p := new(sync.Pool[[]byte])
		require.Equal(t, []byte(nil), p.Get())
		p = p.New(nil)
		require.Equal(t, []byte(nil), p.Get())
	})

	t.Run("returns new value", func(t *testing.T) {
		p := new(sync.Pool[[]byte]).New(func() []byte { return make([]byte, 8) })
		x := p.Get()
		require.NotEqual(t, []byte(nil), x)
		require.Len(t, x, 8)
	})

	t.Run("returns existing value", func(t *testing.T) {
		p := new(sync.Pool[[]byte])
		require.Equal(t, []byte(nil), p.Get())
		x := make([]byte, 8)
		p.Put(x)
		require.Eventually(t, func() bool { return slices.Compare(p.Get(), x) == 0 }, 10*time.Second, time.Second/10)
		require.Equal(t, []byte(nil), p.Get())
	})
}
