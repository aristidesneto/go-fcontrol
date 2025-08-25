package services

import (
	"context"
	"errors"
	"fmt"
	"go-fcontrol-api/src/models"
	"go-fcontrol-api/src/repositories"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type TransactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		repo: *repositories.NewTransactionRepository(),
	}
}

func (s *TransactionService) GetTransactions(ctx context.Context, queryParam models.Transaction) ([]models.Transaction, error) {
	filter := bson.M{}

	// if queryParam.Name != "" {
	// 	filter = bson.M{
	// 		"$and": bson.A{
	// 			bson.M{"name": bson.M{"$regex": queryParam.Name, "$options": "i"}},
	// 		},
	// 	}
	// }

	return s.repo.GetTransactions(ctx, filter)
}

func (s *TransactionService) CreateTransaction(ctx context.Context, transaction models.Transaction) (*models.Transaction, error) {
	// 1. Validação básica
	if transaction.UserId.IsZero() {
		return nil, errors.New("user_id is required")
	}
	if transaction.Amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	if transaction.PaymentMethod == "" {
		return nil, errors.New("payment_method is required")
	}

	// 2. Pagamento e parcelas
	if transaction.PaymentMethod == "cash" && transaction.InstallmentCount > 1 {
		return nil, errors.New("cash payments cannot have installments")
	}

	// Se houver parcelas, calcula os valores e datas
	if transaction.InstallmentCount > 1 {

		if transaction.DueDate == nil {
			return nil, errors.New("due_date is required")
		}

		// Define a data base para as parcelas
		baseDate := *transaction.DueDate

		amountPerInstallment := transaction.Amount / float64(transaction.InstallmentCount)
		installments := make([]models.Installment, transaction.InstallmentCount)

		parcelas := gerarParcelas(baseDate, transaction.InstallmentCount)

		for i, p := range parcelas {
			fmt.Printf("Parcela %02d: %s\n", i+1, p.Format("2006-01-02"))
			installments[i] = models.Installment{
				Number:   i + 1,
				Amount:   amountPerInstallment,
				DueDate:  p,
				PaidDate: nil,
			}
		}

		transaction.Installments = installments
	}

	// 3. Recorrência
	// if transaction.IsRecurring {
	// 	if transaction.RecurrenceRule == nil {
	// 		return nil, errors.New("recurrence rule is required")
	// 	}
	// 	// apenas salva a regra, não gera todas no banco
	// }

	res, err := s.repo.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}
	slog.Info("Trsanction created", "id", res.InsertedID)

	transaction.ID = res.InsertedID.(bson.ObjectID)

	return &transaction, nil
}

// Função que soma meses respeitando fim de mês
func addMonthsEndOfMonth(t time.Time, months int) time.Time {
	year, month, day := t.Date()
	loc := t.Location()

	// cria a data "normal"
	newDate := time.Date(year, month+time.Month(months), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)

	// se "estourou" pro mês seguinte, corrige para último dia do mês esperado
	if newDate.Day() != day {
		newDate = time.Date(year, month+time.Month(months)+1, 0, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
	}

	return newDate
}

// Gera as parcelas
func gerarParcelas(inicio time.Time, qtd int) []time.Time {
	parcelas := make([]time.Time, qtd)
	for i := range parcelas {
		parcelas[i] = addMonthsEndOfMonth(inicio, i)
	}
	return parcelas
}
