package card

import (
	"bank/pkg/bank/types"
	"fmt"
)
func ExampleTotal() {
	
	fmt.Println(Total([]types.Card{
	{
		Balance: 1_000_00,
		Active: true,
	},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: true,
		},
		{
			Balance: 2_000_00,
			Active: true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active: true,
		},
	}))
	//Output: 
	//100000
	//300000
	//0
	//0

	
}

func ExamplePaymentSource() {
	payments := []types.Card {
		{
			PAN: "5058 xxxx xxxx 8888",
			Balance: 1_000_00,
			Active: true,	
		},
		{
			PAN: "5058 xxxx xxxx 8887",
			Balance: 1_000_00,
			Active: false,
		},
		{
			PAN: "5058 xxxx xxxx 8886",
			Balance: -1_000_00,
			Active: false,
		},
	}

	num := PaymentSources(payments)
	for _, pay := range num {
		fmt.Println(pay.Number)		
	}
	
	
}