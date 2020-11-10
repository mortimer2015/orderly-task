/**
* @Author: Hunter
* @Date: 2020/11/2 11:13
**/
package controller

//import (
//	"fmt"
//	"github.com/golang/glog"
//	coreV1 "k8s.io/api/core/v1"
//	"k8s.io/apimachinery/pkg/api/errors"
//	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/util/runtime"
//	utilRuntime "k8s.io/apimachinery/pkg/util/runtime"
//	"k8s.io/apimachinery/pkg/util/wait"
//	batchInformers "k8s.io/client-go/informers/batch/v1"
//	"k8s.io/client-go/kubernetes"
//	"k8s.io/client-go/kubernetes/scheme"
//	typedCoreV1 "k8s.io/client-go/kubernetes/typed/core/v1"
//	batchLister "k8s.io/client-go/listers/batch/v1"
//	"k8s.io/client-go/tools/cache"
//	"k8s.io/client-go/tools/record"
//	"k8s.io/client-go/util/workqueue"
//	"k8s.io/klog/v2"
//	clientSet "k8s.io/orderly-task/pkg/generated/clientset/versioned"
//	taskScheme "k8s.io/orderly-task/pkg/generated/clientset/versioned/scheme"
//	informers "k8s.io/orderly-task/pkg/generated/informers/externalversions/orderly-task/v1alpha1"
//	taskLister "k8s.io/orderly-task/pkg/generated/listers/orderly-task/v1alpha1"
//	"time"
//)
//
//const controllerAgentName = "control"
//
//const (
//	SuccessSynced         = "Synced"
//	MessageResourceSynced = "Student synced successfully"
//	BaseTaskName          = "Task-d67s5ou4a-"
//)
//
//type Controller struct {
//	kubeClientSet kubernetes.Interface
//	taskClientSet clientSet.Interface
//
//	//deploymentsLister appslisters.DeploymentLister
//	//deploymentsSynced cache.InformerSynced
//	jobLister batchLister.JobLister
//	jobSynced cache.InformerSynced
//
//	controlsLister taskLister.TaskLister
//	controlsSynced cache.InformerSynced
//
//	// workQueue is a rate limited work queue. This is used to queue work to be
//	// processed instead of performing it as soon as a change happens. This
//	// means we can ensure we only process a fixed amount of resources at a
//	// time, and makes it easy to ensure we are never processing the same item
//	// simultaneously in two different workers.
//	workQueue workqueue.RateLimitingInterface
//	// recorder is an event recorder for recording Event resources to the
//	// Kubernetes API.
//	recorder record.EventRecorder
//}
//
//// NewController returns a new sample controller
//func NewController(
//	kubeClientSet kubernetes.Interface,
//	taskClientSet clientSet.Interface,
//	jobInformer batchInformers.JobInformer,
//	taskInformer informers.TaskInformer) *Controller {
//
//	// Create event broadcaster
//	// Add sample-controller types to the default Kubernetes Scheme so Events can be
//	// logged for sample-controller types.
//	utilRuntime.Must(taskScheme.AddToScheme(scheme.Scheme))
//	klog.V(4).Info("Creating event broadcaster")
//	eventBroadcaster := record.NewBroadcaster()
//	eventBroadcaster.StartStructuredLogging(0)
//	eventBroadcaster.StartRecordingToSink(&typedCoreV1.EventSinkImpl{Interface: kubeClientSet.CoreV1().Events("")})
//	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, coreV1.EventSource{Component: controllerAgentName})
//
//	controller := &Controller{
//		kubeClientSet:  kubeClientSet,
//		taskClientSet:  taskClientSet,
//		jobLister:      jobInformer.Lister(),
//		jobSynced:      jobInformer.Informer().HasSynced,
//		controlsLister: taskInformer.Lister(),
//		controlsSynced: taskInformer.Informer().HasSynced,
//		workQueue:      workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Foos"),
//		recorder:       recorder,
//	}
//
//	klog.Info("Setting up event handlers")
//
//	taskInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
//		AddFunc: controller.taskCreate,
//		UpdateFunc: func(old, new interface{}) {
//			controller.taskUpdate(new)
//		},
//		DeleteFunc: controller.delete,
//	})
//
//	jobInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
//		AddFunc: controller.add,
//		UpdateFunc: func(old, new interface{}) {
//			jobUpdate(new)
//		},
//		DeleteFunc: controller.delete,
//	})
//
//	return controller
//}
//
//func (c *Controller) add(obj interface{}) {
//	oObj := obj.(v1.Object)
//	fmt.Printf("add task: %s\n", oObj.GetName())
//}
//
//func (c *Controller) update(obj interface{}) {
//	oObj := obj.(v1.Object)
//	fmt.Printf("update task: %s\n", oObj.GetName())
//}
//
//func (c *Controller) delete(obj interface{}) {
//	oObj := obj.(v1.Object)
//	fmt.Printf("delete task: %s\n", oObj.GetName())
//}
//
//func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
//	defer utilRuntime.HandleCrash()
//	defer c.workQueue.ShutDown()
//
//	// Start the informer factories to begin populating the informer caches
//	klog.Info("Starting Foo controller")
//
//	// Wait for the caches to be synced before starting workers
//	klog.Info("Waiting for informer caches to sync")
//	if ok := cache.WaitForCacheSync(stopCh, c.controlsSynced); !ok {
//		return fmt.Errorf("failed to wait for caches to sync")
//	}
//
//	klog.Info("Starting workers")
//	// Launch two workers to process Foo resources
//	for i := 0; i < threadiness; i++ {
//		go wait.Until(c.runWorker, time.Second, stopCh)
//	}
//
//	klog.Info("Started workers")
//	<-stopCh
//	klog.Info("Shutting down workers")
//
//	return nil
//}
//
//func (c *Controller) runWorker() {
//	for c.processNextWorkItem() {
//	}
//}
//
//func (c *Controller) processNextWorkItem() bool {
//	obj, shutdown := c.workQueue.Get()
//
//	if shutdown {
//		return false
//	}
//
//	err := func(obj interface{}) error {
//		defer c.workQueue.Done(obj)
//		var key string
//		var ok bool
//		if key, ok = obj.(string); !ok {
//			c.workQueue.Forget(obj)
//			utilRuntime.HandleError(fmt.Errorf("expected string in workQueue but got %#v", obj))
//			return nil
//		}
//		if err := c.syncHandler(key); err != nil {
//			c.workQueue.AddRateLimited(key)
//			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
//		}
//		c.workQueue.Forget(obj)
//		klog.Infof("Successfully synced '%s'", key)
//		return nil
//	}(obj)
//
//	if err != nil {
//		utilRuntime.HandleError(err)
//		return true
//	}
//
//	return true
//}
//
//func (c *Controller) syncHandler(key string) error {
//	// Convert the namespace/name string into a distinct namespace and name
//	namespace, name, err := cache.SplitMetaNamespaceKey(key)
//	if err != nil {
//		runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
//		return nil
//	}
//
//	// 从缓存中取对象
//	student, err := c.controlsLister.Tasks(namespace).Get(name)
//	if err != nil {
//		// 如果Student对象被删除了，就会走到这里，所以应该在这里加入执行
//		if errors.IsNotFound(err) {
//			glog.Infof("Student对象被删除，请在这里执行实际的删除业务: %s/%s ...", namespace, name)
//
//			return nil
//		}
//
//		runtime.HandleError(fmt.Errorf("failed to list student by: %s/%s", namespace, name))
//
//		return err
//	}
//
//	glog.Infof("这里是student对象的期望状态: %#v ...", student)
//	glog.Infof("实际状态是从业务层面得到的，此处应该去的实际状态，与期望状态做对比，并根据差异做出响应(新增或者删除)")
//
//	c.recorder.Event(student, coreV1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
//	return nil
//}
