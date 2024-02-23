/*  2022 The wren authors.
*
*   All rights reserved use this source of code is governed by a BSD
*   license check licence file.
*
*  Package wren makes the most activities of the credit card
*  it includes, sells, account blocks, orders permissions, debit -
*  calculator.
 */
package wren

import (
	"fmt"

	"github.com/rodrimor/wren/customer"
)

// Sells contains its value, approved
type SaleInfo struct {
	value     float64
	Approved  bool
	ErrStatus error
}

// Order returns whether the sell was approved or not
func (s *SaleInfo) Order(customerId customer.Account, value float64) *SaleInfo {

	// Checking Account and Card status
	err, check := AccountAndCardCheck(customerId)

	if (customerId.Limite - customerId.Debt) < int64(value) {
		// checks whether the shopping items
		//values more than customers limite
		// When it exceeds this might be refused instaintly

		s.Approved = false

		s.ErrStatus = fmt.Errorf("%d Exceeded limit", 51)
		return s

	} else if check == false {

		s.ErrStatus = fmt.Errorf("%d something went wrong", err)

		return s
	}

	return s
}

// AccountAndCardCheck returns whether a error is found or not.
func AccountAndCardCheck(customerId customer.Account) (int, bool) {

	// checking accounts blocks if any of these are true the
	// order will be denied.
	switch {
	case customerId.Statuses.UnSet:
		return 13, false
	case customerId.Statuses.Adm:
		return 14, false
	case customerId.Statuses.Canceled:
		return 15, false
	case customerId.Statuses.DebitinNovation:
		return 16, false
	}

	// checking Card blocks if any of these are true the
	// order will be denied.

	switch {
	case customerId.Credit_card.Status.Canceled:
		return 46, false
	case customerId.Credit_card.Status.Misplaced:
		return 50, false
	case customerId.Credit_card.Status.Renewed:
		return 51, false
	case customerId.Credit_card.Status.Shipping:
		return 46, false
	}

	// if none of those blocks above is true, the check is completed and approved.
	return 0, true
}
