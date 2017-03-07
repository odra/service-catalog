// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package servicecatalog

import (
	api "k8s.io/kubernetes/pkg/api"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_Binding, InType: reflect.TypeOf(&Binding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_BindingCondition, InType: reflect.TypeOf(&BindingCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_BindingList, InType: reflect.TypeOf(&BindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_BindingSpec, InType: reflect.TypeOf(&BindingSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_BindingStatus, InType: reflect.TypeOf(&BindingStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_Broker, InType: reflect.TypeOf(&Broker{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_BrokerCondition, InType: reflect.TypeOf(&BrokerCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_BrokerList, InType: reflect.TypeOf(&BrokerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_BrokerSpec, InType: reflect.TypeOf(&BrokerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_BrokerStatus, InType: reflect.TypeOf(&BrokerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_Instance, InType: reflect.TypeOf(&Instance{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_InstanceCondition, InType: reflect.TypeOf(&InstanceCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_InstanceList, InType: reflect.TypeOf(&InstanceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_InstanceSpec, InType: reflect.TypeOf(&InstanceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_InstanceStatus, InType: reflect.TypeOf(&InstanceStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_ServiceClass, InType: reflect.TypeOf(&ServiceClass{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_ServiceClassList, InType: reflect.TypeOf(&ServiceClassList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_servicecatalog_ServicePlan, InType: reflect.TypeOf(&ServicePlan{})},
	)
}

func DeepCopy_servicecatalog_Binding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Binding)
		out := out.(*Binding)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*api.ObjectMeta)
		}
		if err := DeepCopy_servicecatalog_BindingSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_servicecatalog_BindingStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_servicecatalog_BindingCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingCondition)
		out := out.(*BindingCondition)
		*out = *in
		return nil
	}
}

func DeepCopy_servicecatalog_BindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingList)
		out := out.(*BindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Binding, len(*in))
			for i := range *in {
				if err := DeepCopy_servicecatalog_Binding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_servicecatalog_BindingSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingSpec)
		out := out.(*BindingSpec)
		*out = *in
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		} else {
			out.Parameters = nil
		}
		return nil
	}
}

func DeepCopy_servicecatalog_BindingStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingStatus)
		out := out.(*BindingStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]BindingCondition, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Conditions = nil
		}
		return nil
	}
}

func DeepCopy_servicecatalog_Broker(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Broker)
		out := out.(*Broker)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*api.ObjectMeta)
		}
		if err := DeepCopy_servicecatalog_BrokerSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_servicecatalog_BrokerStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_servicecatalog_BrokerCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerCondition)
		out := out.(*BrokerCondition)
		*out = *in
		return nil
	}
}

func DeepCopy_servicecatalog_BrokerList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerList)
		out := out.(*BrokerList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Broker, len(*in))
			for i := range *in {
				if err := DeepCopy_servicecatalog_Broker(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_servicecatalog_BrokerSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerSpec)
		out := out.(*BrokerSpec)
		*out = *in
		if in.AuthSecret != nil {
			in, out := &in.AuthSecret, &out.AuthSecret
			*out = new(v1.ObjectReference)
			**out = **in
		} else {
			out.AuthSecret = nil
		}
		return nil
	}
}

func DeepCopy_servicecatalog_BrokerStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerStatus)
		out := out.(*BrokerStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]BrokerCondition, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Conditions = nil
		}
		return nil
	}
}

func DeepCopy_servicecatalog_Instance(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Instance)
		out := out.(*Instance)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*api.ObjectMeta)
		}
		if err := DeepCopy_servicecatalog_InstanceSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_servicecatalog_InstanceStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_servicecatalog_InstanceCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceCondition)
		out := out.(*InstanceCondition)
		*out = *in
		return nil
	}
}

func DeepCopy_servicecatalog_InstanceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceList)
		out := out.(*InstanceList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Instance, len(*in))
			for i := range *in {
				if err := DeepCopy_servicecatalog_Instance(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_servicecatalog_InstanceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceSpec)
		out := out.(*InstanceSpec)
		*out = *in
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		} else {
			out.Parameters = nil
		}
		if in.OSBDashboardURL != nil {
			in, out := &in.OSBDashboardURL, &out.OSBDashboardURL
			*out = new(string)
			**out = **in
		} else {
			out.OSBDashboardURL = nil
		}
		if in.OSBLastOperation != nil {
			in, out := &in.OSBLastOperation, &out.OSBLastOperation
			*out = new(string)
			**out = **in
		} else {
			out.OSBLastOperation = nil
		}
		return nil
	}
}

