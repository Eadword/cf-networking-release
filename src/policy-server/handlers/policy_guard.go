package handlers

import (
	"fmt"
	"policy-server/store"
	"policy-server/uaa_client"
)

type PolicyGuard struct {
	CCClient  ccClient
	UAAClient uaaClient
}

func NewPolicyGuard(uaaClient uaaClient, ccClient ccClient) *PolicyGuard {
	return &PolicyGuard{
		CCClient:  ccClient,
		UAAClient: uaaClient,
	}
}

func (g *PolicyGuard) CheckAccess(policyCollection store.PolicyCollection, userToken uaa_client.CheckTokenResponse) (bool, error) {
	for _, scope := range userToken.Scope {
		if scope == "network.admin" {
			return true, nil
		}
	}

	if len(policyCollection.EgressPolicies) > 0 {
		return false, nil
	}

	token, err := g.UAAClient.GetToken()
	if err != nil {
		return false, fmt.Errorf("getting token: %s", err)
	}

	spaceGUIDs, err := g.CCClient.GetSpaceGUIDs(token, uniqueAppGUIDs(policyCollection.Policies))
	if err != nil {
		return false, fmt.Errorf("getting space guids: %s", err)
	}
	for _, guid := range spaceGUIDs {
		space, err := g.CCClient.GetSpace(token, guid)
		if err != nil {
			return false, fmt.Errorf("getting space with guid %s: %s", guid, err)
		}
		if space == nil {
			return false, nil
		}
		userSpace, err := g.CCClient.GetUserSpace(token, userToken.UserID, *space)
		if err != nil {
			return false, fmt.Errorf("getting space with guid %s: %s", guid, err)
		}
		if userSpace == nil {
			return false, nil
		}
	}
	return true, nil
}
func (g *PolicyGuard) CheckEgressPolicyListAccess(userToken uaa_client.CheckTokenResponse) bool {
	for _, scope := range userToken.Scope {
		if scope == "network.admin" {
			return true
		}
	}

	return false
}

func uniqueAppGUIDs(policies []store.Policy) []string {
	var set = make(map[string]struct{})
	for _, policy := range policies {
		set[policy.Source.ID] = struct{}{}
		set[policy.Destination.ID] = struct{}{}
	}
	var appGUIDs = make([]string, 0, len(set))
	for guid, _ := range set {
		appGUIDs = append(appGUIDs, guid)
	}
	return appGUIDs
}
