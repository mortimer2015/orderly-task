/**
* @Author: Hunter
* @Date: 2020/11/6 15:14
**/
package controller

import (
	"context"
	"errors"
	"fmt"
	batchv1 "k8s.io/api/batch/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	orderlytask "k8s.io/orderly-task/pkg/apis/orderlytask/v1alpha1"
	"strings"
)

func (c *Controller) jobUpdate(obj interface{}) {
	oObj := obj.(metaV1.Object)

	if strings.HasPrefix(oObj.GetName(), BaseTaskName) {
		job, err := c.kubeClientSet.BatchV1().Jobs(oObj.GetNamespace()).Get(context.TODO(), oObj.GetName(), metaV1.GetOptions{})
		if err != nil {
			fmt.Println("jobbbbb")
		}
		if job != nil {
			// todo 重复代码
			complete := false
			for _, condition := range job.Status.Conditions {
				if condition.Type == "Complete" && condition.Status == "True" {
					complete = true
				}
			}
			if complete == false {
				return
			}
			// 获取下一个等待执行的任务
			task, err := c.getNextTask(obj)
			if err != nil {
				klog.Error(err.Error())
				return
			}

			// 创建新的job
			klog.Info("Job update: ", task.Name)
		}
	}
}

func (c *Controller) getNextTask(obj interface{}) (*orderlytask.Task, error) {
	oObj := obj.(metaV1.Object)
	// TODO ListOptions的LabelSelector需要优化
	taskList, err := c.taskClientSet.OrderlytaskV1alpha1().Tasks(oObj.GetNamespace()).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	if taskList != nil {
		var retTask = orderlytask.Task{Spec: orderlytask.TaskSpec{
			Order: int32(2147483647),
		}}
		haveTask := false
		for _, task := range taskList.Items {
			//fmt.Println(task.Name)
			if task.Status.Complete != "true" && task.Spec.Order <= retTask.Spec.Order {
				retTask = task
				haveTask = true
			}
		}
		if haveTask {
			return &retTask, nil
		} else {
			return &retTask, errors.New("没有新task")
		}

	} else {
		return nil, err
	}

}

func jobIsComplete(job batchv1.Job) bool {
	for _, condition := range job.Status.Conditions {
		if condition.Type == "Complete" && condition.Status == "True" {
			return true
		}
	}
	return false
}

func (c *Controller) getRunJob(obj interface{}) (*batchv1.Job, error) {
	oObj := obj.(metaV1.Object)
	jobList, err := c.kubeClientSet.BatchV1().Jobs(oObj.GetNamespace()).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return &batchv1.Job{}, err
	}

	if jobList != nil {
		for _, job := range jobList.Items {
			if strings.HasPrefix(job.Name, BaseTaskName) && !jobIsComplete(job) {
				return &job, nil
			}
		}
	}
	return &batchv1.Job{}, errors.New("没有job正在运行")
}

func (c *Controller) jobCreate(t *orderlytask.Task) error {
	// TODO LabelSelector需要优化
	job := batchv1.Job{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      BaseTaskName + t.Name,
			Namespace: t.Namespace,
		},
		Spec: t.Spec.JobSpec,
	}
	job1, err := c.kubeClientSet.BatchV1().Jobs("default").Create(context.TODO(), &job, metaV1.CreateOptions{})
	if err != nil {
		return err
	} else {
		klog.Info("create success: ", job1.Name)
		return nil
	}
}

func (c *Controller) updateJob(job *batchv1.Job, task *orderlytask.Task) error {
	tmpSelector := job.Spec.Selector
	tmpTemplateLabels := job.Spec.Template.Labels
	job.Spec = task.Spec.JobSpec
	job.Spec.Selector = tmpSelector
	job.Spec.Template.Labels = tmpTemplateLabels

	_, err := c.kubeClientSet.BatchV1().Jobs("default").Update(context.TODO(), job,
		metaV1.UpdateOptions{})
	if err != nil {
		return err
	} else {
		return nil
	}

}
