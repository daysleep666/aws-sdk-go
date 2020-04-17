package s3

import "github.com/daysleep666/aws-sdk-go/aws"

func init() {
	initService = func(s *aws.Service) {
		// Support building custom host-style bucket endpoints
		s.Handlers.Build.PushFront(updateHostWithBucket)

		// Require SSL when using SSE keys
		s.Handlers.Validate.PushBack(validateSSERequiresSSL)
		s.Handlers.Build.PushBack(computeSSEKeys)

		// S3 uses custom error unmarshaling logic
		s.Handlers.UnmarshalError.Clear()
		s.Handlers.UnmarshalError.PushBack(unmarshalError)
	}

	initRequest = func(r *aws.Request) {
		switch r.Operation {
		case opPutBucketCORS, opPutBucketLifecycle, opPutBucketPolicy, opPutBucketTagging, opDeleteObjects, opPutObject, opUploadPart:
			// These S3 operations require Content-MD5 to be set
			// 这里不强制校验了...因为我们的seeker是假的
			// 无法通过seeker将数据都读到内存里
			// r.Handlers.Build.PushBack(contentMD5)
		case opGetBucketLocation:
			// GetBucketLocation has custom parsing logic
			r.Handlers.Unmarshal.PushFront(buildGetBucketLocation)
		case opCreateBucket:
			// Auto-populate LocationConstraint with current region
			r.Handlers.Validate.PushFront(populateLocationConstraint)
		}
	}
}
