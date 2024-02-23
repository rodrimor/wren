package customer

import "time"

type Card struct {
	// **ID**: Unique identifier for the credit card, assigned by the issuer.
	// This field is a primary key in the database and cannot be duplicated.
	ID int64

	// **HolderName**: Full name of the cardholder, copied from the `Name` field
	//  of the associated `Account`. Any modifications to the account holder's name
	//  should be reflected here.
	HolderName string

	// **Additional**: Indicates whether this card is an additional card linked
	// to the main account. Additional cards might have different limits, permissions
	// or benefits.
	Additional bool

	// **Status**: Current status of the card. Possible values might include "Active",
	// "Suspended", "Lost", "Stolen", etc.
	Status string
	// **Debt**: Current outstanding balance on the account. This is the sum of
	// all transactions that have not been paid yet.
	Debt float64

	Purchase []LineItem
}

// LineItem stores the name and time created of an specific Purchase
type LineItem struct {
	// Created: Timestamp representing the time when the item was bought.
	// This field uses the time.Time type for accurate time representation.
	// You can access the formatted time using the FormarttedTime function.
	Created time.Time
	// BusinessName: The name of the Business where the item was purchased.
	// This field stores a name as a string.
	BusinessName string
	// Amount: Total Amount paid for the item.
	// This fields uses a float64 to represent the amount with decimals.
	Amount float64
	// transactionId: Unique identifier for the transaction, uses an int64.
	TransactionId int64
}

// FormattedTime formats the created time of a Purchase as HH:MM:SS
func (li *LineItem) FormattedTime() string {
	return li.Created.Format("003:004:005")
}
