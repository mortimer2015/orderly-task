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
	oObj := obj.(v1.Object)
	task, err := c.taskClientSet.OrderlytaskV1alpha1().Tasks(oObj.GetNamespace()).Get(context.TODO(), oObj.GetName(), metaV1.GetOptions{})
	if err != nil {
		klog.Error("获取task失败")
		klog.Error(err.Error())
	}
	if task != nil {
		// TODO 判断是否需要创建job

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
			klog.Fatalf("create %s success", job1.Name)
		}
	} else {
		klog.Info("没有获取到task")
	}
}

func (c *Controller) taskUpdate(obj interface{}) {
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
			tmpSelector := job.Spec.Selector
			tmpTemplateLabels := job.Spec.Template.Labels
			job.Spec = task.Spec
			job.Spec.Selector = tmpSelector
			job.Spec.Template.Labels = tmpTemplateLabels

			job1, err := c.kubeClientSet.BatchV1().Jobs("default").Update(context.TODO(), job,
				metaV1.UpdateOptions{})
			if err != nil {
				klog.Error("更新job失败")
				klog.Error(err.Error())
			} else {
				klog.Info(job1.Name, " 更新成功")
			}
		}

	} else {
		klog.Info("没有获取到task")
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

			err := c.kubeClientSet.BatchV1().Jobs(job.Namespace).Delete(context.TODO(), job.Name,
				metaV1.DeleteOptions{})
			if err != nil {
				klog.Error("删除job失败")
				klog.Error(err.Error())
			}
		}
		// TODO 正在运行的task

		// TODO 移除当前job

		// TODO 下一个任务

	} else {
		klog.Info("没有获取到task")
	}
}
