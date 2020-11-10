/**
* @Author: Hunter
* @Date: 2020/11/6 15:14
**/
package controller

import (
	"context"
	//"k8s.io/orderlytask/pkg/apis/orderlytask/v1alpha1"
	batchV1 "k8s.io/api/batch/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func (c *Controller) taskCreate(obj interface{}) {
	//labels := map[string]string{
	//	"app":        "orderlytask",
	//	"controller": "orderlytask",
	//}
	//parallelism := int32(1)
	oObj := obj.(v1.Object)
	task, err := c.taskClientSet.OrderlytaskV1alpha1().Tasks(oObj.GetNamespace()).Get(context.TODO(), oObj.GetName(), metaV1.GetOptions{})
	if err != nil {
		klog.Error("获取task失败")
		klog.Error(err.Error())
	}
	if task != nil {
		job := batchV1.Job{
			ObjectMeta: metaV1.ObjectMeta{
				Name:      BaseTaskName + task.Name,
				Namespace: task.Namespace,
			},
			Spec: task.Spec,
		}
		job1, err := c.kubeClientSet.BatchV1().Jobs("default").Create(context.TODO(), &job, metaV1.CreateOptions{})
		if err != nil {
			klog.Error("创建job失败")
			klog.Error(err.Error())
		} else {
			klog.Info("create %s success", job1.Name)
		}
	} else {
		klog.Info("没有获取到task")
	}

}

func (c *Controller) taskUpdate(obj interface{}) {
	//labels := map[string]string{
	//	"app":        "orderlytask",
	//	"controller": "orderlytask",
	//}
	//parallelism := int32(1)
	oObj := obj.(v1.Object)
	task, err := c.taskClientSet.OrderlytaskV1alpha1().Tasks(oObj.GetNamespace()).Get(context.TODO(), oObj.GetName(), metaV1.GetOptions{})
	if err != nil {
		klog.Error("获取task失败")
		klog.Error(err.Error())
	}
	if task != nil {
		job := batchV1.Job{
			ObjectMeta: metaV1.ObjectMeta{
				Name:      BaseTaskName + task.Name,
				Namespace: task.Namespace,
			},
			Spec: task.Spec,
		}
		job1, err := c.kubeClientSet.BatchV1().Jobs("default").Update(context.TODO(), &job, metaV1.UpdateOptions{})
		if err != nil {
			klog.Error("创建job失败")
			klog.Error(err.Error())
		} else {
			klog.Info("create %s success", job1.Name)
		}
	} else {
		klog.Info("没有获取到task")
	}

}

//Spec: batchV1.JobSpec{
//	Parallelism: task.Spec.Parallelism,
//	Completions: task.Spec.Completions,
//	ActiveDeadlineSeconds: task.Spec.ActiveDeadlineSeconds,
//	BackoffLimit: task.Spec.BackoffLimit,
//	Selector: task.Spec.Selector,
//	ManualSelector: task.Spec.ManualSelector,
//	Template: task.Spec.Template,
//	TTLSecondsAfterFinished: task.Spec.TTLSecondsAfterFinished,
//	//Template: corev1.PodTemplateSpec{
//	//	ObjectMeta: metaV1.ObjectMeta{
//	//		Labels: labels,
//	//	},
//	//	Spec: corev1.PodSpec{
//	//		Containers: []corev1.Container{
//	//			{
//	//				Name: "aaa",
//	//				Image: "busybox",
//	//				Command: []string{"sleep", "120"},
//	//			},
//	//		},
//	//		RestartPolicy: "OnFailure",
//	//	},
//	//},
//},