func DeepCopy_servicecatalog_InstanceStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceStatus)
		out := out.(*InstanceStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]InstanceCondition, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Conditions = nil
		}
		return nil
	}
}

func DeepCopy_servicecatalog_ServiceClass(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceClass)
		out := out.(*ServiceClass)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*api.ObjectMeta)
		}
		if in.Plans != nil {
			in, out := &in.Plans, &out.Plans
			*out = make([]ServicePlan, len(*in))
			for i := range *in {
				if err := DeepCopy_servicecatalog_ServicePlan(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Plans = nil
		}
		if in.OSBTags != nil {
			in, out := &in.OSBTags, &out.OSBTags
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.OSBTags = nil
		}
		if in.OSBRequires != nil {
			in, out := &in.OSBRequires, &out.OSBRequires
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.OSBRequires = nil
		}
		if in.OSBMaxDBPerNode != nil {
			in, out := &in.OSBMaxDBPerNode, &out.OSBMaxDBPerNode
			*out = new(string)
			**out = **in
		} else {
			out.OSBMaxDBPerNode = nil
		}
		if in.OSBDashboardOAuth2ClientID != nil {
			in, out := &in.OSBDashboardOAuth2ClientID, &out.OSBDashboardOAuth2ClientID
			*out = new(string)
			**out = **in
		} else {
			out.OSBDashboardOAuth2ClientID = nil
		}
		if in.OSBDashboardSecret != nil {
			in, out := &in.OSBDashboardSecret, &out.OSBDashboardSecret
			*out = new(string)
			**out = **in
		} else {
			out.OSBDashboardSecret = nil
		}
		if in.OSBDashboardRedirectURI != nil {
			in, out := &in.OSBDashboardRedirectURI, &out.OSBDashboardRedirectURI
			*out = new(string)
			**out = **in
		} else {
			out.OSBDashboardRedirectURI = nil
		}
		if in.Description != nil {
			in, out := &in.Description, &out.Description
			*out = new(string)
			**out = **in
		} else {
			out.Description = nil
		}
		if in.DisplayName != nil {
			in, out := &in.DisplayName, &out.DisplayName
			*out = new(string)
			**out = **in
		} else {
			out.DisplayName = nil
		}
		if in.ImageURL != nil {
			in, out := &in.ImageURL, &out.ImageURL
			*out = new(string)
			**out = **in
		} else {
			out.ImageURL = nil
		}
		if in.LongDescription != nil {
			in, out := &in.LongDescription, &out.LongDescription
			*out = new(string)
			**out = **in
		} else {
			out.LongDescription = nil
		}
		if in.ProviderDisplayName != nil {
			in, out := &in.ProviderDisplayName, &out.ProviderDisplayName
			*out = new(string)
			**out = **in
		} else {
			out.ProviderDisplayName = nil
		}
		if in.DocumentationURL != nil {
			in, out := &in.DocumentationURL, &out.DocumentationURL
			*out = new(string)
			**out = **in
		} else {
			out.DocumentationURL = nil
		}
		if in.SupportURL != nil {
			in, out := &in.SupportURL, &out.SupportURL
			*out = new(string)
			**out = **in
		} else {
			out.SupportURL = nil
		}
		return nil
	}
}

func DeepCopy_servicecatalog_ServiceClassList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceClassList)
		out := out.(*ServiceClassList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ServiceClass, len(*in))
			for i := range *in {
				if err := DeepCopy_servicecatalog_ServiceClass(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_servicecatalog_ServicePlan(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServicePlan)
		out := out.(*ServicePlan)
		*out = *in
		if in.Description != nil {
			in, out := &in.Description, &out.Description
			*out = new(string)
			**out = **in
		} else {
			out.Description = nil
		}
		if in.Bullets != nil {
			in, out := &in.Bullets, &out.Bullets
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Bullets = nil
		}
		if in.DisplayName != nil {
			in, out := &in.DisplayName, &out.DisplayName
			*out = new(string)
			**out = **in
		} else {
			out.DisplayName = nil
		}
		return nil
	}
}
