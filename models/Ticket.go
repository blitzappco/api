package models

import (
	"api/db"
	"api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Ticket struct {
  ID string `bson:"id" json:"id"`
  AccountID string `bson:"accountID" json:"accountID"`
  Name string `bson:"name" json:"name"`
  City string `bson:"city" json:"city"`
  Price float64 `bson:"price" json:"price"`
  ExpiresAt time.Time `bson:"expiresAt" json:"expiresAt"`
  CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

func GetTickets(accountID string) ([]Ticket, error) {
  cursor, err := db.Tickets.Find(ctx, bson.M{
    "accountID": accountID,
  })

  if err != nil {
    return []Ticket {}, err
  }

  tickets := []Ticket {}

  err = cursor.All(ctx, &tickets)
  if len(tickets) == 0 {
    tickets = []Ticket {}
  }

  return tickets, err
}

func GetLastTicket(accountID string) (Ticket, error) {
  tickets := []Ticket{}
  
  cursor, err := db.Tickets.Find(ctx, 
    bson.M {"accountID": accountID}, 
    options.Find().SetLimit(1).
      SetSort(bson.M {"expiresAt": -1}))
  
  if err != nil {
    return tickets[0], err
  }

  err = cursor.All(ctx, &tickets)
  if err != nil {
    return tickets[0], err
  }  

  return tickets[0], nil
}


func (ticket *Ticket) Create(expiry string) error {
  ticket.ID = utils.GenID(9)

  loc, _ := time.LoadLocation("Europe/Bucharest")
  ticket.CreatedAt = time.Now().In(loc)
  switch expiry {
  case "1h":
    ticket.ExpiresAt = ticket.CreatedAt.Add(time.Hour)
  case "1d":
    ticket.ExpiresAt = ticket.CreatedAt.AddDate(0, 0, 1)
  }

  _, err := db.Tickets.InsertOne(ctx, ticket)
  return err
}