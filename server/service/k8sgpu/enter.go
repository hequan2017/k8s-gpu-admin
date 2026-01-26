package k8sgpu

type ServiceGroup struct {
	ImageRegistryService
	ComputeNodeService
	ProductSpecService
	InstanceService
}
