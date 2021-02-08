package main

import (
	"time"

	paymentModels "github.com/cloudentity/openbanking-sandbox/openbanking/paymentinitiation/models"
)

type PaymentQueue struct {
	repo  UserRepo
	queue chan paymentModels.OBWriteDomesticResponse5
}

func NewPaymentQueue(repo UserRepo) (PaymentQueue, error) {
	return PaymentQueue{
		repo:  repo,
		queue: make(chan paymentModels.OBWriteDomesticResponse5),
	}, nil
}

func (pq *PaymentQueue) Start() {
	for payment := range pq.queue {
		time.Sleep(10 * time.Second)
		pq.repo.SetDomesticPaymentStatus(*payment.Data.DomesticPaymentID, AcceptedSettlementCompleted) //nolint
	}
}
