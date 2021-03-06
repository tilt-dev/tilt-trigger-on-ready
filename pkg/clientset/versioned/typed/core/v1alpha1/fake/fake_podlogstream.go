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

// FakePodLogStreams implements PodLogStreamInterface
type FakePodLogStreams struct {
	Fake *FakeTiltV1alpha1
}

var podlogstreamsResource = schema.GroupVersionResource{Group: "tilt.dev", Version: "v1alpha1", Resource: "podlogstreams"}

var podlogstreamsKind = schema.GroupVersionKind{Group: "tilt.dev", Version: "v1alpha1", Kind: "PodLogStream"}

// Get takes name of the podLogStream, and returns the corresponding podLogStream object, and an error if there is any.
func (c *FakePodLogStreams) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PodLogStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(podlogstreamsResource, name), &v1alpha1.PodLogStream{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodLogStream), err
}

// List takes label and field selectors, and returns the list of PodLogStreams that match those selectors.
func (c *FakePodLogStreams) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PodLogStreamList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(podlogstreamsResource, podlogstreamsKind, opts), &v1alpha1.PodLogStreamList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PodLogStreamList{ListMeta: obj.(*v1alpha1.PodLogStreamList).ListMeta}
	for _, item := range obj.(*v1alpha1.PodLogStreamList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podLogStreams.
func (c *FakePodLogStreams) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(podlogstreamsResource, opts))
}

// Create takes the representation of a podLogStream and creates it.  Returns the server's representation of the podLogStream, and an error, if there is any.
func (c *FakePodLogStreams) Create(ctx context.Context, podLogStream *v1alpha1.PodLogStream, opts v1.CreateOptions) (result *v1alpha1.PodLogStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(podlogstreamsResource, podLogStream), &v1alpha1.PodLogStream{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodLogStream), err
}

// Update takes the representation of a podLogStream and updates it. Returns the server's representation of the podLogStream, and an error, if there is any.
func (c *FakePodLogStreams) Update(ctx context.Context, podLogStream *v1alpha1.PodLogStream, opts v1.UpdateOptions) (result *v1alpha1.PodLogStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(podlogstreamsResource, podLogStream), &v1alpha1.PodLogStream{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodLogStream), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodLogStreams) UpdateStatus(ctx context.Context, podLogStream *v1alpha1.PodLogStream, opts v1.UpdateOptions) (*v1alpha1.PodLogStream, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(podlogstreamsResource, "status", podLogStream), &v1alpha1.PodLogStream{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodLogStream), err
}

// Delete takes name of the podLogStream and deletes it. Returns an error if one occurs.
func (c *FakePodLogStreams) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(podlogstreamsResource, name), &v1alpha1.PodLogStream{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodLogStreams) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(podlogstreamsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PodLogStreamList{})
	return err
}

// Patch applies the patch and returns the patched podLogStream.
func (c *FakePodLogStreams) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PodLogStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(podlogstreamsResource, name, pt, data, subresources...), &v1alpha1.PodLogStream{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodLogStream), err
}
