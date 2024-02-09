package main

import (
	"fmt"
	"os"
)

// TODO after checking resources after buy: I have enough resources, making you a coffee!
// TODO after checking resources after buy: Sorry, not enough water!

func main() {
	fmt.Println()  // So the start of the output is easier to distinguish from the file paths and such in the terminal

	dollarsInMachine := 550
	waterSupply := 400
	milkSupply := 540
	coffeeBeanSupply := 120
	cupsSupply := 9

	for {
		switch action := getUserAction(); action {
		case "buy":
			order := getUserOrder()
			if order == "back" {
				continue
			}

			if waterSupply < waterRequired(order) {
				fmt.Println("Sorry, not enough water!")
				fmt.Println()
				continue
			} else if milkSupply < milkRequired(order) {
				fmt.Println("Sorry, not enough milk!")
				fmt.Println()
				continue
			} else if coffeeBeanSupply < coffeeBeansRequired(order) {
				fmt.Println("Sorry, not enough coffee beans!")
				fmt.Println()
				continue
			} else if cupsSupply < 1 {
				fmt.Println("Sorry, not enough disposable cups!")
				fmt.Println()
				continue
			} else {
				fmt.Println("I have enough resources, making you a coffee!")
				fmt.Println()
			}

			dollarsInMachine += costOfOrder(order)
			waterSupply -= waterRequired(order)
			milkSupply -= milkRequired(order)
			coffeeBeanSupply -= coffeeBeansRequired(order)
			cupsSupply -= 1  // All orders require 1 disposable cup
		case "fill":
			waterSupply += getFillAmount("ml of water")
			milkSupply += getFillAmount("ml of milk")
			coffeeBeanSupply += getFillAmount("grams of coffee beans")
			cupsSupply += getFillAmount("disposable cups")
			fmt.Println()
		case "take":
			fmt.Printf("I gave you $%d\n", dollarsInMachine)
			fmt.Println()
			dollarsInMachine = 0
		case "remaining":
			printSupplies(dollarsInMachine, waterSupply, milkSupply, coffeeBeanSupply, cupsSupply)
		case "exit":
			os.Exit(0)
		}
	}
}


func printSupplies(dollarsInMachine int, waterSupply int, milkSupply int, coffeeBeanSupply int, cupsSupply int) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", waterSupply)
	fmt.Printf("%d ml of milk\n", milkSupply)
	fmt.Printf("%d g of coffee beans\n", coffeeBeanSupply)
	fmt.Printf("%d disposable cups\n", cupsSupply)
	fmt.Printf("$%d of money\n", dollarsInMachine)

	fmt.Println()
	return
}


func getUserAction() (action string) {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	fmt.Scan(&action)

	fmt.Println()
	return action
}


func getUserOrder() (order string) {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, 4 - back")
	fmt.Scan(&order)

	return order
}


func getFillAmount(item string) (amount int) {
	fmt.Printf("Write how many %s you want to add:\n", item)
	fmt.Scan(&amount)

	return amount
}


func costOfOrder(order string) (cost int) {
	switch order {
	case "1":
		return 4
	case "2":
		return 7
	case "3":
		return 6
	default:
		return 0
	}
}


func waterRequired(order string) (mlOfWater int) {
	switch order {
	case "1":
		return 250
	case "2":
		return 350
	case "3":
		return 200
	default:
		return 0
	}
}


func milkRequired(order string) (mlOfMilk int) {
	switch order {
	case "1":
		return 0
	case "2":
		return 75
	case "3":
		return 100
	default:
		return 0
	}
}


func coffeeBeansRequired(order string) (gramsOfCoffeeBeans int) {
	switch order {
	case "1":
		return 16
	case "2":
		return 20
	case "3":
		return 12
	default:
		return 0
	}
}
