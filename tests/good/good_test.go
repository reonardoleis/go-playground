package good

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestControllerExecute(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	controller := Controller{}

	err := controller.Execute()
	require.Error(err)
	assert.Equal(ErrNotFound, err)
}
