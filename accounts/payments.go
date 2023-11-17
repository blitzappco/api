package accounts

import (
	"api/models"
	"api/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"go.mongodb.org/mongo-driver/bson"
)

func payments (acc fiber.Router) {
  payments := acc.Group("/payments")

  payments.Post("/topup-intent", models.AccountMiddleware, func (c *fiber.Ctx) error {
    var account models.Account
    utils.GetLocals(c, "account", &account)

    var body map[string]int64
    json.Unmarshal(c.Body(), &body)

    // amountToCharge := body["amount"] + 100 + 25 * body["amount"] / 1000
    amountToCharge := body["amount"]

    params := &stripe.PaymentIntentParams {
      Customer: stripe.String(account.StripeCustomerID),
      PaymentMethodTypes: []*string {
        stripe.String("card"),
      },
      Amount: stripe.Int64(amountToCharge),
      Currency: stripe.String(string(stripe.CurrencyRON)),
    }
    intent, _ := paymentintent.New(params);

    topupIntent := models.TopupIntent {
      Amount: float64(body["amount"]) / 100,
      AccountID: account.ID,
      StripeIntentID: intent.ID,
      Confirmed: false,
    }
    err := topupIntent.Create()
    if err != nil {
      return utils.MessageError(c, "nu s-a putut crea")
    }


    return c.JSON(bson.M {
      "clientSecret": intent.ClientSecret,
      "name": account.LastName + " " + account.FirstName,
      "id": intent.ID,
    })
  })

  payments.Post("/topup-confirm", models.AccountMiddleware, func (c *fiber.Ctx) error {
    var body map[string]string
    json.Unmarshal(c.Body(), &body)
    stripeIntentID := body["stripeIntentID"]

    var intent models.TopupIntent
    err := intent.Find(stripeIntentID)
    if err != nil {
      return utils.MessageError(c, "nu s-a putut gasi tranzactia")
    }
    err = intent.Confirm()
    if err != nil {
      return utils.MessageError(c ,"nu s-a confirma")
    }

    return c.JSON("ok")
  })
}