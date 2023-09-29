package main

import (
	"errors"
	"log"
)

// Models an account in a bank
type Account struct {
	id           uint8
	name         string
	balance      uint32
	transactions []Transaction
}

// All of the possible payment methods
type PaymentMethod string

const (
	CREDIT PaymentMethod = "C"
	DEBIT  PaymentMethod = "D"
	CASH   PaymentMethod = "S"
)

type TransactionState string

const (
	OPEN    TransactionState = "O"
	EXPIRED TransactionState = "E"
	CLOSED  TransactionState = "C"
)

// Models the transaction one account can make to another
type Transaction struct {
	id                 uint8
	value              uint32
	sender             *Account
	recipient          *Account
	state              TransactionState
	paymentMethod      PaymentMethod
	transactionHandler TransactionHandler
}

// Interface for handling paying transactions
type TransactionHandler interface {
	// Pays an open transaction
	// Returns an error if the transaction is invalid
	pay(t *Transaction) error
}

// Pays transaction
func (t *Transaction) makePayment() error {
	err := t.transactionHandler.pay(t)

	if err != nil {
		return err
	}

	return nil
}

// Chooses what handler should be used with each transaction
func (t *Transaction) selectTransactionHandler() error {
	switch t.paymentMethod {
	case CREDIT:
		t.transactionHandler = &CreditTransactionHandler{}
		return nil
	case CASH:
		t.transactionHandler = &CashTransactionHandler{}
		return nil
	case DEBIT:
		t.transactionHandler = &DebitTransactionHandler{}
		return nil
	default:
		return errors.New("Could find a valid handler")
	}
}

// Models dependencies used to pay a transaction of type credit
type CreditTransactionHandler struct{}

// Handles transactions of type credit
func (th *CreditTransactionHandler) pay(t *Transaction) error {
	if t.sender.id == t.recipient.id {
		return errors.New("One account can't make a transaction to itself")
	}

	if t.state == CLOSED {
		return errors.New("Can't pay an already closed transaction")
	}

	if t.state == EXPIRED {
		return errors.New("Transaction expired")
	}

	if t.sender.balance < t.value {
		return errors.New("Sender doesn't have enough balance to make transaction")
	}

	t.sender.balance -= t.value
	t.recipient.balance += t.value
	t.state = CLOSED

	return nil
}

// Models dependencies used to pay a transaction of type cash
type CashTransactionHandler struct{}

// Handles transactions of type cash
func (th *CashTransactionHandler) pay(t *Transaction) error {
	if t.sender.id == t.recipient.id {
		return errors.New("One account can't make a transaction to itself")
	}

	if t.state == CLOSED {
		return errors.New("Can't pay an already closed transaction")
	}

	if t.state == EXPIRED {
		return errors.New("Transaction expired")
	}

	if t.sender.balance < t.value {
		return errors.New("Sender doesn't have enough balance to make transaction")
	}

	t.sender.balance -= t.value
	t.recipient.balance += t.value
	t.state = CLOSED

	return nil
}

// Models dependencies used to pay a transaction of type debit
type DebitTransactionHandler struct{}

// Handles transactions of type debit
func (th *DebitTransactionHandler) pay(t *Transaction) error {
	if t.sender.id == t.recipient.id {
		return errors.New("One account can't make a transaction to itself")
	}

	if t.state == CLOSED {
		return errors.New("Can't pay an already closed transaction")
	}

	if t.state == EXPIRED {
		return errors.New("Transaction expired")
	}

	if t.sender.balance < t.value {
		return errors.New("Sender doesn't have enough balance to make transaction")
	}

	t.sender.balance -= t.value
	t.recipient.balance += t.value
	t.state = CLOSED

	return nil
}

func main() {
	gustavo := &Account{
		id:      1,
		name:    "My first account",
		balance: 150,
	}

	pedro := &Account{
		id:      2,
		name:    "Online store",
		balance: 5,
	}

	transaction := &Transaction{
		id:            1,
		value:         55,
		sender:        gustavo,
		recipient:     pedro,
		state:         OPEN,
		paymentMethod: CASH,
	}

	err := transaction.selectTransactionHandler()

	if err != nil {
		log.Println(err)
	}

	err = transaction.makePayment()

	if err != nil {
		log.Println(err)
	}
}
