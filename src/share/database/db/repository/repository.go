package repository

import (
	"context"
	"errors"
	"github.com/reddit-clone/src/share/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AbstractRepository[T any] interface {
	Create(ctx context.Context, entity any) error
	Update(ctx context.Context, id string, update any) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]T, error)
	GetById(ctx context.Context, id string) (T, error)
	Drop(ctx context.Context) error
}
type MongoDbAbstractRepository[T any] struct {
	Collection *mongo.Collection
}

func (m *MongoDbAbstractRepository[T]) Create(ctx context.Context, entity any) error {
	_, err := m.Collection.InsertOne(ctx, entity)
	return err
}

func (m *MongoDbAbstractRepository[T]) Update(ctx context.Context, id string, update any) error {
	obi, err := pkg.StringToObjectId(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": obi}
	_, updateErr := m.Collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	return updateErr
}

func (m *MongoDbAbstractRepository[T]) Delete(ctx context.Context, id string) error {
	obi, err := pkg.StringToObjectId(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": obi}
	_, deleteErr := m.Collection.DeleteOne(ctx, filter)
	return deleteErr
}

func (m *MongoDbAbstractRepository[T]) GetAll(ctx context.Context) ([]T, error) {
	cursor, err := m.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var entities []T
	err = cursor.All(ctx, &entities)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return []T{}, nil
		}
		return nil, err
	}
	return entities, nil
}
func (m *MongoDbAbstractRepository[T]) Drop(ctx context.Context) error {
	err := m.Collection.Drop(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (m *MongoDbAbstractRepository[T]) GetById(ctx context.Context, id string) (T, error) {
	obi, err := pkg.StringToObjectId(id)
	if err != nil {
		return zero((*T)(nil)), err
	}
	filter := bson.M{"_id": obi}
	var data T
	decodeErr := m.Collection.FindOne(ctx, filter, nil).Decode(&data)
	if decodeErr != nil {
		return zero((*T)(nil)), decodeErr
	}
	return data, nil
}

func zero[T any](t *T) T {
	return *new(T)
}
