// ===============================
// LISKOV SUBSTITUTION PRINCIPLE (LSP)
// ===============================
//
// Definition:
// Subtypes must be substitutable for their base types
// without breaking expected behavior.
//
// Key Concept:
// - Implementing an interface is NOT enough
// - Implementation must honor the behavioral contract
//
// In this example:
//
// - Bird interface defines Fly() behavior
// - Sparrow implements Bird and can fly → ✅ safe
// - Ostrich cannot fly → ❌ would break Bird contract
//
// Why This Matters:
//
// - Any code that accepts Bird can safely call Fly()
// - No surprises, no runtime errors
//
// Real Backend Analogy:
//
// - Interface: PaymentMethod { Process(amount float64) error }
// - Any implementation must actually process payment
// - Returning nil without processing violates LSP
//
// Mental Check:
// If replacing a base type with a subtype breaks code,
// the Liskov Substitution Principle is violated.
//
// Key Takeaways:
//
// - Separate interfaces if behavior differs (e.g., RunningBird vs FlyingBird)
// - Always honor the "promise" of the interface
// - LSP ensures reliable polymorphism in Go


// =========================================
// GOOD EXAMPLE - Liskov Substitution Principle (LSP)
// =========================================
//
// Definition Reminder:
// If S is a subtype of T, objects of type T
// should be replaceable with objects of type S
// without breaking the program.
//
// Key Idea:
// Implementing an interface is not enough;
// the implementation must satisfy the behavioral contract.

package main

import "fmt"

// Bird interface defines the behavior of a bird
type Bird interface {
	Fly()
}

// Sparrow is a normal bird that can fly
type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying")
}

// Ostrich is a bird that cannot fly, so it should NOT implement Fly()
// Instead, we separate behaviors to avoid LSP violation
// For example, we could have a RunningBird interface for ostrich

// This is an example of adhering to LSP: 
// any struct implementing Bird can safely be used where Bird is expected
func MakeBirdFly(b Bird) {
	b.Fly() // We assume every Bird can fly
}

func main() {
	sparrow := Sparrow{}

	// Works fine because Sparrow behaves as expected
	MakeBirdFly(sparrow)

	// If we tried to pass an Ostrich here, it would violate LSP
	// because Ostrich cannot fly and would break MakeBirdFly
}
