package main

import (
	"fmt"
	"os"
)

type Bill struct {
	Name  string
	Items map[string]float64
	Tip   float64
}

// Make new bills
func newBill(name string) Bill {
	b := Bill{
		Name:  name,
		Items: map[string]float64{},
		Tip:   0,
	}
	return b
}

// Add item to bill
func (b *Bill) addItem(name string, price float64) {
	b.Items[name] = price
}

// Format the bill
func (b *Bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	// List items
	for k, v := range b.Items {
		fs += fmt.Sprintf("%-25v £%v\n", k+":", v)
		total += v
	}

	// Add tip
	fs += fmt.Sprintf("%-25v £%v\n", "Tip:", b.Tip)

	// Add total
	fs += fmt.Sprintf("%-25v £%0.2f", "Total:", total+b.Tip)

	return fs
}

// Update tip
func (b *Bill) updateTip(tip float64) {
	(*b).Tip = tip
}

// Save bill
func (b *Bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.Name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to file")
}
