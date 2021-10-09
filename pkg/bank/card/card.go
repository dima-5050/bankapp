package card

import (
	"bank/pkg/bank/types"
)

func Total( cards []types.Card) types.Money  {
	sum := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}

		if card.Balance <=0 {
			continue
		}

		sum+=card.Balance
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource  {
	var paymentCardSlice []types.PaymentSource
	var paymentCard types.PaymentSource
	for _, card := range cards {
		if !card.Active {
			continue
		}
		
		if card.Balance <= 0 {
			continue
		}

		paymentCard.Number = string(card.PAN)
		paymentCard.Balance = card.Balance
		paymentCard.Type = card.Name
		paymentCardSlice= append(paymentCardSlice, paymentCard)
	}
	return paymentCardSlice
}