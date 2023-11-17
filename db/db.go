package db

import (
	"api/env"
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RDB
var ctx = context.Background()
var RDB *redis.Client
var Client *mongo.Client

var Accounts *mongo.Collection
var Balances *mongo.Collection
var TopupIntents *mongo.Collection
var Tickets *mongo.Collection
var TicketTypes *mongo.Collection
var Trips *mongo.Collection

func InitDB() {
  var err error

  Client, err = mongo.Connect(
    ctx,
    options.Client().ApplyURI(env.MongoURI),
  )
  
  if err != nil {
    log.Fatal(err)
  }

  // loading collections
  Accounts = GetCollection("accounts", Client)
  Balances = GetCollection("balances", Client)
  TopupIntents = GetCollection("topupIntents", Client)
  Tickets = GetCollection("tickets", Client)
  TicketTypes = GetCollection("ticketTypes", Client)
  Trips = GetCollection("trips", Client)

  fmt.Println("connected to mongodb")
}

func GetCollection(collectionName string, client *mongo.Client) (*mongo.Collection) {
  return client.Database("dev").Collection(collectionName)
}

func InitCache(RedisOptions *redis.Options) {
  RDB = redis.NewClient(RedisOptions)
  
  pong, _ := RDB.Ping(context.Background()).Result()
  if pong == "PONG" {
    fmt.Println("connected to redis")
  } else {
    fmt.Println("not connected to redis")
  }
}

func Set(key string, value string) error {
  err := RDB.Set(ctx, key, value, 0).Err()

  return err
}

func Get(key string) (string, error) {
  val, err := RDB.Get(ctx, key).Result()

  return val, err
}

func Del(key string) error {
  _, err := RDB.Del(ctx, key).Result()

  return err
}