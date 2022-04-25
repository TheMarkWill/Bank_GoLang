package main

import (
	"bank/accounts"
	"bank/clients"
)

func PayBillet(account verifyAccount, amountBillet float64) {
	account.Withdraw(amountBillet)
}

type verifyAccount interface {
	Withdraw(amountWithdraw float64) accounts.History
}

func main() {
	clientSilvia := clients.Owner{
		Name:       "Silvia",
		CPF:        "1234",
		Occupation: "Designer",
	}

	accountSilvia := accounts.CheckingAccount{Owner: clientSilvia}
	accountSilvia.Deposit(100)

	PayBillet(&accountSilvia, 25)

	// fmt.Println(accountSilvia.GetAmount())
	// fmt.Println(accountSilvia.GetHistory())

	accountSilvia.PrintHistory()
}
