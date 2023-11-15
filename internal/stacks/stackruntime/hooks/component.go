package hooks

import "github.com/hashicorp/terraform/internal/stacks/stackaddrs"

// ComponentInstances is the argument type for the ComponentExpanded hook
// callback, which signals the result of expanding a component into zero or
// more instances.
type ComponentInstances struct {
	ComponentAddr stackaddrs.AbsComponent
	InstanceAddrs []stackaddrs.AbsComponentInstance
}
