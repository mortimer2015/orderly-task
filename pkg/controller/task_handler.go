/**
* @Author: Hunter
* @Date: 2020/11/6 15:14
**/
package controller

import (
	"context"
	"strings"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func (c *Controller) taskCreate(obj interface{}) {
	return
}

func (c *Controller) taskUpdate(obj interface{}) {
	oObj := obj.(v1.Object)

	// 非task任务不处理
	if !strings.HasPrefix(oObj.GetName(), BaseTaskName) {
		return
	}

	// 判断是否有正在运行的job
	job, err := c.getRunJob(obj)
	if err != nil {
		klog.Error(err.Error())
		//return
	}
	if job != nil {

		task, err := c.taskClientSet.OrderlytaskV1alpha1().Tasks(oObj.GetNamespace()).Get(context.TODO(), oObj.GetName(),
			metaV1.GetOptions{})
		if err != nil {
			klog.Error("获取task失败")
			klog.Error(err.Error())
		}
		// 有正在运行的job
		if job.Name != "" {
			// task和job对应时，更新job
			if job.Name == BaseTaskName+oObj.GetName() {
				err := c.updateJob(job, task)
				if err != nil {
					klog.Error(err.Error())
				}
			}
		} else {
			// 没有正在运行的job
			err := c.jobCreate(task)
			if err != nil {
				klog.Error("创建任务失败")
				klog.Error(err.Error())
			}
		}
	}

}

func (c *Controller) taskDelete(obj interface{}) {
	oObj := obj.(v1.Object)
	task, err := c.taskClientSet.OrderlytaskV1alpha1().Tasks(oObj.GetNamespace()).Get(context.TODO(), oObj.GetName(),
		metaV1.GetOptions{})
	if err != nil {
		klog.Error("获取task失败")
		klog.Error(err.Error())
	}
	if task != nil {
		job, err := c.kubeClientSet.BatchV1().Jobs(task.Namespace).Get(context.TODO(), BaseTaskName+task.Name,
			metaV1.GetOptions{})
		if err != nil {
			klog.Error("获取Job失败")
			klog.Error(err.Error())
		}

		if job != nil {
			// 移除当前job
			err := c.kubeClientSet.BatchV1().Jobs(job.Namespace).Delete(context.TODO(), job.Name,
				metaV1.DeleteOptions{})
			if err != nil {
				klog.Error("删除job失败")
				klog.Error(err.Error())
			}
		}

	} else {
		klog.Info("没有获取到task")
	}
}
