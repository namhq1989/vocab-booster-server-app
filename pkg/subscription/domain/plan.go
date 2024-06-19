package domain

type Plan string

const (
	PlanUnknown          Plan = ""
	PlanFree             Plan = "free"
	PlanPremiumMonthly   Plan = "premiumMonthly"
	PlanPremiumQuarterly Plan = "premiumQuarterly"
	PlanPremiumYearly    Plan = "premiumYearly"
)

func (p Plan) String() string {
	return string(p)
}

func (p Plan) IsValid() bool {
	return p != PlanUnknown
}

func ToPlan(value string) Plan {
	switch value {
	case PlanFree.String():
		return PlanFree
	case PlanPremiumMonthly.String():
		return PlanPremiumMonthly
	case PlanPremiumQuarterly.String():
		return PlanPremiumQuarterly
	case PlanPremiumYearly.String():
		return PlanPremiumYearly
	default:
		return PlanUnknown
	}
}

func (p Plan) IsFree() bool {
	return p == PlanFree
}

func (p Plan) IsPremium() bool {
	return p.IsValid() && !p.IsFree()
}
