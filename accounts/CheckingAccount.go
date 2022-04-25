package accounts

import (
	"fmt"
	"time"

	"bank/clients"
)

type CheckingAccount struct {
	Owner         clients.Owner
	NumberAgency  int
	NumberAccount int
	amount        float64
	History       []History
}

func (c *CheckingAccount) Withdraw(amountWithdraw float64) History {
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

func (c *CheckingAccount) Deposit(amountDeposit float64) History {
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

func (c *CheckingAccount) Transfer(amountToSend float64, targetAccount *CheckingAccount) History {
	canSendMoney := amountToSend <= c.amount && amountToSend > 0

	if canSendMoney {
		c.amount -= amountToSend
		targetAccount.Deposit(amountToSend)

		return c.createEvent(History{
			description: "Transferência realizado com sucesso para " + targetAccount.Owner.Name + ".",
			amount:      amountToSend,
			isSuccess:   true,
			eventType:   Debit,
			dateAt:      time.Now(),
		})
	} else {
		return c.createEvent(History{
			description: "Erro em transferência para " + targetAccount.Owner.Name + ".",
			amount:      amountToSend,
			isSuccess:   true,
			eventType:   Debit,
			dateAt:      time.Now(),
		})
	}
}

func (c *CheckingAccount) GetAmount() float64 {
	return c.amount
}

func (c *CheckingAccount) GetHistory() []History {
	return c.History
}

func (c *CheckingAccount) PrintHistory() {
	for _, event := range c.History {
		fmt.Println(event.dateAt.Format("02-01-2006"), ":", event.description)
		fmt.Println(string(event.eventType), "-", event.amount)
		fmt.Println()
	}
}

func (c *CheckingAccount) createEvent(event History) History {
	c.History = append(c.History, event)
	return event
}
