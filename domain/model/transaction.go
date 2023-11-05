package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(trasaction *Transaction) error
	Save(trasaction *Transaction) error
	Find(trasaction *Transaction) error
}
type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id; type:uuid;" valid:"notnull"`
	Amount            float64  `json:"amount" valid:"notnull" gorm:"type:float"`
	PixkeyTo          *Pixkey  `valid:"-"`
	PixkeyIdTo        string   `gorm:"column:pix_key_id_to;type:uuid;" valid:"notnull"`
	Status            string   `json:"status" valid:"notnull" gorm:"type:varchar(20)"`
	Description       string   `json:"description"  gorm:"type:varchar(255)" valid:"-"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

func (transaction *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)

	if transaction.Amount <= 0 {
		return errors.New("the amount must be greater tha 0")
	}

	if transaction.Status != TransactionPending && transaction.Status != TransactionCompleted && transaction.Status != TransactionError && transaction.Status != TransactionPending {
		return errors.New("invalid status for the transaction")
	}

	if transaction.PixkeyTo.AccountID == transaction.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")

	}

	if err != nil {
		return err
	}

	return nil

}

func NewTransaction(accountfrom *Account, amount float64, pixKeyTo *Pixkey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountfrom,
		Amount:      amount,
		PixkeyTo:    pixKeyTo,
		Status:      TransactionPending,
		Description: description,
	}
	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (trasaction *Transaction) Complete() error {
	trasaction.Status = TransactionCompleted
	trasaction.UpdatedAt = time.Now()

	err := trasaction.isValid()

	return err
}

func (trasaction *Transaction) Cancel(description string) error {
	trasaction.Status = TransactionError
	trasaction.UpdatedAt = time.Now()
	trasaction.Description = description
	err := trasaction.isValid()

	return err
}

func (trasaction *Transaction) Confirm() error {
	trasaction.Status = TransactionConfirmed
	trasaction.UpdatedAt = time.Now()

	err := trasaction.isValid()

	return err
}
