package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank account
	var bankBalance int
	var wg sync.WaitGroup
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Initial account balance: $%d.00\n", bankBalance)

	// define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made
	for _, income := range incomes {

		go func(income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				// this line causes test to fall in infinite loop
				//fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(income)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("Final bank balance: $%d.00\n", bankBalance)
}