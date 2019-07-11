package main

import (
	"context"
	"fmt"
	"log"

	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	clientPkg "sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	// Kubernetesの設定
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	scheme := runtime.NewScheme()
	batchv1.AddToScheme(scheme)
	option := clientPkg.Options{
		Scheme: scheme,
	}

	client, err := clientPkg.New(config, option)
	if err != nil {
		fmt.Println(err)
		return
	}

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cowsay",
			Namespace: "default",
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

	ctx := context.Background()

	err = client.Create(ctx, job)
	if err != nil {
		log.Printf("Failed to create job %v\n", err)
	}
}
