package agent

type bundleFeature map[string]float32

type agent struct {
	budget float32
	price_rule float32
	bundle *[]string
	bundle_amount bundleFeature
	bundle_participation bundleFeature
	bundle_MRS bundleFeature
	bundle_price_min bundleFeature
	bundle_price_max bundleFeature
	bundle_price_average bundleFeature
	bundle_optimal bundleFeature
}

func NewAgent(bundles *[]string) agent {

	var bundle_features_zeroes map[string]float32 
	
	for _, name := range *bundles {
		bundle_features_zeroes[name] = 0
	}

	a := agent{
		budget: 0,
		price_rule: 0.5,
		bundle: bundles,
		bundle_amount: bundle_features_zeroes,
		bundle_participation: bundle_features_zeroes,
		bundle_MRS: bundle_features_zeroes,
		bundle_price_min: bundle_features_zeroes,
		bundle_price_max: bundle_features_zeroes,
		
	}

	a.randomParticipation()
	a.randomAmounts()
	a.mrsCalculator()
	a.bundle_price_min = a.bundle_MRS
	a.bundle_price_max = a.bundle_MRS
	a.bundle_price_average = a.bundle_MRS
	a.calculateBudget()
	a.calculateOptimalAmount()

	return a
}



// One of the goods is money, so the agent has to calculate his or her total
// budget by converting the total inventory of goods into money
// The problem is that since there is no price stablished (it maybe emerges from the
// transactions of the agents), there is no fixed budget so the first price for
// the first agent is going to be his MRS between money and other goods
// and then is going to be updated for the last price he or she payed in
// grand exchange. But also could be a set of prices, max and min or something like that

// The bundles inventory must also contain price for each good, 
// There should be a method for calculate the RMS between money and other goods,
// The good that is going to serve as money should be clarified,

// When the agent goes to the market there should be methods for selling and buying
// The agent will try to buy at the less price seen, but is going to try to sell at 
// what price? That should vary from time to time