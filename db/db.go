package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Car struct {
	Reg   string             `bson:"reg"`
	Make  string             `bson:"name"`
	Model string             `bson:"model"`
	ID    primitive.ObjectID `bson:"_id"`
	Hired bool               `bson:"hired"`
	Rate  int                `bson:"rate"`
}

type Customer struct {
	Name    string             `bson:"name"`
	ID      primitive.ObjectID `bson:"_id"`
	Balance float32            `bson:"balance"`
}

type Reservation struct {
	ID       primitive.ObjectID `bson:"_id"`
	Car      primitive.ObjectID `bson:"car"`
	Customer primitive.ObjectID `bson:"customer"`
}

func ConnectToDB(ctx context.Context) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	database := client.Database("carpark")

	return database, nil
}

func (c *Car) insert(ctx context.Context) error {
	if dbInter := ctx.Value("db"); dbInter != nil {
		db, ok := dbInter.(mongo.Database)
		if !ok {
			return errors.New("db key in context does not yeild a mongo db")
		}
		carCollection := db.Collection("cars")
		_, err := carCollection.InsertOne(ctx, c)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("No db key in context")
}

func (c *Customer) insert(ctx context.Context) error {
	if dbInter := ctx.Value("db"); dbInter != nil {
		db, ok := dbInter.(mongo.Database)
		if !ok {
			return errors.New("db key in context does not yeild a mongo db")
		}
		customerCollection := db.Collection("customer")
		_, err := customerCollection.InsertOne(ctx, c)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("No db key in context")
}

func (r *Reservation) insert(ctx context.Context) error {
	if dbInter := ctx.Value("db"); dbInter != nil {
		db, ok := dbInter.(mongo.Database)
		if !ok {
			return errors.New("db key in context does not yeild a mongo db")
		}
		reservationCollection := db.Collection("reservations")
		_, err := reservationCollection.InsertOne(ctx, r)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("No db key in context")
}
