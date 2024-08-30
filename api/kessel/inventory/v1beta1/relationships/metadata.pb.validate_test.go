package relationships

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetadataValid(t *testing.T) {
	meta := Metadata{}
	err := meta.ValidateAll()

	assert.NoError(t, err)
}


	err := meta.ValidateAll()
	assert.ErrorContains(t, err, "invalid Metadata.Labels[0]")
}
