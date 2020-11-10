module k8s.io/Orderly_task

go 1.15

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	k8s.io/api v0.19.3
	k8s.io/apimachinery v0.19.3
	k8s.io/client-go v0.19.3
	k8s.io/code-generator v0.19.3
	k8s.io/klog/v2 v2.2.0
)

replace (
	k8s.io/api => k8s.io/api v0.19.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.3
	k8s.io/client-go => k8s.io/client-go v0.19.3
	k8s.io/code-generator => k8s.io/code-generator v0.19.3
)
