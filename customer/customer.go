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

	// **Statuses**: Comma-separated list of current statuses associated with the
	// account. These statuses might indicate account activity, restrictions,
	// or other relevant information.
	Statuses string
}
