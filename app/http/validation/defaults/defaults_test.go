package defaults

import (
	"testing"

	"github.com/jadevelopmentgrp/Tickets-Import-API/utils"
	"github.com/stretchr/testify/assert"
)

func TestNil(t *testing.T) {
	var myString *string
	ApplyDefaults(NewDefaultApplicator[*string](NilCheck[string], &myString, utils.Ptr("hello")))
	assert.Equal(t, "hello", *myString)
}
