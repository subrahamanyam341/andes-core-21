package factory

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-core-21/core/check"
	"github.com/subrahamanyam341/andes-core-21/marshal"
)

func TestNewMarshalizer_UnknownTypeShouldErr(t *testing.T) {
	t.Parallel()

	mrs, err := NewMarshalizer("unknown")

	assert.True(t, check.IfNil(mrs))
	assert.True(t, errors.Is(err, marshal.ErrUnknownMarshalizer))
}

func TestNewMarshalizer_JsonShouldWork(t *testing.T) {
	t.Parallel()

	mrs, err := NewMarshalizer(JsonMarshalizer)

	jsonMrs := (*marshal.JsonMarshalizer)(nil)
	assert.Nil(t, err)
	assert.IsType(t, jsonMrs, mrs)
}

func TestNewMarshalizer_GogoPotobufShouldWork(t *testing.T) {
	t.Parallel()

	mrs, err := NewMarshalizer(GogoProtobuf)

	protoMrs := (*marshal.GogoProtoMarshalizer)(nil)
	assert.Nil(t, err)
	assert.IsType(t, protoMrs, mrs)
}
