package helpers

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/xmatters/xmatters-go"
)

func ResourceErrorDiagnostic(resource string, err error) diag.Diagnostic {
	var xmErr xmatters.XMattersError
	errors.As(err, &xmErr)
	return diag.NewErrorDiagnostic(
		fmt.Sprintf("Error creating or updating xMatters %s.", resource),
		fmt.Sprintf("xMatters API Error: %d - %s.\n  %s", xmErr.Code, xmErr.Reason, xmErr.Message),
	)
}

func DatasourceErrorDiagnostic(datasource string, err error) diag.Diagnostic {
	var xmErr xmatters.XMattersError
	errors.As(err, &xmErr)
	return diag.NewErrorDiagnostic(
		fmt.Sprintf("Error getting xMatters %s.", datasource),
		fmt.Sprintf("xMatters API Error: %d - %s.\n  %s", xmErr.Code, xmErr.Reason, xmErr.Message),
	)
}

func DeleteErrorDiagnostic(resource string, err error) diag.Diagnostic {
	var xmErr xmatters.XMattersError
	errors.As(err, &xmErr)
	return diag.NewErrorDiagnostic(
		fmt.Sprintf("Error deleting xMatters %s.", resource),
		fmt.Sprintf("xMatters API Error: %d - %s.\n  %s", xmErr.Code, xmErr.Reason, xmErr.Message),
	)
}
