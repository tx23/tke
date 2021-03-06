/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	platform "tkestack.io/tke/api/platform"
)

// RegistriesGetter has a method to return a RegistryInterface.
// A group's client should implement this interface.
type RegistriesGetter interface {
	Registries() RegistryInterface
}

// RegistryInterface has methods to work with Registry resources.
type RegistryInterface interface {
	Create(ctx context.Context, registry *platform.Registry, opts v1.CreateOptions) (*platform.Registry, error)
	Update(ctx context.Context, registry *platform.Registry, opts v1.UpdateOptions) (*platform.Registry, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*platform.Registry, error)
	List(ctx context.Context, opts v1.ListOptions) (*platform.RegistryList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *platform.Registry, err error)
	RegistryExpansion
}

// registries implements RegistryInterface
type registries struct {
	client rest.Interface
}

// newRegistries returns a Registries
func newRegistries(c *PlatformClient) *registries {
	return &registries{
		client: c.RESTClient(),
	}
}

// Get takes name of the registry, and returns the corresponding registry object, and an error if there is any.
func (c *registries) Get(ctx context.Context, name string, options v1.GetOptions) (result *platform.Registry, err error) {
	result = &platform.Registry{}
	err = c.client.Get().
		Resource("registries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Registries that match those selectors.
func (c *registries) List(ctx context.Context, opts v1.ListOptions) (result *platform.RegistryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &platform.RegistryList{}
	err = c.client.Get().
		Resource("registries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested registries.
func (c *registries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("registries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a registry and creates it.  Returns the server's representation of the registry, and an error, if there is any.
func (c *registries) Create(ctx context.Context, registry *platform.Registry, opts v1.CreateOptions) (result *platform.Registry, err error) {
	result = &platform.Registry{}
	err = c.client.Post().
		Resource("registries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(registry).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a registry and updates it. Returns the server's representation of the registry, and an error, if there is any.
func (c *registries) Update(ctx context.Context, registry *platform.Registry, opts v1.UpdateOptions) (result *platform.Registry, err error) {
	result = &platform.Registry{}
	err = c.client.Put().
		Resource("registries").
		Name(registry.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(registry).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the registry and deletes it. Returns an error if one occurs.
func (c *registries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("registries").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched registry.
func (c *registries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *platform.Registry, err error) {
	result = &platform.Registry{}
	err = c.client.Patch(pt).
		Resource("registries").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
