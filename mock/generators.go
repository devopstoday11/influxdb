package mock

import (
	"testing"
	"time"

	influxdb "github.com/influxdata/influxdb/v2"
)

// IDGenerator is mock implementation of influxdb.IDGenerator.
type IDGenerator struct {
	IDFn func() influxdb.ID
}

// ID generates a new influxdb.ID from a mock function.
func (g IDGenerator) ID() influxdb.ID {
	return g.IDFn()
}

// NewIDGenerator is a simple way to create immutable id generator
func NewIDGenerator(s string, t *testing.T) IDGenerator {
	t.Helper()

	id, err := influxdb.IDFromString(s)
	if err != nil {
		t.Fatal(err)
	}

	return NewStaticIDGenerator(*id)
}

// NewStaticIDGenerator returns an IDGenerator which produces the ID
// provided to this function on a call to ID().
func NewStaticIDGenerator(id influxdb.ID) IDGenerator {
	return IDGenerator{
		IDFn: func() influxdb.ID {
			return id
		},
	}
}

// NewIncrementingIDGenerator returns an ID generator which starts at the
// provided ID and increments on each call to ID().
func NewIncrementingIDGenerator(start influxdb.ID) IDGenerator {
	return IDGenerator{
		IDFn: func() influxdb.ID {
			defer func() { start++ }()
			return start
		},
	}
}

type MockIDGenerator struct {
	Last  *influxdb.ID
	Count int
}

const FirstMockID int = 65536

func NewMockIDGenerator() *MockIDGenerator {
	return &MockIDGenerator{
		Count: FirstMockID,
	}
}

func (g *MockIDGenerator) ID() influxdb.ID {
	id := influxdb.ID(g.Count)
	g.Count++

	g.Last = &id

	return id
}

// NewTokenGenerator is a simple way to create immutable token generator.
func NewTokenGenerator(s string, err error) TokenGenerator {
	return TokenGenerator{
		TokenFn: func() (string, error) {
			return s, err
		},
	}
}

// TokenGenerator is mock implementation of influxdb.TokenGenerator.
type TokenGenerator struct {
	TokenFn func() (string, error)
}

// Token generates a new influxdb.Token from a mock function.
func (g TokenGenerator) Token() (string, error) {
	return g.TokenFn()
}

// TimeGenerator stores a fake value of time.
type TimeGenerator struct {
	FakeValue time.Time
}

// Now will return the FakeValue stored in the struct.
func (g TimeGenerator) Now() time.Time {
	return g.FakeValue
}
