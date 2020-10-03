/*
Copyright 2020 The Crossplane Authors.

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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"

// GetProviderConfigReference of this ProviderConfigUsage.
func (p *ProviderConfigUsage) GetProviderConfigReference() runtimev1alpha1.Reference {
	return p.ProviderConfigReference
}

// GetResourceReference of this ProviderConfigUsage.
func (p *ProviderConfigUsage) GetResourceReference() runtimev1alpha1.TypedReference {
	return p.ResourceReference
}

// SetProviderConfigReference of this ProviderConfigUsage.
func (p *ProviderConfigUsage) SetProviderConfigReference(r runtimev1alpha1.Reference) {
	p.ProviderConfigReference = r
}

// SetResourceReference of this ProviderConfigUsage.
func (p *ProviderConfigUsage) SetResourceReference(r runtimev1alpha1.TypedReference) {
	p.ResourceReference = r
}
