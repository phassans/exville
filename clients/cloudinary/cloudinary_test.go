package cloudinary

import (
	"io"
	"strings"
	"testing"

	"github.com/phassans/exville/common"
	"github.com/stretchr/testify/require"
)

func TestClient_Upload(t *testing.T) {
	cloudinaryClient := NewCloudinaryClient(common.GetLogger())
	//prepare the reader instances to encode
	values := map[string]io.Reader{
		"file":          mustOpen("../../upload_images/IMG_9614.JPG"), // lets assume its this file
		"upload_preset": strings.NewReader(UPLOAD_PRESET),
		//"public_id":     strings.NewReader("test"),
		//"folder":        strings.NewReader("upload_images"),
	}
	err := cloudinaryClient.Upload(values)
	require.NoError(t, err)
}
