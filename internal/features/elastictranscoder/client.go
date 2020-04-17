//Package elastictranscoder provides gucumber integration tests suppport.
package elastictranscoder

import (
	"github.com/daysleep666/aws-sdk-go/internal/features/shared"
	"github.com/daysleep666/aws-sdk-go/service/elastictranscoder"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@elastictranscoder", func() {
		World["client"] = elastictranscoder.New(nil)
	})
}
