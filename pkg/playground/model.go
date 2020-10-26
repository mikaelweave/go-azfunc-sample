package playground

import (
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
)

// Playground is a resource group for our system
type Playground struct {
	ID       *string            `json:"id,omitempty"`
	Name     *string            `json:"name,omitempty"`
	Location *string            `json:"location,omitempty"`
	OwnerID  *string            `json:"ownerId,omitempty"`
	Tags     map[string]*string `json:"tags"`
}

func groupToPlayground(g resources.Group) Playground {
	var p Playground
	p.ID = g.ID
	p.Name = g.Name
	p.Location = g.Location
	p.Tags = g.Tags

	if val, ok := g.Tags["OwnerId"]; ok {
		p.OwnerID = val
	}

	return p
}
