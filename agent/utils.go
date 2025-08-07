package agent

import (
	"math/rand"
	"sort"
)

func (a agent) randomParticipation() {
    total := 100
	parts := len(*a.assets)

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
	for name, _ := range a.assets_participation {
		a.assets_participation[name] = result[counter]
		counter += 1
	}

}

func (a agent) randomAmounts() {
	for name, _ := range a.assets_amount {
		a.assets_amount[name] = float32(rand.Float64()*99 + 1)
	}
}

func (a agent) mrsCalculator() {
	money_name := (*a.assets)[0]
	for name, _ := range a.assets_MRS {
		quotient_amount := a.assets_amount[money_name] / a.assets_amount[name]
		quotient_participation := a.assets_participation[name] / a.assets_participation[money_name]
		a.assets_MRS[name] = quotient_participation * quotient_amount
	}
}