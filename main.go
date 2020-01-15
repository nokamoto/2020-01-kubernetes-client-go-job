package main

import (
	batch "k8s.io/api/batch/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	println(cfg.String())

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	client := clientset.BatchV1().Jobs("default")
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#running-an-example-job
	var backoffLimit int32 = 4
	job := &batch.Job{
		TypeMeta: v1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "pi",
		},
		Spec: batch.JobSpec{
			BackoffLimit: &backoffLimit,
			Template: v12.PodTemplateSpec{
				Spec: v12.PodSpec{
					Containers: []v12.Container{
						{
							Name:    "pi",
							Image:   "perl",
							Command: []string{"perl", "-Mbignum=bpi", "-wle", "print bpi(2000)"},
						},
					},
					RestartPolicy: "Never",
				},
			},
		},
	}

	job, err = client.Create(job)
	// if err != nil {
	// 	panic(err) // (?) panic: jobs.batch "pi" already exists
	// }

	println(job)
	println(err)
}
