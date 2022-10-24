package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goproj/internal/result"
	"log"
	"os"
	"time"
)

type Connection interface {
	Close()
	DB() *mongo.Database
}

type conn struct {
	client *mongo.Client
}

func NewConnection() Connection {
	var c conn
	var err error
	auth := getAuth()
	c.client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://" + auth.Hostname + ":" + auth.Port).SetAuth(options.Credential{
		Username: auth.UserId,
		Password: auth.Password,
	}))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = c.client.Connect(ctx)

	if err != nil {
		log.Panicln(err.Error())
	}
	return &c
}

func (c *conn) Close() {
	err := c.client.Disconnect(context.Background())
	if err != nil {
		return
	}
}

func (c *conn) DB() *mongo.Database {
	const dbName = "smartPassword"
	return c.client.Database(dbName)
}

type Auth struct {
	UserId   string `json:"userid"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

func getAuth() Auth {
	data, err := os.ReadFile("configs/mongodb_auth.json")
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		log.Fatalf("Failed to read file: %v\n", err)
	}

	var u Auth
	json.NewDecoder(bytes.NewBuffer(data)).Decode(&u)

	return u
}

func FindOne[T any](collection string, filter bson.M) (T, error) {
	conn := NewConnection()
	defer conn.Close()

	col := conn.DB().Collection(collection)

	var item T

	data := col.FindOne(context.Background(), filter)
	if err := data.Decode(&item); err != nil {
		if err == mongo.ErrNoDocuments {
			return item, nil
		}
		return item, err
	}
	return item, nil
}

func FindAll[T any](collection string, filter bson.M, option *options.FindOptions) (*result.PageResult[[]T], error) {
	conn := NewConnection()
	defer conn.Close()

	col := conn.DB().Collection(collection)

	data, err := col.Find(context.Background(), filter, option)
	if err != nil {
		return nil, err
	}

	var results = make([]T, 0)
	if err := data.All(context.Background(), &results); err != nil {
		return nil, err
	}

	count, _ := col.CountDocuments(context.Background(), filter)
	return &result.PageResult[[]T]{
		Item:       results,
		TotalCount: int(count),
	}, nil
}

func AddDocument[T any](collection string, doc T) error {
	conn := NewConnection()
	defer conn.Close()

	col := conn.DB().Collection(collection)

	response, err := col.InsertOne(context.Background(), doc)
	if err != nil {
		return err
	}
	var insertedInfra bson.M
	query := bson.D{{"_id", response.InsertedID}}
	if err := col.FindOne(context.Background(), query).Decode(&insertedInfra); err != nil {
		return err
	}
	return nil
}

func DropCollection(collection string) error {
	conn := NewConnection()
	defer conn.Close()

	col := conn.DB().Collection(collection)

	if err := col.Drop(context.Background()); err != nil {
		return err
	}

	return nil
}
