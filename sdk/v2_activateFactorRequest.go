package sdk

type ActivateFactorRequest struct {
	Attestation      string `json:"attestation,omitempty"`
	ClientData       string `json:"clientData,omitempty"`
	PassCode         string `json:"passCode,omitempty"`
	RegistrationData string `json:"registrationData,omitempty"`
	StateToken       string `json:"stateToken,omitempty"`
}

func NewActivateFactorRequest() *ActivateFactorRequest {
	return &ActivateFactorRequest{}
}

func (a *ActivateFactorRequest) IsUserFactorInstance() bool {
	return true
}
