package accounts

import (
	"api/models"
	"api/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func trips(r fiber.Router) {
  g := r.Group("/trips")

  g.Post("/", models.AccountMiddleware, func (c *fiber.Ctx) error {
    var trip models.Trip
    json.Unmarshal(c.Body(), &trip)

    var account models.Account
    utils.GetLocals(c, "account", &account)

    trip.AccountID = account.ID

    err := trip.Insert()
    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(trip)
  })

  g.Get("/", models.AccountMiddleware, func (c *fiber.Ctx) error {
    var account models.Account
    utils.GetLocals(c, "account", &account)

    trips, err := models.GetTrips(account.ID)
    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(trips)
  })
}