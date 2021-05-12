/*
Copyright 2020 The Tilt Dev Authors

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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/tilt-dev/tilt/pkg/apis/core/v1alpha1"
)

// FakeUIResources implements UIResourceInterface
type FakeUIResources struct {
	Fake *FakeTiltV1alpha1
}

var uiresourcesResource = schema.GroupVersionResource{Group: "tilt.dev", Version: "v1alpha1", Resource: "uiresources"}

var uiresourcesKind = schema.GroupVersionKind{Group: "tilt.dev", Version: "v1alpha1", Kind: "UIResource"}

// Get takes name of the uIResource, and returns the corresponding uIResource object, and an error if there is any.
func (c *FakeUIResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UIResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(uiresourcesResource, name), &v1alpha1.UIResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UIResource), err
}

// List takes label and field selectors, and returns the list of UIResources that match those selectors.
func (c *FakeUIResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UIResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(uiresourcesResource, uiresourcesKind, opts), &v1alpha1.UIResourceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UIResourceList{ListMeta: obj.(*v1alpha1.UIResourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.UIResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested uIResources.
func (c *FakeUIResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(uiresourcesResource, opts))
}

// Create takes the representation of a uIResource and creates it.  Returns the server's representation of the uIResource, and an error, if there is any.
func (c *FakeUIResources) Create(ctx context.Context, uIResource *v1alpha1.UIResource, opts v1.CreateOptions) (result *v1alpha1.UIResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(uiresourcesResource, uIResource), &v1alpha1.UIResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UIResource), err
}

// Update takes the representation of a uIResource and updates it. Returns the server's representation of the uIResource, and an error, if there is any.
func (c *FakeUIResources) Update(ctx context.Context, uIResource *v1alpha1.UIResource, opts v1.UpdateOptions) (result *v1alpha1.UIResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(uiresourcesResource, uIResource), &v1alpha1.UIResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UIResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUIResources) UpdateStatus(ctx context.Context, uIResource *v1alpha1.UIResource, opts v1.UpdateOptions) (*v1alpha1.UIResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(uiresourcesResource, "status", uIResource), &v1alpha1.UIResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UIResource), err
}

// Delete takes name of the uIResource and deletes it. Returns an error if one occurs.
func (c *FakeUIResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(uiresourcesResource, name), &v1alpha1.UIResource{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUIResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(uiresourcesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UIResourceList{})
	return err
}

// Patch applies the patch and returns the patched uIResource.
func (c *FakeUIResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UIResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(uiresourcesResource, name, pt, data, subresources...), &v1alpha1.UIResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UIResource), err
}
