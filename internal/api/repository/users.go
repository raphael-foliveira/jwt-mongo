package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/raphael-foliveira/fiber-mongo/internal/api/models"
	"github.com/raphael-foliveira/fiber-mongo/internal/api/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Users interface {
	Create(c context.Context, dto schemas.CreateUser) (*models.User, error)
	List(c context.Context) ([]models.User, error)
	Get(c context.Context, id string) (*models.User, error)
	GetByEmail(c context.Context, email string) (*models.User, error)
	Delete(c context.Context, id string) error
	Update(c context.Context, user *models.User) (*models.User, error)
}

type usersMongo struct {
	collection *mongo.Collection
}

func NewUsers(dbClient *mongo.Client) *usersMongo {
	collection := dbClient.Database("fibermongo").Collection("users")
	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "email", Value: 1},
				{Key: "username", Value: 1},
			},
			Options: options.Index().SetUnique(true)})
	if err != nil {
		panic(err)
	}
	return &usersMongo{collection}
}

func (u *usersMongo) Create(c context.Context, dto schemas.CreateUser) (*models.User, error) {
	dto.CreatedAt = time.Now()
	result, err := u.collection.InsertOne(c, dto)
	if err != nil {
		return nil, err
	}
	userId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("error casting inserted id")
	}
	return &models.User{
		ID:        userId.Hex(),
		Username:  dto.Username,
		Email:     dto.Email,
		CreatedAt: dto.CreatedAt,
		Password:  dto.Password,
		Token:     "",
	}, nil
}

func (u *usersMongo) List(c context.Context) ([]models.User, error) {
	var users []models.User
	cursor, err := u.collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(c, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *usersMongo) Get(c context.Context, id string) (*models.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user models.User
	if err := u.collection.FindOne(c, bson.M{"_id": objectId}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usersMongo) GetByEmail(c context.Context, email string) (*models.User, error) {
	var user models.User
	if err := u.collection.FindOne(c, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usersMongo) Delete(c context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := u.collection.DeleteOne(c, bson.M{"_id": objectId})
	if result.DeletedCount == 0 {
		return fmt.Errorf("user not found")
	}
	return err
}

func (u *usersMongo) Update(c context.Context, user *models.User) (*models.User, error) {
	objectId, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return nil, err
	}
	updateResult, err := u.collection.UpdateOne(c, bson.M{"_id": objectId}, bson.M{"$set": schemas.UpdateUser{
		Password: user.Password,
		Token:    user.Token,
	}})
	if err != nil {
		return nil, err
	}
	if updateResult.ModifiedCount == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}
