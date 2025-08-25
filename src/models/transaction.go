package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type TransactionType string

const (
	IncomeTransaction   TransactionType = "income"   // Receita
	ExpenseTransaction  TransactionType = "expense"  // Despesa
	TransferTransaction TransactionType = "transfer" // Transferência entre contas
)

type PaymentMethod string

const (
	CashPayment       PaymentMethod = "cash"          // Dinheiro
	BankTransfer      PaymentMethod = "bank_transfer" // Transferencia bancária
	CreditCardPayment PaymentMethod = "credit_card"   // Cartão de Crédito
	DebitCardPayment  PaymentMethod = "debit_card"    // Cartão de débito
	PixPayment        PaymentMethod = "pix"           // PIX
)

// type RecurrenceType string

// const (
// 	NoneRecurrence    RecurrenceType = "none"    // Lançamento único
// 	MonthlyRecurrence RecurrenceType = "monthly" // Fixa mensal
// 	WeeklyRecurrence  RecurrenceType = "weekly"  // Ex: assinatura semanal
// 	YearlyRecurrence  RecurrenceType = "yearly"  // Ex: seguro anual
// )

type Installment struct {
	Number   int        `form:"number" bson:"number" json:"number" binding:"required,gt=0,number"` // Nº da parcela (1, 2, 3...)
	DueDate  time.Time  `form:"due_date" bson:"due_date" json:"due_date"`                          // Data de vencimento
	Amount   float64    `form:"amount" bson:"amount" json:"amount"`                                // Valor da parcela
	PaidDate *time.Time `form:"paid_date" bson:"paid_date" json:"paid_date" binding:"datetime"`    // Data em que foi paga (se houver)
}

type Transaction struct {
	ID         bson.ObjectID `form:"id" bson:"_id,omitempty" json:"id"`
	UserId     bson.ObjectID `form:"user_id" bson:"user_id,omitempty" json:"user_id"`
	CategoryID bson.ObjectID `form:"category_id" bson:"category_id,omitempty" json:"category_id"`

	Type            TransactionType `form:"type" bson:"type" json:"type" binding:"oneof=income expense"`      // Receita ou despesa
	Amount          float64         `form:"amount" bson:"amount" json:"amount" binding:"gt=0"`                // Valor total da despesa
	Description     string          `form:"description" bson:"description" json:"description"`                // Descrição
	PaymentMethod   PaymentMethod   `form:"payment_method" bson:"payment_method" json:"payment_method"`       // Forma de pagamento
	TransactionDate time.Time       `form:"transaction_date" bson:"transaction_date" json:"transaction_date"` // Data de vencimento
	DueDate         *time.Time      `form:"due_date" bson:"due_date" json:"due_date"`                         // útil para cartões / boletos
	Paid            *time.Time      `form:"paid" bson:"paid" json:"paid"`                                     // Data de pagamento

	// Recorrência
	IsRecurring   bool       `form:"is_recurring" bson:"is_recurring" json:"is_recurring" binding:"boolean"`         // Lançamento recorrente
	RecurrenceEnd *time.Time `form:"recurrence_end" bson:"recurrence_end,omitempty" json:"recurrence_end,omitempty"` // Data fim da recorrência

	// Parcelamento
	// IsInstallment    bool          `form:"is_installment" bson:"is_installment" json:"is_installment"`                              // Se é parcelado
	InstallmentCount int           `form:"installment_count" bson:"installment_count,omitempty" json:"installment_count,omitempty"` // total de parcela
	Installments     []Installment `form:"installments" bson:"installments,omitempty" json:"installments,omitempty"`                // Parcelas

	// Metadata
	CreatedAt time.Time  `form:"created_at" bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time `form:"updated_at" bson:"updated_at" json:"updated_at"`
}
