package landscape

import "github.com/nicodigos/sugarscape/agent"

type Cell struct {
	X int
	Y int
	Agent *agent.Agent
	MaxSugar int
	CurrentSugar int
}