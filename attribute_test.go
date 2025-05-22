package otelx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint(t *testing.T) {
	kv := Uint("negative", 9223372036854775800)
	assert.EqualValues(t, 9223372036854775800, kv.Value.AsInt64())

	kv = Uint("negative", 9223372036854775880)
	assert.EqualValues(t, -9223372036854775736, kv.Value.AsInt64())
}

func TestUint64(t *testing.T) {
	kv := Uint64("negative", 9223372036854775800)
	assert.EqualValues(t, 9223372036854775800, kv.Value.AsInt64())

	kv = Uint64("negative", 9223372036854775880)
	assert.EqualValues(t, -9223372036854775736, kv.Value.AsInt64())
}
