package main

import (
	"log"

	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	jobClient := clientset.BatchV1().Jobs("default")

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-cowsay",
		},
		Spec: batchv1.JobSpec{
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "cowsay",
							Image: "docker/whalesay",
							Command: []string{
								"cowsay",
								"Hello World!",
							},
						},
					},
					RestartPolicy: apiv1.RestartPolicyNever,
				},
			},
		},
	}

	job, err = jobClient.Create(job)
	if err != nil {
		log.Printf("Failed to create job: %v", err)
	}
}

func int32Ptr(i int32) *int32 { return &i }
