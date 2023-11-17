package models

import (
	"api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type TicketType struct {
  ID string `bson:"id" json:"id"`
  Name string `bson:"name" json:"name"`
  Fare float64 `bson:"fare" json:"fare"`
  City string `bson:"city" json:"city"`
  Expiry string `bson:"expiry" json:"expiry"`
}

func GetTicketTypes(city string) ([]TicketType, error) {
    cursor, err := db.TicketTypes.Find(ctx, bson.M{
      "city": city,
    })

    if err != nil {
      return []TicketType {}, err
    }

    ticketTypes := []TicketType {}

    err = cursor.All(ctx, &ticketTypes)
    if len(ticketTypes) == 0 {
      ticketTypes = []TicketType {}
    }

    return ticketTypes, err
}

func GetTicketType(id string) (TicketType, error) {
  var ticketType TicketType

  err := db.TicketTypes.FindOne(
    ctx, 
    bson.M {
      "id": id,
    },
  ).Decode(&ticketType)

  return ticketType, err
}