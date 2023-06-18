package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Add item to bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// Format the bill
func (b *bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v £%v\n", k+":", v)
		total += v
	}

	// Add tip
	fs += fmt.Sprintf("%-25v £%v\n", "Tip:", b.tip)

	// Add total
	fs += fmt.Sprintf("%-25v £%0.2f", "Total:", total+b.tip)

	return fs
}

// Update tip
func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
}

// Save bill
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to file")
}
