// Create order app for coffee
package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID int
	Customer string
	Coffee string
	Size string
	Status string
}

func processOrder(order Order, statusChannel chan Order){
	time.Sleep(2 * time.Second)
	order.Status = "Completed"
	statusChannel <- order

}

func main(){
	fmt.Println("This is main function")
	statusChannel := make(chan Order)
	orders := []Order{
		{ID:1, Customer:"Alice", Coffee: "Espresso", Size:"small"},
		{ID:2, Customer:"Bob", Coffee: "Cappuccino", Size:"small"},
		{ID:3, Customer:"Arzoo", Coffee: "Moccha", Size:"small"},
		{ID:4, Customer:"Ujjwol", Coffee: "Espresso", Size:"small"},
	}
	for _, order := range orders{
		go processOrder(order, statusChannel)
	}
	for i:=0; i<len(orders); i++{
		o := <-statusChannel
		fmt.Printf("Order %d: %s (%s) for %s is %s\n",o.ID, o.Coffee, o.Size, o.Customer, o.Status)
	}
}