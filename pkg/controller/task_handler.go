/**
* @Author: Hunter
* @Date: 2020/11/6 15:14
**/
package controller

import (
	"context"
	"errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func (c *Controller) taskCreate(obj interface{}) {
	oObj := obj.(metaV1.Object)
	klog.Info("新增task：", oObj.GetName())
	return
}

func (c *Controller) taskUpdate(obj interface{}) {
	oObj := obj.(v1.Object)

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
		if job.Name == BaseTaskName+oObj.GetName() {
			// task和job对应时，更新job
			err := c.updateJob(job, task)
			if err != nil {
				klog.Error(err.Error())
			}
		} else {
			// 判断是否有其他正在运行的job，没有则创建
			c.judgeAndRunJob(obj)
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

func (c *Controller) taskComplete(namespace, taskName string) error {
	task, err := c.taskClientSet.OrderlytaskV1alpha1().Tasks(namespace).Get(context.TODO(), taskName, metaV1.GetOptions{})
	if err != nil {
		return err
	}
	if task != nil {
		if task.Status.Complete == taskComplete {
			return nil
		}

		task.Status.Complete = taskComplete
		_, err := c.taskClientSet.OrderlytaskV1alpha1().Tasks(task.Namespace).Update(context.TODO(), task,
			metaV1.UpdateOptions{})
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return errors.New("未获取到task")
	}
}
