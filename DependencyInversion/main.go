// =====================================================
// DEPENDENCY INVERSION PRINCIPLE (DIP)
// =====================================================
//
// Definition:
//
// 1) High-level modules should not depend on low-level modules.
//    Both should depend on abstractions.
//
// 2) Abstractions should not depend on details.
//    Details should depend on abstractions.
//
// -----------------------------------------------------
// What is High-Level?
// → Business logic (core application rules)
//
// What is Low-Level?
// → Database, external APIs, file systems, frameworks
//
// -----------------------------------------------------
// Why DIP?
//
// - Keeps business logic independent
// - Makes infrastructure replaceable
// - Enables easier testing
// - Supports OCP (Open/Closed Principle)
//
// -----------------------------------------------------
// Golden Rule in Go:
//
// Define interfaces where they are CONSUMED,
// not where they are IMPLEMENTED.
//
// -----------------------------------------------------
// Memory Trick:
//
// Business logic depends on abstraction.
// Implementation depends on abstraction.
//

package main

import "fmt"

// Low-level module (PDF implementation)
type PDFGenerator struct{}

func (p PDFGenerator) Generate(content string) {
	fmt.Println("Generating PDF with content:", content)
}

// High-level module (Business logic)
type ReportService struct {
	pdf PDFGenerator // ❌ depends on concrete implementation
}

func (r ReportService) CreateReport() {
	content := "Annual Financial Report"
	r.pdf.Generate(content)
}


// =============================================
// GOOD EXAMPLE
// =============================================

// Abstraction (defined by high-level module)
type ReportGenerator interface {
	Generate(content string)
}

// High-level module
type ReportServiceOne struct {
	generator ReportGenerator // ✅ depends on abstraction
}

func NewReportServiceOne(generator ReportGenerator) *ReportServiceOne {
	return &ReportServiceOne{generator: generator}
}

func (r ReportServiceOne) CreateReport() {
	content := "Annual Financial Report"
	r.generator.Generate(content)
}

func main() {
	pdf := PDFGenerator{}
	service := NewReportServiceOne(pdf)

	service.CreateReport()
}
