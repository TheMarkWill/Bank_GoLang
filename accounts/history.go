package accounts

import "time"

type EventType string

const (
	Debit    EventType = "debit"
	Withdraw           = "withdraw"
	Transfer           = "transfer"
)

type History struct {
	description string
	amount      float64
	isSuccess   bool
	eventType   EventType
	dateAt      time.Time
}
