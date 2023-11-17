package models

import (
	"api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trip struct {
  AccountID string `json:"accountID" bson:"accountID"`
  Name string `json:"name" bson:"name"`
  Address string `json:"address" bson:"address"`
  Type string `json:"type" bson:"type"`
  Longitude float64 `json:"longitude" bson:"longitude"`
  Latitude float64 `json:"latitude" bson:"latitude"`

  DateTime time.Time `json:"datetime" bson:"datetime"`
}

func (t *Trip) Insert() error {
  t.DateTime = time.Now()

  _, err := db.Trips.InsertOne(ctx, t)

  return err
}

func GetTrips(accountID string) ([]Trip, error) {
  trips := []Trip {}

  cursor, err := db.Trips.Find(ctx, bson.M {
    "accountID": accountID,
  }, options.Find().SetSort(bson.M {
    "datetime": -1,
  }))
  if err != nil {
    return trips, err
  }

  err = cursor.All(ctx, &trips)

  return trips, err
} 