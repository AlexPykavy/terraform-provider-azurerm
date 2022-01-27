package endpoints

import (
	"encoding/json"
	"fmt"
)

var _ DeliveryRuleCondition = DeliveryRuleRequestUriCondition{}

type DeliveryRuleRequestUriCondition struct {
	Parameters RequestUriMatchConditionParameters `json:"parameters"`

	// Fields inherited from DeliveryRuleCondition
}

var _ json.Marshaler = DeliveryRuleRequestUriCondition{}

func (s DeliveryRuleRequestUriCondition) MarshalJSON() ([]byte, error) {
	type wrapper DeliveryRuleRequestUriCondition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeliveryRuleRequestUriCondition: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryRuleRequestUriCondition: %+v", err)
	}
	decoded["name"] = "RequestUri"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeliveryRuleRequestUriCondition: %+v", err)
	}

	return encoded, nil
}
