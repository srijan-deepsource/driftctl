package remote

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/cloudskiff/driftctl/pkg/alerter"
	"github.com/cloudskiff/driftctl/pkg/remote/alerts"
	"github.com/cloudskiff/driftctl/pkg/remote/common"
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleResourceEnumerationError(err error, alerter alerter.AlerterInterface) error {
	listError, ok := err.(*remoteerror.ResourceScanningError)
	if !ok {
		return err
	}

	rootCause := listError.RootCause()

	// We cannot use status.FromError() method since AWS errors are weak since they compose error interface
	// without implementing Error() method and thus this trigger a nil panic when returning an unknown error from
	// status.FromError()
	// As a workaround we duplicated the logic from status.FromError here
	if _, ok := rootCause.(interface{ GRPCStatus() *status.Status }); ok {
		return handleGoogleEnumerationError(alerter, listError, status.Convert(rootCause))
	}

	reqerr, ok := rootCause.(awserr.RequestFailure)
	if ok {
		return handleAWSError(alerter, listError, reqerr)
	}

	// This handles access denied errors like the following:
	// aws_s3_bucket_policy: AccessDenied: Error listing bucket policy <policy_name>
	if strings.Contains(rootCause.Error(), "AccessDenied") {
		alerts.SendEnumerationAlert(common.RemoteAWSTerraform, alerter, listError)
		return nil
	}

	if strings.HasPrefix(
		rootCause.Error(),
		"Your token has not been granted the required scopes to execute this query.",
	) {
		alerts.SendEnumerationAlert(common.RemoteGithubTerraform, alerter, listError)
		return nil
	}

	return err
}

func HandleResourceDetailsFetchingError(err error, alerter alerter.AlerterInterface) error {
	listError, ok := err.(*remoteerror.ResourceScanningError)
	if !ok {
		return err
	}

	rootCause := listError.RootCause()

	if shouldHandleGoogleDetailsFetchingError(listError) {
		alerts.SendDetailsFetchingAlert(common.RemoteGoogleTerraform, alerter, listError)
		return nil
	}

	// This handles access denied errors like the following:
	// iam_role_policy: error reading IAM Role Policy (<policy>): AccessDenied: User: <role_arn> ...
	if strings.HasPrefix(rootCause.Error(), "AccessDeniedException") ||
		strings.Contains(rootCause.Error(), "AccessDenied") ||
		strings.Contains(rootCause.Error(), "AuthorizationError") {
		alerts.SendDetailsFetchingAlert(common.RemoteAWSTerraform, alerter, listError)
		return nil
	}

	return err
}

func handleAWSError(alerter alerter.AlerterInterface, listError *remoteerror.ResourceScanningError, reqerr awserr.RequestFailure) error {
	if reqerr.StatusCode() == 403 || (reqerr.StatusCode() == 400 && strings.Contains(reqerr.Code(), "AccessDenied")) {
		alerts.SendEnumerationAlert(common.RemoteAWSTerraform, alerter, listError)
		return nil
	}

	return reqerr
}

func handleGoogleEnumerationError(alerter alerter.AlerterInterface, err *remoteerror.ResourceScanningError, st *status.Status) error {
	if st.Code() == codes.PermissionDenied {
		alerts.SendEnumerationAlert(common.RemoteGoogleTerraform, alerter, err)
		return nil
	}
	return err
}

func shouldHandleGoogleDetailsFetchingError(err *remoteerror.ResourceScanningError) bool {
	errMsg := err.RootCause().Error()

	// Check if this is a Google related error
	if !strings.Contains(errMsg, "googleapi") {
		return false
	}

	if strings.Contains(errMsg, "Error 403") {
		return true
	}

	return false
}
