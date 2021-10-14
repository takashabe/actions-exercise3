package app_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takashabe/actions-exercise3/src/app"
)

func TestRun(t *testing.T) {
	err := app.Run()
	assert.NoError(t, err)
}
