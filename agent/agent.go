package agent

type agent struct {
	assets assetInventory
}

func NewAgent(names []string) agent {
	return agent{assets: newAssetInventoryRandom(names)}
}


func (a agent) redoParticipationAssets() {
	new_participation_list := randomPartition(len(a.assets))

	counter := 0
	for key := range a.assets {
		asset := a.assets[key]
		asset.participation = new_participation_list[counter]
		a.assets[key] = asset 
		counter += 1
	}
}

// One of the goods is money, so the agent has to calculate his or her total
// budget by converting the total inventory of goods into money
// The problem is that since there is no price stablished (it maybe emerges from the
// transactions of the agents), there is no fixed budget so the first price for
// the first agent is going to be his RMS between money and other goods
// and then is going to be updated for the last price he or she payed in
// grand exchange. But also could be a set of prices, max and min or something like that

// The assets inventory must also contain price for each good, 
// There should be a method for calculate the RMS between money and other goods,
// The good that is going to serve as money should be clarified,

// When the agent goes to the market there should be methods for selling and buying
// The agent will try to buy at the less price seen, but is going to try to sell at 
// what price? That should vary from time to time