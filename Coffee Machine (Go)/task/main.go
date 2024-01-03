package main

import "fmt"

func main() {
	var cupCount int
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cupCount)

	waterNeeded := cupCount * 200
	milkNeeded := cupCount * 50
	coffeeNeeded := cupCount * 15

	fmt.Printf("For %d cups of coffee you will need:\n", cupCount)
	fmt.Printf("%d ml of water\n", waterNeeded)
	fmt.Printf("%d ml of milk\n", milkNeeded)
	fmt.Printf("%d g of coffee beans\n", coffeeNeeded)
}
