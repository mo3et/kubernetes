/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// NodeResourceSlices returns a NodeResourceSliceInformer.
	NodeResourceSlices() NodeResourceSliceInformer
	// PodSchedulingContexts returns a PodSchedulingContextInformer.
	PodSchedulingContexts() PodSchedulingContextInformer
	// ResourceClaims returns a ResourceClaimInformer.
	ResourceClaims() ResourceClaimInformer
	// ResourceClaimParameters returns a ResourceClaimParametersInformer.
	ResourceClaimParameters() ResourceClaimParametersInformer
	// ResourceClaimTemplates returns a ResourceClaimTemplateInformer.
	ResourceClaimTemplates() ResourceClaimTemplateInformer
	// ResourceClasses returns a ResourceClassInformer.
	ResourceClasses() ResourceClassInformer
	// ResourceClassParameters returns a ResourceClassParametersInformer.
	ResourceClassParameters() ResourceClassParametersInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// NodeResourceSlices returns a NodeResourceSliceInformer.
func (v *version) NodeResourceSlices() NodeResourceSliceInformer {
	return &nodeResourceSliceInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PodSchedulingContexts returns a PodSchedulingContextInformer.
func (v *version) PodSchedulingContexts() PodSchedulingContextInformer {
	return &podSchedulingContextInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ResourceClaims returns a ResourceClaimInformer.
func (v *version) ResourceClaims() ResourceClaimInformer {
	return &resourceClaimInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ResourceClaimParameters returns a ResourceClaimParametersInformer.
func (v *version) ResourceClaimParameters() ResourceClaimParametersInformer {
	return &resourceClaimParametersInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ResourceClaimTemplates returns a ResourceClaimTemplateInformer.
func (v *version) ResourceClaimTemplates() ResourceClaimTemplateInformer {
	return &resourceClaimTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ResourceClasses returns a ResourceClassInformer.
func (v *version) ResourceClasses() ResourceClassInformer {
	return &resourceClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ResourceClassParameters returns a ResourceClassParametersInformer.
func (v *version) ResourceClassParameters() ResourceClassParametersInformer {
	return &resourceClassParametersInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
