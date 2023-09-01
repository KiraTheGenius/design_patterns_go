package main

import (
	"fmt"
	"log"
)

type Payment interface {
	Pay()
}

type Saman struct {
}

func (s *Saman) Pay() {
	fmt.Println("Paying by Saman IPG")
}

type AP struct {
}

func (a *AP) Pay() {
	fmt.Println("Paying by AP IPG")
}

type PaymentType uint8

const (
	PaymentSaman PaymentType = iota
	PaymentAP
)

func PaymentFactory(t PaymentType) Payment {
	switch t {
	case PaymentSaman:
		return new(Saman)
	case PaymentAP:
		return new(AP)
	default:
		log.Fatal("undefined payment type")
		return nil
	}
}

func main() {
	payment := PaymentFactory(PaymentAP)
	payment.Pay()
}
