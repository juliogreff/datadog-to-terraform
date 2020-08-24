package types

type ResourceWrapper struct {
	Resource `hcl:"resource"`
}

type Resource struct {
	Type     string `hcl:",key"`
	Name     string `hcl:",key"`
	*Board   `hcl:"dashboard,squash"`
	*Monitor `hcl:"monitor,squash"`
}
