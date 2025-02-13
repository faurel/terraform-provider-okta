package sdk

import "encoding/json"

type OktaSignOnPolicyRuleSignonSessionActions struct {
	MaxSessionIdleMinutes        int64  `json:"-"`
	MaxSessionIdleMinutesPtr     *int64 `json:"maxSessionIdleMinutes"`
	MaxSessionLifetimeMinutes    int64  `json:"-"`
	MaxSessionLifetimeMinutesPtr *int64 `json:"maxSessionLifetimeMinutes"`
	UsePersistentCookie          *bool  `json:"usePersistentCookie,omitempty"`
}

func NewOktaSignOnPolicyRuleSignonSessionActions() *OktaSignOnPolicyRuleSignonSessionActions {
	return &OktaSignOnPolicyRuleSignonSessionActions{
		MaxSessionIdleMinutes:        120,
		MaxSessionIdleMinutesPtr:     Int64Ptr(120),
		MaxSessionLifetimeMinutes:    0,
		MaxSessionLifetimeMinutesPtr: Int64Ptr(0),
		UsePersistentCookie:          boolPtr(false),
	}
}

func (a *OktaSignOnPolicyRuleSignonSessionActions) IsPolicyInstance() bool {
	return true
}

func (a *OktaSignOnPolicyRuleSignonSessionActions) MarshalJSON() ([]byte, error) {
	type Alias OktaSignOnPolicyRuleSignonSessionActions
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.MaxSessionIdleMinutes != 0 {
		result.MaxSessionIdleMinutesPtr = Int64Ptr(a.MaxSessionIdleMinutes)
	}
	if a.MaxSessionLifetimeMinutes != 0 {
		result.MaxSessionLifetimeMinutesPtr = Int64Ptr(a.MaxSessionLifetimeMinutes)
	}
	return json.Marshal(&result)
}

func (a *OktaSignOnPolicyRuleSignonSessionActions) UnmarshalJSON(data []byte) error {
	type Alias OktaSignOnPolicyRuleSignonSessionActions

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.MaxSessionIdleMinutesPtr != nil {
		a.MaxSessionIdleMinutes = *result.MaxSessionIdleMinutesPtr
		a.MaxSessionIdleMinutesPtr = result.MaxSessionIdleMinutesPtr
	}
	if result.MaxSessionLifetimeMinutesPtr != nil {
		a.MaxSessionLifetimeMinutes = *result.MaxSessionLifetimeMinutesPtr
		a.MaxSessionLifetimeMinutesPtr = result.MaxSessionLifetimeMinutesPtr
	}
	return nil
}
