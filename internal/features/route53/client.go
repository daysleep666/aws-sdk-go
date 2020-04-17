//Package route53 provides gucumber integration tests suppport.
package route53

import (
	"github.com/daysleep666/aws-sdk-go/internal/features/shared"
	"github.com/daysleep666/aws-sdk-go/service/route53"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@route53", func() {
		World["client"] = route53.New(nil)
	})
}
