package accounts

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func balance(acc fiber.Router) {
  balance := acc.Group("/balance")

  balance.Get("", models.AccountMiddleware, func (c *fiber.Ctx) error {
    id := fmt.Sprintf("%v", c.Locals("id"))

    balance, err := models.GetBalance(id)
    if err != nil {
      return utils.MessageError(c, "Nu s-a putut gasi")
    }

    return c.JSON(balance)
  })

  balance.Post("/test-charge", models.AccountMiddleware, func (c *fiber.Ctx) error {
    var body map[string]float64
    json.Unmarshal(c.Body(), &body)
    value := body["value"]
    
    id := fmt.Sprintf("%v", c.Locals("id"))

    balance, err := models.ChargeBalance(id, value)
    if err != nil {
      return utils.MessageError(c, "eroare")
    }

    return c.JSON(balance)
  })

  balance.Post("/test-topup", models.AccountMiddleware, func (c *fiber.Ctx) error {
    var body map[string]float64
    json.Unmarshal(c.Body(), &body)
    value := body["value"]
    
    id := fmt.Sprintf("%v", c.Locals("id"))

    balance, err := models.TopupBalance(id, value)
    if err != nil {
      return utils.MessageError(c, "eroare")
    }

    return c.JSON(balance)
  })
}