/**
* @Author: Hunter
* @Date: 2020/11/6 15:14
**/
package controller

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func (c *Controller) jobUpdate(obj interface{}) {
	oObj := obj.(metaV1.Object)

	job, err := c.kubeClientSet.BatchV1().Jobs(oObj.GetNamespace()).Get(context.TODO(), oObj.GetName(), metaV1.GetOptions{})
	if err != nil {
		fmt.Println("jobbbbb")
	}
	if job != nil {
		if strings.HasPrefix(job.Name, BaseTaskName) {
			complete := false

			for _, condition := range job.Status.Conditions {
				if condition.Type == "Complete" && condition.Status == "True" {
					complete = true
				}
			}
			if complete == false {
				return
			}
			//task := c.getNextTask(obj)
			//fmt.Println(task.Name)
		}
		//orderlytask.Task
		fmt.Printf("Job update: %s\n", job.GetName())
	}

}

//func (c *Controller) getNextTask(obj interface{}) (task *orderlytask.Task) {
//	selector := labels.Selector()
//	tasks, err := c.controlsLister.Tasks(obj.Namespace).List(&selector)
//	if err != nil {
//
//	}
//	for task := range(tasks){
//		fmt.Println(task.Name)
//	}
//	return task
//}
