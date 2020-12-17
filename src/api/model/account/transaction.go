package account

import "time"

type Transaction struct {
	Id string
	Type string
	Amount float64
	EffectiveDate time.Time
}