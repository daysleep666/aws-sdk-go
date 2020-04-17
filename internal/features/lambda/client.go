//Package lambda provides gucumber integration tests suppport.
package lambda

import (
	"github.com/daysleep666/aws-sdk-go/internal/features/shared"
	"github.com/daysleep666/aws-sdk-go/service/lambda"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@lambda", func() {
		World["client"] = lambda.New(nil)
	})
}
