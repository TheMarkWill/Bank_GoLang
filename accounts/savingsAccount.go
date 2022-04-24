package accounts

import (
	"bank/clients"
	"time"
)

type SavingsAccount struct {
	Owner                                  clients.Owner
	NumberAgency, NumberAccount, Operation int
	amount                                 float64
	History                                []History
}

func (c *SavingsAccount) Withdraw(amountWithdraw float64) History {
	canWithdraw := amountWithdraw <= c.amount

	if canWithdraw {
		c.amount -= amountWithdraw

		return c.createEvent(History{
			description: "Saque realizado com sucesso",
			amount:      amountWithdraw,
			isSuccess:   true,
			eventType:   Withdraw,
			dateAt:      time.Now(),
		})
	} else {
		return c.createEvent(History{
			description: "Erro ao tentar realizar saque",
			amount:      amountWithdraw,
			isSuccess:   false,
			eventType:   Withdraw,
			dateAt:      time.Now(),
		})
	}
}

func (c *SavingsAccount) Deposit(amountDeposit float64) History {
	canDeposit := amountDeposit > 0

	if canDeposit {
		c.amount += amountDeposit

		return c.createEvent(History{
			description: "Deposito realizado com sucesso",
			amount:      amountDeposit,
			isSuccess:   true,
			eventType:   Debit,
			dateAt:      time.Now(),
		})
	} else {
		return c.createEvent(History{
			description: "Erro ao fazer deposito",
			amount:      amountDeposit,
			isSuccess:   false,
			eventType:   Debit,
			dateAt:      time.Now(),
		})
	}
}

func (c *SavingsAccount) GetAmount() float64 {
	return c.amount
}

func (c *SavingsAccount) createEvent(event History) History {
	c.History = append(c.History, event)
	return event
}
