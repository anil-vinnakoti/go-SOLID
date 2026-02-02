// =========================================
// OPEN / CLOSED PRINCIPLE (OCP)
// =========================================
//
// Definition:
// Software entities should be OPEN for extension
// but CLOSED for modification.
//
// Meaning:
// We should be able to add new behavior
// without changing existing working code.
//
// In this example:
//
// - Notification is an interface.
// - EmailService and SmsService implement Notification.
// - SendNotification depends on the interface, NOT concrete types.
//
// Why this follows OCP:
//
// If we want to add a new notification type (e.g., PushNotification),
// we DO NOT modify:
//
//   - Notification interface
//   - SendNotification function
//
// We only create a new struct that implements Send().
//
// This makes the system:
//
// ✅ Easy to extend
// ✅ Safe from breaking existing code
// ✅ Scalable for future features
//
// Key Design Idea:
// Depend on abstractions (interfaces), not concrete implementations.
//
// Bad Design (Violates OCP):
// Using if-else or switch on type strings.
// Every new feature would require modifying existing code.
//
// Good Design (Follows OCP):
// Use interfaces so new behavior is added via new implementations.
//
// Mental Check:
// If adding a new feature forces me to edit old tested code,
// I am probably violating OCP.

// =============
// BAD EXAMPLE - Violates Open / Closed Principle (OCP)
// ===============
// // PaymentProcessor handles different payment types
// type PaymentProcessor struct{}

// // This function violates OCP
// func (p PaymentProcessor) ProcessPayment(method string, amount float64) {
// 	if method == "credit" {
// 		fmt.Println("Processing credit card payment of", amount)
// 	} else if method == "paypal" {
// 		fmt.Println("Processing PayPal payment of", amount)
// 	} else if method == "upi" {
// 		fmt.Println("Processing UPI payment of", amount)
// 	} else {
// 		fmt.Println("Unsupported payment method")
// 	}
// }

// func main() {
// 	processor := PaymentProcessor{}

// 	processor.ProcessPayment("credit", 1000)
// 	processor.ProcessPayment("paypal", 2000)
// }

// ======== PERFECT EXAMPLE ==========

package main

import "fmt"

type Notification interface {
	Send()
}

type EmailService struct{}

func (e EmailService) Send() {
	fmt.Println("Sending email...")
}

type SmsService struct{}

func (s SmsService) Send() {
	fmt.Println("Sending SMS...")
}

func SendNotification(n Notification) {
	n.Send()
}

func main() {
	email := EmailService{}
	sms := SmsService{}

	SendNotification(email)
	SendNotification(sms)
}

// Future case

// type SlackService struct{}
// func (s SlackService) Send(){
// 	fmt.Println("sending slack notification...")
// }
