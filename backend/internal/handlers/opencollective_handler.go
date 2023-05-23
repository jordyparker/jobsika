package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type WebhookPayload struct {
	CreatedAt    time.Time `json:"createdAt"`
	ID           int       `json:"id"`
	CollectiveID int       `json:"CollectiveId"`
	Type         string    `json:"type"`
	Data         struct {
		FromCollective struct {
			ID            int    `json:"id"`
			Type          string `json:"type"`
			Slug          string `json:"slug"`
			Name          string `json:"name"`
			TwitterHandle string `json:"twitterHandle"`
			GithubHandle  string `json:"githubHandle"`
			RepositoryURL string `json:"repositoryUrl"`
			Image         string `json:"image"`
		} `json:"fromCollective"`
		Collective struct {
			ID            int         `json:"id"`
			Type          string      `json:"type"`
			Slug          string      `json:"slug"`
			Name          string      `json:"name"`
			TwitterHandle interface{} `json:"twitterHandle"`
			GithubHandle  string      `json:"githubHandle"`
			RepositoryURL string      `json:"repositoryUrl"`
			Image         string      `json:"image"`
		} `json:"collective"`
		Transaction struct {
			ID                                int         `json:"id"`
			Kind                              string      `json:"kind"`
			Type                              string      `json:"type"`
			UUID                              string      `json:"uuid"`
			Group                             string      `json:"group"`
			Amount                            int         `json:"amount"`
			IsDebt                            bool        `json:"isDebt"`
			OrderID                           int         `json:"OrderId"`
			Currency                          string      `json:"currency"`
			IsRefund                          bool        `json:"isRefund"`
			ExpenseID                         interface{} `json:"ExpenseId"`
			CreatedAt                         time.Time   `json:"createdAt"`
			TaxAmount                         interface{} `json:"taxAmount"`
			Description                       string      `json:"description"`
			CollectiveID                      int         `json:"CollectiveId"`
			HostCurrency                      string      `json:"hostCurrency"`
			CreatedByUserID                   int         `json:"CreatedByUserId"`
			FromCollectiveID                  int         `json:"FromCollectiveId"`
			AmountInHostCurrency              int         `json:"amountInHostCurrency"`
			HostFeeInHostCurrency             int         `json:"hostFeeInHostCurrency"`
			NetAmountInHostCurrency           int         `json:"netAmountInHostCurrency"`
			PlatformFeeInHostCurrency         int         `json:"platformFeeInHostCurrency"`
			UsingGiftCardFromCollectiveID     interface{} `json:"UsingGiftCardFromCollectiveId"`
			NetAmountInCollectiveCurrency     int         `json:"netAmountInCollectiveCurrency"`
			AmountSentToHostInHostCurrency    int         `json:"amountSentToHostInHostCurrency"`
			PaymentProcessorFeeInHostCurrency int         `json:"paymentProcessorFeeInHostCurrency"`
			FormattedAmount                   string      `json:"formattedAmount"`
			FormattedAmountWithInterval       string      `json:"formattedAmountWithInterval"`
		} `json:"transaction"`
	} `json:"data"`
}

func OpenCollectiveWebhook(c *gin.Context) {
	var payload WebhookPayload
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Info("OpenCollectiveWebhook was triggered!")
}
