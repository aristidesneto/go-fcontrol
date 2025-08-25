package repositories

import (
	"context"
	"go-fcontrol-api/src/configs"
	"go-fcontrol-api/src/models"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TransactionRepository struct {
	collection *mongo.Collection
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		collection: configs.MongoDatabase.Collection("transactions"),
	}
}

func (r *TransactionRepository) GetTransactions(ctx context.Context, filter bson.M) ([]models.Transaction, error) {
	slog.Debug("Searching transactions", "filter", filter)
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	users := make([]models.Transaction, 0)
	if err := cursor.All(ctx, &users); err != nil {
		return make([]models.Transaction, 0), err
	}

	return users, nil
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction models.Transaction) (*mongo.InsertOneResult, error) {
	slog.Debug("Creating transaction")
	res, err := r.collection.InsertOne(ctx, transaction)
	if err != nil {
		slog.Debug("Error to create transaction", "error", err)
		return nil, err
	}

	return res, nil
}
