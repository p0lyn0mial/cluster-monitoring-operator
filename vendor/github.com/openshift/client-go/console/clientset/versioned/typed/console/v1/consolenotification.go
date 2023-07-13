// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/openshift/api/console/v1"
	consolev1 "github.com/openshift/client-go/console/applyconfigurations/console/v1"
	scheme "github.com/openshift/client-go/console/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConsoleNotificationsGetter has a method to return a ConsoleNotificationInterface.
// A group's client should implement this interface.
type ConsoleNotificationsGetter interface {
	ConsoleNotifications() ConsoleNotificationInterface
}

// ConsoleNotificationInterface has methods to work with ConsoleNotification resources.
type ConsoleNotificationInterface interface {
	Create(ctx context.Context, consoleNotification *v1.ConsoleNotification, opts metav1.CreateOptions) (*v1.ConsoleNotification, error)
	Update(ctx context.Context, consoleNotification *v1.ConsoleNotification, opts metav1.UpdateOptions) (*v1.ConsoleNotification, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ConsoleNotification, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ConsoleNotificationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConsoleNotification, err error)
	Apply(ctx context.Context, consoleNotification *consolev1.ConsoleNotificationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ConsoleNotification, err error)
	ConsoleNotificationExpansion
}

// consoleNotifications implements ConsoleNotificationInterface
type consoleNotifications struct {
	client rest.Interface
}

// newConsoleNotifications returns a ConsoleNotifications
func newConsoleNotifications(c *ConsoleV1Client) *consoleNotifications {
	return &consoleNotifications{
		client: c.RESTClient(),
	}
}

// Get takes name of the consoleNotification, and returns the corresponding consoleNotification object, and an error if there is any.
func (c *consoleNotifications) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ConsoleNotification, err error) {
	result = &v1.ConsoleNotification{}
	err = c.client.Get().
		Resource("consolenotifications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConsoleNotifications that match those selectors.
func (c *consoleNotifications) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ConsoleNotificationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ConsoleNotificationList{}
	err = c.client.Get().
		Resource("consolenotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested consoleNotifications.
func (c *consoleNotifications) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("consolenotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a consoleNotification and creates it.  Returns the server's representation of the consoleNotification, and an error, if there is any.
func (c *consoleNotifications) Create(ctx context.Context, consoleNotification *v1.ConsoleNotification, opts metav1.CreateOptions) (result *v1.ConsoleNotification, err error) {
	result = &v1.ConsoleNotification{}
	err = c.client.Post().
		Resource("consolenotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(consoleNotification).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a consoleNotification and updates it. Returns the server's representation of the consoleNotification, and an error, if there is any.
func (c *consoleNotifications) Update(ctx context.Context, consoleNotification *v1.ConsoleNotification, opts metav1.UpdateOptions) (result *v1.ConsoleNotification, err error) {
	result = &v1.ConsoleNotification{}
	err = c.client.Put().
		Resource("consolenotifications").
		Name(consoleNotification.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(consoleNotification).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the consoleNotification and deletes it. Returns an error if one occurs.
func (c *consoleNotifications) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("consolenotifications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *consoleNotifications) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("consolenotifications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched consoleNotification.
func (c *consoleNotifications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConsoleNotification, err error) {
	result = &v1.ConsoleNotification{}
	err = c.client.Patch(pt).
		Resource("consolenotifications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied consoleNotification.
func (c *consoleNotifications) Apply(ctx context.Context, consoleNotification *consolev1.ConsoleNotificationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ConsoleNotification, err error) {
	if consoleNotification == nil {
		return nil, fmt.Errorf("consoleNotification provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(consoleNotification)
	if err != nil {
		return nil, err
	}
	name := consoleNotification.Name
	if name == nil {
		return nil, fmt.Errorf("consoleNotification.Name must be provided to Apply")
	}
	result = &v1.ConsoleNotification{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("consolenotifications").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}