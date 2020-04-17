//Package cognitosync provides gucumber integration tests suppport.
package cognitosync

import (
	"github.com/daysleep666/aws-sdk-go/internal/features/shared"
	"github.com/daysleep666/aws-sdk-go/service/cognitosync"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cognitosync", func() {
		World["client"] = cognitosync.New(nil)
	})
}
