package models

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

type Product struct {
	ID           string   `json:"id"`
	UserId       int      `json:"user_id"`
	Name         string   `json:"name"`
	Quantity     int      `json:"quantity"`
	LastPrice    float64  `json:"last_price"`
	IsActive     bool     `json:"is_active"`
	ObserverList []string `json:"observer_list"`
}

func (p *Product) Consume(amt int, db *firestore.Client, ctx context.Context) error {
	p.Quantity -= amt
	if p.Quantity < 5 {
		return p.NotifyAll(db, ctx)
	}
	return nil
}

func (p *Product) Register(u string) {
	p.ObserverList = append(p.ObserverList, u)
}

func (p *Product) NotifyAll(db *firestore.Client, ctx context.Context) error {

	n := NewNotification(LowStock, fmt.Sprintf("Running Low on %s", p.Name), fmt.Sprintf("Running Low on %s, last time the price was %.2f", p.Name, p.LastPrice))

	for _, observerId := range p.ObserverList {
		doc := db.Collection("Users").Doc(observerId)
		_, _, err := doc.Collection("Notifications").Add(ctx, n)
		return err
	}
	return nil
}
