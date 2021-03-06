package remote

import (
	"fmt"

	"github.com/cloudskiff/driftctl/pkg/alerter"
	"github.com/cloudskiff/driftctl/pkg/remote/aws"
)

var supportedRemotes = []string{
	aws.RemoteAWSTerraform,
}

func IsSupported(remote string) bool {
	for _, r := range supportedRemotes {
		if r == remote {
			return true
		}
	}
	return false
}

func Activate(remote string, alerter *alerter.Alerter) error {
	switch remote {
	case aws.RemoteAWSTerraform:
		return aws.Init(alerter)
	default:
		return fmt.Errorf("unsupported remote '%s'", remote)
	}
}

func GetSupportedRemotes() []string {
	return supportedRemotes
}
