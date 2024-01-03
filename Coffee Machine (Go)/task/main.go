package main

import "fmt"

func main() {
	// Define resource constants
	mlWaterPerCup := 200
	mlMilkPerCup := 50
	gramsCoffeePerCup := 15

	// Gather the amount of resources the coffee machine has
	var waterSupply int
	var milkSupply int
	var coffeeGramSupply int

	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&waterSupply)

	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&milkSupply)

	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&coffeeGramSupply)

	/* Figure out how many cups of coffee the machine can make.
	This will be determined by the resource that can make the fewest cups.
	*/
	cupsSupply := waterSupply / mlWaterPerCup

	if milkSupply/mlMilkPerCup < cupsSupply {
		cupsSupply = milkSupply / mlMilkPerCup
	}

	if coffeeGramSupply/gramsCoffeePerCup < cupsSupply {
		cupsSupply = coffeeGramSupply / gramsCoffeePerCup
	}

	// Get the requested amount of cups of coffee
	var cupsOrdered int
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cupsOrdered)

	// Respond to the user
	if cupsOrdered < cupsSupply {
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", cupsSupply-cupsOrdered)
	} else if cupsOrdered == cupsSupply {
		fmt.Println("Yes, I can make that amount of coffee")
	} else {
		fmt.Printf("No, I can make only %d cups of coffee\n", cupsSupply)
	}
}
