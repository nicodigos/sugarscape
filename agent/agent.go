package agent

type assetFeature map[string]float32

type agent struct {
	assets *[]string
	assets_amount assetFeature
	assets_participation assetFeature
	assets_MRS assetFeature
	assets_price_min assetFeature
	assets_price_max assetFeature
}

func NewAgent(assets *[]string) agent {

	var asset_features_zeroes map[string]float32 
	
	for _, name := range *assets {
		asset_features_zeroes[name] = 0
	}

	a := agent{
		assets: assets,
		assets_amount: asset_features_zeroes,
		assets_participation: asset_features_zeroes,
		assets_MRS: asset_features_zeroes,
		assets_price_min: asset_features_zeroes,
		assets_price_max: asset_features_zeroes,
	}

	a.randomParticipation()
	a.randomAmounts()
	a.mrsCalculator()

	a.assets_price_min = a.assets_MRS
	a.assets_price_max = a.assets_MRS
	

	return a
}



// One of the goods is money, so the agent has to calculate his or her total
// budget by converting the total inventory of goods into money
// The problem is that since there is no price stablished (it maybe emerges from the
// transactions of the agents), there is no fixed budget so the first price for
// the first agent is going to be his MRS between money and other goods
// and then is going to be updated for the last price he or she payed in
// grand exchange. But also could be a set of prices, max and min or something like that

// The assets inventory must also contain price for each good, 
// There should be a method for calculate the RMS between money and other goods,
// The good that is going to serve as money should be clarified,

// When the agent goes to the market there should be methods for selling and buying
// The agent will try to buy at the less price seen, but is going to try to sell at 
// what price? That should vary from time to time