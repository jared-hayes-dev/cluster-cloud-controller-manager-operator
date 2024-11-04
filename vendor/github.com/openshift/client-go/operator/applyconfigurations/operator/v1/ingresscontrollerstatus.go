// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// IngressControllerStatusApplyConfiguration represents a declarative configuration of the IngressControllerStatus type for use
// with apply.
type IngressControllerStatusApplyConfiguration struct {
	AvailableReplicas          *int32                                        `json:"availableReplicas,omitempty"`
	Selector                   *string                                       `json:"selector,omitempty"`
	Domain                     *string                                       `json:"domain,omitempty"`
	EndpointPublishingStrategy *EndpointPublishingStrategyApplyConfiguration `json:"endpointPublishingStrategy,omitempty"`
	Conditions                 []OperatorConditionApplyConfiguration         `json:"conditions,omitempty"`
	TLSProfile                 *configv1.TLSProfileSpec                      `json:"tlsProfile,omitempty"`
	ObservedGeneration         *int64                                        `json:"observedGeneration,omitempty"`
	NamespaceSelector          *metav1.LabelSelectorApplyConfiguration       `json:"namespaceSelector,omitempty"`
	RouteSelector              *metav1.LabelSelectorApplyConfiguration       `json:"routeSelector,omitempty"`
}

// IngressControllerStatusApplyConfiguration constructs a declarative configuration of the IngressControllerStatus type for use with
// apply.
func IngressControllerStatus() *IngressControllerStatusApplyConfiguration {
	return &IngressControllerStatusApplyConfiguration{}
}

// WithAvailableReplicas sets the AvailableReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AvailableReplicas field is set to the value of the last call.
func (b *IngressControllerStatusApplyConfiguration) WithAvailableReplicas(value int32) *IngressControllerStatusApplyConfiguration {
	b.AvailableReplicas = &value
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *IngressControllerStatusApplyConfiguration) WithSelector(value string) *IngressControllerStatusApplyConfiguration {
	b.Selector = &value
	return b
}

// WithDomain sets the Domain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Domain field is set to the value of the last call.
func (b *IngressControllerStatusApplyConfiguration) WithDomain(value string) *IngressControllerStatusApplyConfiguration {
	b.Domain = &value
	return b
}

// WithEndpointPublishingStrategy sets the EndpointPublishingStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndpointPublishingStrategy field is set to the value of the last call.
func (b *IngressControllerStatusApplyConfiguration) WithEndpointPublishingStrategy(value *EndpointPublishingStrategyApplyConfiguration) *IngressControllerStatusApplyConfiguration {
	b.EndpointPublishingStrategy = value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *IngressControllerStatusApplyConfiguration) WithConditions(values ...*OperatorConditionApplyConfiguration) *IngressControllerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithTLSProfile sets the TLSProfile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSProfile field is set to the value of the last call.
func (b *IngressControllerStatusApplyConfiguration) WithTLSProfile(value configv1.TLSProfileSpec) *IngressControllerStatusApplyConfiguration {
	b.TLSProfile = &value
	return b
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *IngressControllerStatusApplyConfiguration) WithObservedGeneration(value int64) *IngressControllerStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithNamespaceSelector sets the NamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NamespaceSelector field is set to the value of the last call.
func (b *IngressControllerStatusApplyConfiguration) WithNamespaceSelector(value *metav1.LabelSelectorApplyConfiguration) *IngressControllerStatusApplyConfiguration {
	b.NamespaceSelector = value
	return b
}

// WithRouteSelector sets the RouteSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RouteSelector field is set to the value of the last call.
func (b *IngressControllerStatusApplyConfiguration) WithRouteSelector(value *metav1.LabelSelectorApplyConfiguration) *IngressControllerStatusApplyConfiguration {
	b.RouteSelector = value
	return b
}
