package webapplicationfirewallpolicies

type RateLimitRule struct {
	Action                     ActionType              `json:"action"`
	EnabledState               *CustomRuleEnabledState `json:"enabledState,omitempty"`
	MatchConditions            []MatchCondition        `json:"matchConditions"`
	Name                       string                  `json:"name"`
	Priority                   int64                   `json:"priority"`
	RateLimitDurationInMinutes int64                   `json:"rateLimitDurationInMinutes"`
	RateLimitThreshold         int64                   `json:"rateLimitThreshold"`
}
