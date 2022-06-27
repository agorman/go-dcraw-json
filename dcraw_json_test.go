package dcraw

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	data, err := GetImageData(context.Background(), "./testdata/RAW_KODAK_DC50.KDC")
	assert.Nil(t, err)

	assert.Equal(t, "./testdata/RAW_KODAK_DC50.KDC", data.Filename)
	assert.Equal(t, int64(861782447), data.Timestamp)
	assert.Equal(t, 0, data.ISO)
	assert.Equal(t, 0.00435, data.Shutter)
	assert.Equal(t, 6.2, data.Aperture)
	assert.Equal(t, 37.0, data.Focal)

	assert.Equal(t, "Kodak", data.Camera.Make)
	assert.Equal(t, "DC50", data.Camera.Model)

	assert.Equal(t, uint(96), data.ThumbSize.Width)
	assert.Equal(t, uint(64), data.ThumbSize.Height)

	assert.Equal(t, uint(768), data.FullSize.Width)
	assert.Equal(t, uint(512), data.FullSize.Height)

	assert.Nil(t, data.GPS)
}
