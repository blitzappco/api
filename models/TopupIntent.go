package models

import (
	"api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type TopupIntent struct {
  AccountID string `bson:"accountID" json:"accountID"`
  StripeIntentID string `bson:"stripeIntentID" json:"stripeIntentID"`
  Amount float64 `bson:"amount" json:"amount"`
  Confirmed bool `bson:"confirmed" json:"confirmed"`
}

func (intent *TopupIntent) Create() error {
  _, err := db.TopupIntents.InsertOne(ctx, intent)

  return err
}

func (intent *TopupIntent) Find(stripeIntentID string) error {
  return db.TopupIntents.FindOne(ctx, bson.M{
    "stripeIntentID": stripeIntentID,
  }).Decode(&intent)
}

func (intent *TopupIntent) Confirm() error {
  if !intent.Confirmed {
    _, err := TopupBalance(intent.AccountID, intent.Amount)
  
    if err != nil {
      return err
    }
    _, err = db.TopupIntents.UpdateOne(ctx, bson.M{
      "stripeIntentID": intent.StripeIntentID,
    }, bson.M {
      "$set": bson.M{
        "confirmed": true,
      },
    })

    intent.Confirmed = true
    return err
  } else {
    return nil
  }

}