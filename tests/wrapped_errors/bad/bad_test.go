package bad

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestControllerExecute(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	controller := Controller{}

	err := controller.Execute()
	require.Error(err)

	// fix me
	assert.Equal(
		err,
		errors.Wrap(
			errors.Wrap(
				errors.Wrap(
					errors.Wrap(ErrNotFound, StrRepositoryFail), StrServiceFail), StrProviderFail), StrControllerFail),
	)
}
