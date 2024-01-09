package main

import (
	"fmt"
)

func main() {
	fmt.Println()  // So the start of the output is easier to distinguish from the file paths and such in the terminal

	dollarsInMachine := 550
	waterSupply := 400
	milkSupply := 540
	coffeeBeanSupply := 120
	cupsSupply := 9
	printSupplies(dollarsInMachine, waterSupply, milkSupply, coffeeBeanSupply, cupsSupply)

	switch action := getUserAction(); action {
	case "buy":
		order := getUserOrder()
		fmt.Println()
		printRemainingSupplies(order, dollarsInMachine, waterSupply, milkSupply, coffeeBeanSupply, cupsSupply)
	case "fill":
		waterSupply += getFillAmount("ml of water")
		milkSupply += getFillAmount("ml of milk")
		coffeeBeanSupply += getFillAmount("grams of coffee beans")
		cupsSupply += getFillAmount("disposable cups")
		fmt.Println()
		printSupplies(dollarsInMachine, waterSupply, milkSupply, coffeeBeanSupply, cupsSupply)
	case "take":
		fmt.Printf("I gave you $%d\n", dollarsInMachine)
		fmt.Println()
		dollarsInMachine = 0
		printSupplies(dollarsInMachine, waterSupply, milkSupply, coffeeBeanSupply, cupsSupply)
	}
}


func printSupplies(dollarsInMachine int, waterSupply int, milkSupply int, coffeeBeanSupply int, cupsSupply int) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", waterSupply)
	fmt.Printf("%d ml of milk\n", milkSupply)
	fmt.Printf("%d g of coffee beans\n", coffeeBeanSupply)
	fmt.Printf("%d disposable cups\n", cupsSupply)
	fmt.Printf("$%d of money\n\n", dollarsInMachine)

	return
}


func printRemainingSupplies(order string, dollarsInMachine, waterSupply, milkSupply, coffeeBeanSupply, cupsSupply int) {
	switch order {
	case "1":
		waterSupply -= 250
		coffeeBeanSupply -= 16
		dollarsInMachine += 4
	case "2":
		waterSupply -= 350
		milkSupply -= 75
		coffeeBeanSupply -= 20
		dollarsInMachine += 7
	case "3":
		waterSupply -= 200
		milkSupply -= 100
		coffeeBeanSupply -= 12
		dollarsInMachine += 6
	}
	// Every order requires one disposable cup
	cupsSupply -= 1

	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", waterSupply)
	fmt.Printf("%d ml of milk\n", milkSupply)
	fmt.Printf("%d g of coffee beans\n", coffeeBeanSupply)
	fmt.Printf("%d disposable cups\n", cupsSupply)
	fmt.Printf("$%d of money\n\n", dollarsInMachine)

	return
}


func getUserAction() (action string) {
	fmt.Println("Write action (buy, fill, take):")
	fmt.Scan(&action)

	return action
}


func getUserOrder() (order string) {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino")
	fmt.Scan(&order)

	return order
}


func getFillAmount(item string) (amount int) {
	fmt.Printf("Write how many %s you want to add:\n", item)
	fmt.Scanf("%d", &amount)

	return amount
}
