package webapplicationfirewallpolicies

type MatchCondition struct {
	MatchValue      []string         `json:"matchValue"`
	MatchVariable   WafMatchVariable `json:"matchVariable"`
	NegateCondition *bool            `json:"negateCondition,omitempty"`
	Operator        Operator         `json:"operator"`
	Selector        *string          `json:"selector,omitempty"`
	Transforms      *[]TransformType `json:"transforms,omitempty"`
}
