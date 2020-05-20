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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	logagent "tkestack.io/tke/api/logagent"
)

// FakeLogAgents implements LogAgentInterface
type FakeLogAgents struct {
	Fake *FakeLogagent
}

var logagentsResource = schema.GroupVersionResource{Group: "logagent.tkestack.io", Version: "", Resource: "logagents"}

var logagentsKind = schema.GroupVersionKind{Group: "logagent.tkestack.io", Version: "", Kind: "LogAgent"}

// Get takes name of the logAgent, and returns the corresponding logAgent object, and an error if there is any.
func (c *FakeLogAgents) Get(ctx context.Context, name string, options v1.GetOptions) (result *logagent.LogAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(logagentsResource, name), &logagent.LogAgent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*logagent.LogAgent), err
}

// List takes label and field selectors, and returns the list of LogAgents that match those selectors.
func (c *FakeLogAgents) List(ctx context.Context, opts v1.ListOptions) (result *logagent.LogAgentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(logagentsResource, logagentsKind, opts), &logagent.LogAgentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &logagent.LogAgentList{ListMeta: obj.(*logagent.LogAgentList).ListMeta}
	for _, item := range obj.(*logagent.LogAgentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logAgents.
func (c *FakeLogAgents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(logagentsResource, opts))
}

// Create takes the representation of a logAgent and creates it.  Returns the server's representation of the logAgent, and an error, if there is any.
func (c *FakeLogAgents) Create(ctx context.Context, logAgent *logagent.LogAgent, opts v1.CreateOptions) (result *logagent.LogAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(logagentsResource, logAgent), &logagent.LogAgent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*logagent.LogAgent), err
}

// Update takes the representation of a logAgent and updates it. Returns the server's representation of the logAgent, and an error, if there is any.
func (c *FakeLogAgents) Update(ctx context.Context, logAgent *logagent.LogAgent, opts v1.UpdateOptions) (result *logagent.LogAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(logagentsResource, logAgent), &logagent.LogAgent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*logagent.LogAgent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogAgents) UpdateStatus(ctx context.Context, logAgent *logagent.LogAgent, opts v1.UpdateOptions) (*logagent.LogAgent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(logagentsResource, "status", logAgent), &logagent.LogAgent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*logagent.LogAgent), err
}

// Delete takes name of the logAgent and deletes it. Returns an error if one occurs.
func (c *FakeLogAgents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(logagentsResource, name), &logagent.LogAgent{})
	return err
}

// Patch applies the patch and returns the patched logAgent.
func (c *FakeLogAgents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *logagent.LogAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(logagentsResource, name, pt, data, subresources...), &logagent.LogAgent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*logagent.LogAgent), err
}