package main

import "fmt"

func main() {
	hpPrinter := &HpPrinter{}
	canonPrinter := &CanonPrinter{}

	pc := NewMacPc(hpPrinter)
	pc.Print("hello")
	pc.SetPrinter(canonPrinter)
	pc.Print("world!")
}

type Computer interface {
	Print(text string)
	SetPrinter(printer Printer)
}

type Printer interface {
	PrintSomething(text string)
}

type MacPc struct {
	printer Printer
}

func NewMacPc(printer Printer) Computer {
	return &MacPc{printer: printer}
}

func (m *MacPc) SetPrinter(printer Printer) {
	m.printer = printer
}

func (m *MacPc) Print(text string) {
	fmt.Println("Mac PC printing")
	m.printer.PrintSomething(text)
}

type HpPrinter struct{}

func (h *HpPrinter) PrintSomething(text string) {
	fmt.Println("HP printer prints:", text)
}

type CanonPrinter struct{}

func (c *CanonPrinter) PrintSomething(text string) {
	fmt.Println("Canon printer prints:", text)
}
