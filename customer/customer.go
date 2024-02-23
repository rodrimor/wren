// Copyright 2024 Mateus Rodrigues Moreira github.com/rodrimor
//
// Licensed under the BSD-3-Clause License (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/BSD-3-Clause
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package customer defines types and functions for managing customer accounts.
//
// This package provides a basic implementation of customer accounts, including
// fields for name, ID, credit limit, current debt, and account statuses.
//
// More features and functionalities can be added to this package in the future.

package customer

import (
	"fmt"
	"time"
)

type Account struct {
	// **ID**: Unique identifier for the customer account. This field is a primary
	// key in the database and cannot be duplicated.
	ID int64

	// **Name**: Full name of the account holder. Use the customer's legal name
	// as it appears on identification documents.
	Name string

	// **Limit**: Maximum amount of credit allowed on the account. This value
	// cannot be exceeded by transactions.
	Limit float64

	// **AccountStatus**: Comma-separated list of current statuses associated with the
	// account. These statuses might indicate account activity, restrictions,
	// or other relevant information.
	AccountStatus Status
}

// Status represents the status of an account or a credit card.
type Status struct {
	// UnSet indicates whether the account or credit card is still not active.
	// Both SetStatus function for cards and account attribute false automatically to this property.
	UnSet bool

	// Warning indicates an account block with a message edited by the system's administrator.
	Warning bool

	// Canceled indicates whether the account has been canceled.
	Canceled bool

	// DebitNegotiation indicates a status where the customer is negotiating a debit.
	// It represents an account block.
	DebitNegotiation bool

	// Message represents a message associated with the status.
	Message string

	// Deleted indicates whether the credit card has been deleted.
	Deleted bool

	// Loss indicates a card block due to loss.
	Loss bool
	// Created represents when a status was set.
	Created time.Time
}

// SetStatusAccount updates the status of an account.
// Automatically sets the Unset status as false and them
// set the proper status accourding to a message and status
func (acs *Account) SetStatusAccount(newstatus, msg string) {
	// automatically unset false
	acs.AccountStatus.UnSet = false

	// if stataments for the new status.
	if newstatus == "indebt" {
		// Status indebt when customers are negotiating their debitts
		// this status cannot be changed on AccountUnset function until
		// the account's debit is equals zero.
		acs.AccountStatus.DebitNegotiation = true
		acs.AccountStatus.Message = msg

	} else if newstatus == "canceled" {
		// canceled account can unset and set anytime the customers want,
		// AccountUnset function won't allow the return of the customer only
		// when the credit card limit is equals zero.
		acs.AccountStatus.Canceled = true
		acs.AccountStatus.Message = msg

	} else if newstatus == "admin" {
		// Admin block can be set whenever a customer asks for it or due to the
		// need of updates and further documentation.
		acs.AccountStatus.Message = msg
		acs.AccountStatus.Warning = true
	}
	acs.AccountStatus.Created = time.Now()
	fmt.Printf("New Status created at %v", acs.AccountStatus.Created)
}
