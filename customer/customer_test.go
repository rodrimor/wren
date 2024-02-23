package customer_test

import (
	"testing"
	"time"

	. "github.com/rodrimor/wren/customer"
)

func TestLineItem(t *testing.T) {
	now := time.Now()
	item := LineItem{
		Created:       now,
		BusinessName:  "Example Store, inc",
		Amount:        9.99,
		TransactionId: 123456,
	}

	// Test created time
	if item.Created != now {
		t.Errorf("Expected created time to be %v, got %v", now, item.Created)
	}

	if item.BusinessName != "Example Store, inc" {
		t.Errorf("Expected Business name %s, got %s", "Example Store, inc", item.BusinessName)
	}
}
