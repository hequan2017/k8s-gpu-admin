package k8sgpu

type RouterGroup struct {
	ImageRegistryRouter
	ComputeNodeRouter
	ProductSpecRouter
	InstanceRouter
}
