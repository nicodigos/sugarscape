package agent

import (
	"math/rand"
	"sort"
)

func (a agent) randomParticipation() {
    total := 100
	parts := len(*a.bundle)

    cuts := make([]int, parts-1)
    for i := 0; i < parts-1; i++ {
        cuts[i] = rand.Intn(total + 1)
    }

    cuts = append(cuts, 0, total)
    sort.Ints(cuts)

    result := make([]float32, parts)
    for i := 0; i < parts; i++ {
        chunk := cuts[i+1] - cuts[i]
        result[i] = float32(chunk) / 100.0 
    }

	counter := 0
	for name, _ := range a.bundle_participation {
		a.bundle_participation[name] = result[counter]
		counter += 1
	}

}

func (a agent) randomAmounts() {
	for name, _ := range a.bundle_amount {
		a.bundle_amount[name] = float32(rand.Float64()*99 + 1)
	}
}

func (a agent) mrsCalculator() {
	money_name := (*a.bundle)[0]
	for name, _ := range a.bundle_MRS {
		quotient_amount := a.bundle_amount[money_name] / a.bundle_amount[name]
		quotient_participation := a.bundle_participation[name] / a.bundle_participation[money_name]
		a.bundle_MRS[name] = quotient_participation * quotient_amount
	}
}

func (a agent) calculateAveragePrice() {
	for name, _ := range a.bundle_price_max {
		price := (a.bundle_price_max[name] * a.price_rule) + (a.bundle_price_min[name] * a.price_rule)
		a.bundle_price_average[name] = price 
	}
}

func (a agent) calculateBudget() {
	var budget float32 
	for name, _ := range a.bundle_price_max {
		budget += a.bundle_price_average[name] * a.bundle_amount[name]
	}

	a.budget = budget	         
}

func (a agent) calculateOptimalAmount() {
	for name, _ := range a.bundle_optimal {
		bundle_optimal_money := a.budget * a.bundle_participation[name]
		a.bundle_optimal[name] = bundle_optimal_money / a.bundle_price_average[name]
	}
}