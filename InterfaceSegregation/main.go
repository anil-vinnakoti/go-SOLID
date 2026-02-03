// =========================================
// INTERFACE SEGREGATION PRINCIPLE (ISP)
// =========================================
//
// Definition:
// Clients should not be forced to depend on methods
// they do not use.
//
// Meaning:
// Avoid large, "fat" interfaces.
// Prefer small, focused interfaces.
//
// Why?
// Large interfaces create:
// ❌ Tight coupling
// ❌ Harder testing (big mocks)
// ❌ Unnecessary dependencies
// ❌ Confusing design
//
// In Go, small interfaces are preferred.
// Example: io.Reader has only one method.
//
// Mental Check:
// If some implementations leave methods empty
// or panic because they are not needed,
// ISP is likely being violated.

// =========================================
// BAD EXAMPLE - Violates ISP
// =========================================
//
// Problem:
// The interface is too large and forces
// implementations to define methods
// they do not actually support.
//

package main

import "fmt"

type Machine interface {
	Print()
	Scan()
	Fax()
}

type SimplePrinterOne struct{}

func (s SimplePrinterOne) Print() {
	fmt.Println("Printing document")
}

// SimplePrinterOne does not support Scan or Fax,
// but it is forced to implement them.

func (s SimplePrinterOne) Scan() {
	panic("Scan not supported")
}

func (s SimplePrinterOne) Fax() {
	panic("Fax not supported")
}

// Why this violates ISP:
//
// - SimplePrinterOne depends on methods it does not use.
// - Unnecessary methods lead to panic or empty logic.
// - Interface is too broad.

// =========================================
// GOOD EXAMPLE - Follows ISP
// =========================================
//
// Solution:
// Split the large interface into smaller,
// focused interfaces.
//

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type Faxer interface {
	Fax()
}

// SimplePrinter only implements Printer
type SimplePrinter struct{}

func (s SimplePrinter) Print() {
	fmt.Println("Printing document")
}

// AdvancedMachine implements multiple small interfaces
type AdvancedMachine struct{}

func (a AdvancedMachine) Print() {
	fmt.Println("Printing document")
}

func (a AdvancedMachine) Scan() {
	fmt.Println("Scanning document")
}

func (a AdvancedMachine) Fax() {
	fmt.Println("Sending fax")
}

// Why this follows ISP:
//
// - SimplePrinter depends only on Printer.
// - No unused methods.
// - No panic implementations.
// - Flexible and scalable design.
