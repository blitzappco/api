package models

import (
	"api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type Balance struct {
  ID string `bson:"id" json:"id"`
  Value float64 `bson:"value" json:"value"`
}

func (balance *Balance) Create() error {
  // creating account
  _, err := db.Balances.InsertOne(ctx, balance)

  return err
}

func GetBalance(id string) (Balance, error) {
  var balance Balance

  err := db.Balances.FindOne(
    ctx,
    bson.M {
      "id": id,
    },
  ).Decode(&balance)

  return balance, err
}

func UpdateBalanceValue(id string, oldValue float64, newValue float64) error {
  _, err := db.Balances.UpdateOne(
    ctx,
    bson.M{"id": id},
    bson.M{
      "$set": bson.M {
        "value": oldValue + newValue,
      },
    },
  )

  return err
}

func ChargeBalance(id string, value float64) (Balance, error) {
  balance, err := GetBalance(id)
  if err != nil {
    return Balance {}, err
  }

  if balance.Value >= value {
    err = UpdateBalanceValue(id, balance.Value, -1 * value)
  }

  balance.Value -= value

  return balance, err
}

func TopupBalance(id string, value float64) (Balance, error) {
  balance, err := GetBalance(id)
  if err != nil {
    return Balance {}, err
  }

  err = UpdateBalanceValue(id, balance.Value, value)

  balance.Value += value

  return balance, err
}