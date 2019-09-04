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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/n3wscott/sources/pkg/apis/sources/v1alpha1"
	scheme "github.com/n3wscott/sources/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CronJobSourcesGetter has a method to return a CronJobSourceInterface.
// A group's client should implement this interface.
type CronJobSourcesGetter interface {
	CronJobSources(namespace string) CronJobSourceInterface
}

// CronJobSourceInterface has methods to work with CronJobSource resources.
type CronJobSourceInterface interface {
	Create(*v1alpha1.CronJobSource) (*v1alpha1.CronJobSource, error)
	Update(*v1alpha1.CronJobSource) (*v1alpha1.CronJobSource, error)
	UpdateStatus(*v1alpha1.CronJobSource) (*v1alpha1.CronJobSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CronJobSource, error)
	List(opts v1.ListOptions) (*v1alpha1.CronJobSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CronJobSource, err error)
	CronJobSourceExpansion
}

// cronJobSources implements CronJobSourceInterface
type cronJobSources struct {
	client rest.Interface
	ns     string
}

// newCronJobSources returns a CronJobSources
func newCronJobSources(c *SourcesV1alpha1Client, namespace string) *cronJobSources {
	return &cronJobSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cronJobSource, and returns the corresponding cronJobSource object, and an error if there is any.
func (c *cronJobSources) Get(name string, options v1.GetOptions) (result *v1alpha1.CronJobSource, err error) {
	result = &v1alpha1.CronJobSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cronjobsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CronJobSources that match those selectors.
func (c *cronJobSources) List(opts v1.ListOptions) (result *v1alpha1.CronJobSourceList, err error) {
	result = &v1alpha1.CronJobSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cronjobsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cronJobSources.
func (c *cronJobSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cronjobsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a cronJobSource and creates it.  Returns the server's representation of the cronJobSource, and an error, if there is any.
func (c *cronJobSources) Create(cronJobSource *v1alpha1.CronJobSource) (result *v1alpha1.CronJobSource, err error) {
	result = &v1alpha1.CronJobSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cronjobsources").
		Body(cronJobSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cronJobSource and updates it. Returns the server's representation of the cronJobSource, and an error, if there is any.
func (c *cronJobSources) Update(cronJobSource *v1alpha1.CronJobSource) (result *v1alpha1.CronJobSource, err error) {
	result = &v1alpha1.CronJobSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cronjobsources").
		Name(cronJobSource.Name).
		Body(cronJobSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cronJobSources) UpdateStatus(cronJobSource *v1alpha1.CronJobSource) (result *v1alpha1.CronJobSource, err error) {
	result = &v1alpha1.CronJobSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cronjobsources").
		Name(cronJobSource.Name).
		SubResource("status").
		Body(cronJobSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the cronJobSource and deletes it. Returns an error if one occurs.
func (c *cronJobSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cronjobsources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cronJobSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cronjobsources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cronJobSource.
func (c *cronJobSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CronJobSource, err error) {
	result = &v1alpha1.CronJobSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cronjobsources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
