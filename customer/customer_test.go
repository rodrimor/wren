package customer_test

import (
	"testing"

	. "github.com/rodrimor/wren/customer"
)

func TestSetAccount(t *testing.T) {
	var ac Account
	ac.SetStatusAccount("canceled", "Account canceled due to customer request")
	if !ac.AccountStatus.Canceled {
		t.Errorf("Expected status equals true, got false")
	}
	ac.AccountStatus.Canceled = false
	ac.SetStatusAccount("indebt", "indebt account")
	if !ac.AccountStatus.DebitNegotiation {
		t.Errorf("Expected status equals true, got false")
	}
	ac.AccountStatus.DebitNegotiation = false
	ac.SetStatusAccount("admin", "Account block due to the need of further documentation")
	if !ac.AccountStatus.Warning {
		t.Errorf("Expected Warning")
	}

}
