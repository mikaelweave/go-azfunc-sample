package playground

import "context"

// ListPlaygrounds returns a list of all active playgrounds in subscription
func ListPlaygrounds(ctx context.Context) ([]Playground, error) {
	groupsClient, err := getResourceGroupsClient()
	if err != nil {
		return nil, err
	}

	// Get resource groups created by our Playground System
	groups, err := groupsClient.ListComplete(ctx, "tagName eq 'System' and tagValue eq 'Playground'", nil)
	if err != nil {
		return nil, err
	}

	playgrounds := make([]Playground, 0)
	for _, v := range *groups.Response().Value {
		playgrounds = append(playgrounds, groupToPlayground(v))
	}

	return playgrounds, nil
}
