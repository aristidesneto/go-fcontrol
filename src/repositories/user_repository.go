package repositories

// var userCollection = configs.GetCollection(configs.DB, "users")
// var userCollection = configs.MongoDatabase.Collection("users")

// func GetUsers() ([]models.UserResponse, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	cursor, err := userCollection.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var users []models.UserResponse
// 	if err := cursor.All(ctx, &users); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	res, err := userCollection.InsertOne(ctx, user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return res, nil
// }
