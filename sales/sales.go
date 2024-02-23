package sales

import "time"

type Products struct {
	ID         int
	Name       string
	price      float64
	LastUpdate time.Time
}
