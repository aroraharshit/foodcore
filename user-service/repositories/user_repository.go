package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/aroraharshit/go-foodcore/user-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.RegistertUserInsertion) (string, error)
	FindByEmail(ctx context.Context, email string) (bool, error)
	FindByMobile(ctx context.Context, mobile string) (bool, primitive.ObjectID, error)
	FetchHashPassword(ctx context.Context, userId primitive.ObjectID) (string, error)
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

func (ur *userRepository) FindByMobile(ctx context.Context, mobile string) (bool, primitive.ObjectID, error) {
	result := ur.opts.userCollection.FindOne(ctx, bson.M{"mobile": mobile})
	err := result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, primitive.NilObjectID, errors.New("no user exists from this mobile number")
		}
		return false, primitive.NilObjectID, err
	}

	type response struct {
		UserId primitive.ObjectID `json:"userId" bson:"_id"`
	}

	var res response

	if err := result.Decode(&res); err != nil {
		fmt.Println(err)
		return false, primitive.NilObjectID, err
	}

	return true, res.UserId, nil
}

func (ur *userRepository) FetchHashPassword(ctx context.Context, userId primitive.ObjectID) (string, error) {
	type HashedPassword struct {
		StoredPassword string `bson:"password"`
	}

	var password HashedPassword

	result := ur.opts.userCollection.FindOne(ctx, bson.M{"_id": userId}, options.FindOne().SetProjection(bson.M{"password": 1}))
	err := result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("no user exists from this mobile number")
		}
		return "", err
	}

	err = result.Decode(&password)
	if err != nil {
		fmt.Printf("Error decoding password")
		return "", err
	}

	return password.StoredPassword, nil
}
