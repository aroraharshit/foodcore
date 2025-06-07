package repositories

import (
	"context"

	"github.com/aroraharshit/go-foodcore/user-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.RegistertUserInsertion) (string, error)
	FindByEmail(ctx context.Context, email string) (bool, error)
}

type userRepositoryOpts struct {
	userCollection *mongo.Collection
}

type userRepository struct {
	opts userRepositoryOpts
}

func NewUserRepository(userCollection *mongo.Collection) UserRepository {
	return &userRepository{opts: userRepositoryOpts{userCollection: userCollection}}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *models.RegistertUserInsertion) (string, error) {
	result, err := ur.opts.userCollection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (bool, error) {
	count, err := ur.opts.userCollection.CountDocuments(ctx, bson.M{"email": email})
	return count > 0, err
}
