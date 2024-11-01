package pgimpl

import (
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // we import the driver used by sqlx
	"github.com/stretchr/testify/assert"
)

func TestPingNoDb(t *testing.T) {
	impl := pgFlagImpl{}
	assert.Error(t, impl.Ping())
}
