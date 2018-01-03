package rmsbase

type RiskRule struct {
	ID           string
	RiskType     string
	InstrumentID string
	Account      string
	StrategyName string
	Indicator    string
	Condition    string
	Threshold    float32
	Action       int
}

const (
	RiskTypeInStrumentID string = "INSTRUMENTID"
	RiskTypeStrategy     string = "STRATEGY"
	RiskTypeAccount      string = "ACCOUNT"

	ActionMail int = 1
	ActionStop int = 2
)
