package sdk

import (
	"encoding/json"
)

type GroupProfileMap map[string]interface{}

type GroupProfile struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	GroupProfileMap
}

func (a *GroupProfile) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	var profile map[string]interface{}
	err := json.Unmarshal(data, &profile)
	if err != nil {
		return err
	}
	a.Name, _ = profile["name"].(string)
	a.Description, _ = profile["description"].(string)
	delete(profile, "name")
	delete(profile, "description")
	a.GroupProfileMap = profile
	return nil
}

func (a GroupProfile) MarshalJSON() ([]byte, error) {
	if len(a.GroupProfileMap) == 0 {
		return json.Marshal(&struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		}{
			Name:        a.Name,
			Description: a.Description,
		})
	}
	if a.Name != "" {
		a.GroupProfileMap["name"] = a.Name
	}
	if a.Description != "" {
		a.GroupProfileMap["description"] = a.Description
	}
	return json.Marshal(a.GroupProfileMap)
}
