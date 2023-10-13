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
)

type Users interface {
	Create(c context.Context, dto schemas.CreateUser) (*models.User, error)
	List(c context.Context) ([]models.User, error)
	Get(c context.Context, id string) (*models.User, error)
	GetByEmail(c context.Context, email string) (*models.User, error)
	Delete(c context.Context, id string) error
	Update(c context.Context, id string, updateUserDto *schemas.UpdateUser) (*models.User, error)
}

type mongoUsers struct {
	collection *mongo.Collection
}

func NewUsers(dbClient *mongo.Client) *mongoUsers {
	return &mongoUsers{dbClient.Database("fibermongo").Collection("users")}
}

func (u *mongoUsers) Create(c context.Context, dto schemas.CreateUser) (*models.User, error) {
	dto.CreatedAt = time.Now()
	result, err := u.collection.InsertOne(c, dto)
	if err != nil {
		return nil, err
	}
	userId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("error casting inserted id")
	}
	return u.Get(c, userId.Hex())
}

func (u *mongoUsers) List(c context.Context) ([]models.User, error) {
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

func (u *mongoUsers) Get(c context.Context, id string) (*models.User, error) {
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

func (u *mongoUsers) GetByEmail(c context.Context, email string) (*models.User, error) {
	var user models.User
	if err := u.collection.FindOne(c, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *mongoUsers) Delete(c context.Context, id string) error {
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

func (u *mongoUsers) Update(c context.Context, id string, updateUserDto *schemas.UpdateUser) (*models.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	updateResult, err := u.collection.UpdateOne(c, bson.M{"_id": objectId}, bson.M{"$set": updateUserDto})
	if err != nil {
		return nil, err
	}
	if updateResult.ModifiedCount == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u.Get(c, id)
}
