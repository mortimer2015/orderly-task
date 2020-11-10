/**
* @Author: Hunter
* @Date: 2020/11/6 15:14
**/
package controller

import (
	"context"
	"fmt"
	batchV1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func jobUpdate(obj interface{}) {
	job := obj.(v1.Object)
	//status := job

	fmt.Printf("Job handle: %s\n", job.GetName())
}

var (
	masterURL  = ""
	kubeconfig = "/Users/hunter/.kube/config"
)

func GetClient() *kubernetes.Clientset {
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}
	return kubeClient
}

func jobCreate() {
	client := GetClient()
	labels := map[string]string{
		"app":        "nginx",
		"controller": "aaa",
	}

	job := batchV1.Job{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "aaa",
			Namespace: "aaa",
		},
		Spec: batchV1.JobSpec{
			//Parallelism: 1,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "aaa",
							Image: "busybox",
						},
					},
				},
			},
		},
	}

	job1, err := client.BatchV1().Jobs("default").Create(context.TODO(), &job, metaV1.CreateOptions{})
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Printf("%s", job1.Name)
	}

}
