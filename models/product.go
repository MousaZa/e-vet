package models

import (
	"github.com/MousaZa/e-vet/interfaces"
)

type Product struct {
	ID           string                `json:"id"`
	UserId       int                   `json:"user_id"`
	Name         string                `json:"name"`
	Quantity     int                   `json:"quantity"`
	LastPrice    float64               `json:"last_price"`
	IsActive     bool                  `json:"is_active"`
	ObserverList []interfaces.Observer `json:"observer_list"`
}

func (p *Product) Consume(amt int) {
	p.Quantity -= amt
	if p.Quantity < 5 {
		p.NotifyAll()
	}
}

func (p *Product) Register(o interfaces.Observer) {
	p.ObserverList = append(p.ObserverList, o)
}

func (p *Product) DeRegister(o interfaces.Observer) {
	p.ObserverList = removeFromslice(p.ObserverList, o)
}

func (p *Product) NotifyAll() {
	for _, observer := range p.ObserverList {
		observer.Update(p.Name)
	}
}

func removeFromslice(observerList []interfaces.Observer, observerToRemove interfaces.Observer) []interfaces.Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
