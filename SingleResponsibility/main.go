package main

import "fmt"

type OrderRepository struct{}

func (o OrderRepository) Save(orderID int) {
	fmt.Printf("Saving order %d to database\n", orderID)
}

type PaymentService struct{}

func (p PaymentService) Process(amount float64) {
	fmt.Printf("Processing payment of %.2f\n", amount)
}

type EmailService struct{}

func (e EmailService) Send() {
	fmt.Println("Sending confirmation email")
}

type InvoiceService struct{}

func (i InvoiceService) Generate(orderID int) {
	fmt.Printf("Generating invoice for order %d\n", orderID)
}

type OrderService struct {
	repo    OrderRepository
	payment PaymentService
	email   EmailService
	invoice InvoiceService
}

func (os OrderService) PlaceOrder(orderId int, amount int){
	os.repo.Save(orderId)
	os.payment.Process(float64(amount))
	os.email.Send()
	os.invoice.Generate(orderId)
}
// ===============================
// SINGLE RESPONSIBILITY PRINCIPLE (SRP)
// ===============================
//
// Definition:
// A struct should have only ONE reason to change.
//
// In this example:
//
// OrderRepository  → Responsible only for saving orders to DB.
// PaymentService   → Responsible only for processing payments.
// EmailService     → Responsible only for sending emails.
// InvoiceService   → Responsible only for generating invoices.
// OrderService     → Responsible only for coordinating the order workflow.
//
// Why this follows SRP:
//
// - If database logic changes → Only OrderRepository changes.
// - If payment gateway changes → Only PaymentService changes.
// - If email provider changes → Only EmailService changes.
// - If invoice format changes → Only InvoiceService changes.
// - If order flow changes → Only OrderService changes.
//
// Each struct has exactly ONE responsibility.
// Each struct has exactly ONE reason to change.
//
// Mental Check:
// If I can give more than one reason for a struct to change,
// then SRP is violated.
//
// Benefits:
// - Easier maintenance
// - Better testability
// - Cleaner separation of concerns
// - More scalable backend design
//
// Backend Architecture Mapping:
// Handler → Service → Repository
//
// SRP ensures each layer has a clear and focused responsibility.
