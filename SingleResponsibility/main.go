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


// =========================================
// BAD EXAMPLE - Violates Single Responsibility Principle (SRP)
// =========================================
//
// Definition Reminder:
// A struct should have only ONE reason to change.
//
// Problem in this example:
// OrderService is handling MULTIPLE responsibilities:
//
// 1. Saving order to database
// 2. Processing payment
// 3. Sending confirmation email
// 4. Generating invoice
//
// Why this violates SRP:
//
// If database logic changes → this struct changes
// If payment gateway changes → this struct changes
// If email provider changes → this struct changes
// If invoice format changes → this struct changes
//
// That means this struct has MULTIPLE reasons to change.
//
// This makes the code:
//
// ❌ Hard to maintain
// ❌ Hard to test
// ❌ Tightly coupled
// ❌ Difficult to scale
//
// Proper design would separate these responsibilities

// into different structs/services.

// type OrderService struct{}

// func (o OrderService) PlaceOrder(orderID int, amount float64) {

// 	// Responsibility 1: Database logic
// 	fmt.Printf("Saving order %d to database\n", orderID)

// 	// Responsibility 2: Payment processing
// 	fmt.Printf("Processing payment of %.2f\n", amount)

// 	// Responsibility 3: Email sending
// 	fmt.Println("Sending confirmation email")

// 	// Responsibility 4: Invoice generation
// 	fmt.Printf("Generating invoice for order %d\n", orderID)
// }

// func main() {
// 	service := OrderService{}
// 	service.PlaceOrder(1, 5000)
// }



// =============== PERFECT EXAMPLE ===============
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

func (os OrderService) PlaceOrder(orderId int, amount int) {
	os.repo.Save(orderId)
	os.payment.Process(float64(amount))
	os.email.Send()
	os.invoice.Generate(orderId)
}
