package agent

import (
	"math/rand"
	"sort"
)

type asset struct {
	amount float32
	participation float32
}

type assetInventory map[string]asset


func newAsset(amount float32, participation float32) asset {
	return asset{amount: amount, participation: participation}
}

func newAssetInventoryRandom(names []string) assetInventory {
	assetInventory := make(map[string]asset)
	participation_list := randomPartition(len(names))
	for i, name := range names {
		assetInventory[name] = asset{amount: float32(rand.Float64()*99 + 1),
									participation: participation_list[i]}
	
} 
return assetInventory
}

func randomPartition(parts int) []float32 {
    total := 100

    if parts <= 0 {
        return []float32{}
    }

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

    return result
}
