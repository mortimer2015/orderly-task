# orderly-task

#### 介绍
基于k8s CRD功能实现有序job功能



#### 软件架构
软件架构说明： 略🥶


#### 调试

1.  mkdir $GOPATH/src/k8s.io
2.  cd $GOPATH/src/k8s.io
3.  git clone https://github.com/mortimer2015/orderly-task.git
4.  cd orderly-task
5.  go mod vendor
6.  go run cmd/orderly-task/main.go --master="" --kubeconfig="~/.kube/config"

#### 使用说明

1.  日志输出到了标准输出，需要保存日志的话，使用 >>输出到文件
2.  crd文件参考artifacts/crd.yaml
3.  task创建示例参考artifacts/example-foo5.yaml和artifacts/example-foo10.yaml
4.  如artifacts/example-foo5.yaml的中的`order`是定义task的执行顺序，由小到大逐步执行，`jobSpec`的格式和Job的Spec定义格式完全一致
5.  如果想使用多个不同task组，只需在不同的命名空间下创建task即可，本controller会自动按照不同的命名空间进行各自task的调度
6.  mac版二进制包下载，请查看Release
7.  创建CRD`kubectl apply -f artifacts/crd.yaml`
8.  创建task`kubectl apply -f artifacts/example-foo5.yaml`
9.  查看task`kubectl get task`

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

