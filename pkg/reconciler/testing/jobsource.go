/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testing

import (
	"context"

	"github.com/n3wscott/sources/pkg/apis/sources/v1alpha1"
	"github.com/n3wscott/sources/pkg/reconciler/jobsource/resources"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobSourceOption func(*v1alpha1.JobSource)

func NewJobSource(name string, options ...JobSourceOption) *v1alpha1.JobSource {
	js := &v1alpha1.JobSource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
		},
	}

	for _, option := range options {
		option(js)
	}

	js.SetDefaults(context.Background())
	return js
}

func WithFakeJobContainer(js *v1alpha1.JobSource) {
	js.Spec.Template.Spec.Containers = append(js.Spec.Template.Spec.Containers, corev1.Container{
		Name:  "Steve",
		Image: "grc.io/fakeimage",
	})
}

type JobOption func(*batchv1.Job)

func NewJob(js *v1alpha1.JobSource, options ...JobOption) *batchv1.Job {
	job := resources.MakeJob(js)

	for _, option := range options {
		option(job)
	}

	return job
}
