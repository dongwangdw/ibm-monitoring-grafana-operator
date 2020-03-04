// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Grafana) DeepCopyInto(out *Grafana) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Grafana.
func (in *Grafana) DeepCopy() *Grafana {
	if in == nil {
		return nil
	}
	out := new(Grafana)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Grafana) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfig) DeepCopyInto(out *GrafanaConfig) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = new(GrafanaConfigPath)
		**out = **in
	}
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = new(GrafanaConfigServer)
		**out = **in
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = new(GrafanaConfigUser)
		**out = **in
	}
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(GrafanaConfigLog)
		**out = **in
	}
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(GrafanaConfigAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(GrafanaConfigAuthProxy)
		(*in).DeepCopyInto(*out)
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(GrafanaConfigSecurity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfig.
func (in *GrafanaConfig) DeepCopy() *GrafanaConfig {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAuth) DeepCopyInto(out *GrafanaConfigAuth) {
	*out = *in
	if in.DisableLoginForm != nil {
		in, out := &in.DisableLoginForm, &out.DisableLoginForm
		*out = new(bool)
		**out = **in
	}
	if in.DisableSignoutMenu != nil {
		in, out := &in.DisableSignoutMenu, &out.DisableSignoutMenu
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAuth.
func (in *GrafanaConfigAuth) DeepCopy() *GrafanaConfigAuth {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigAuthProxy) DeepCopyInto(out *GrafanaConfigAuthProxy) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.AutoSignUp != nil {
		in, out := &in.AutoSignUp, &out.AutoSignUp
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigAuthProxy.
func (in *GrafanaConfigAuthProxy) DeepCopy() *GrafanaConfigAuthProxy {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigAuthProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigLog) DeepCopyInto(out *GrafanaConfigLog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigLog.
func (in *GrafanaConfigLog) DeepCopy() *GrafanaConfigLog {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigPath) DeepCopyInto(out *GrafanaConfigPath) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigPath.
func (in *GrafanaConfigPath) DeepCopy() *GrafanaConfigPath {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigSecurity) DeepCopyInto(out *GrafanaConfigSecurity) {
	*out = *in
	if in.DisableInitialAdminCreation != nil {
		in, out := &in.DisableInitialAdminCreation, &out.DisableInitialAdminCreation
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigSecurity.
func (in *GrafanaConfigSecurity) DeepCopy() *GrafanaConfigSecurity {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigServer) DeepCopyInto(out *GrafanaConfigServer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigServer.
func (in *GrafanaConfigServer) DeepCopy() *GrafanaConfigServer {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfigUser) DeepCopyInto(out *GrafanaConfigUser) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfigUser.
func (in *GrafanaConfigUser) DeepCopy() *GrafanaConfigUser {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfigUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDatasource) DeepCopyInto(out *GrafanaDatasource) {
	*out = *in
	if in.IsDefault != nil {
		in, out := &in.IsDefault, &out.IsDefault
		*out = new(bool)
		**out = **in
	}
	if in.Editable != nil {
		in, out := &in.Editable, &out.Editable
		*out = new(bool)
		**out = **in
	}
	in.JSONData.DeepCopyInto(&out.JSONData)
	out.SecureJSONData = in.SecureJSONData
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDatasource.
func (in *GrafanaDatasource) DeepCopy() *GrafanaDatasource {
	if in == nil {
		return nil
	}
	out := new(GrafanaDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaIngress) DeepCopyInto(out *GrafanaIngress) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaIngress.
func (in *GrafanaIngress) DeepCopy() *GrafanaIngress {
	if in == nil {
		return nil
	}
	out := new(GrafanaIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaList) DeepCopyInto(out *GrafanaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Grafana, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaList.
func (in *GrafanaList) DeepCopy() *GrafanaList {
	if in == nil {
		return nil
	}
	out := new(GrafanaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaResources) DeepCopyInto(out *GrafanaResources) {
	*out = *in
	if in.Grafana != nil {
		in, out := &in.Grafana, &out.Grafana
		*out = new(int)
		**out = **in
	}
	if in.Dashboard != nil {
		in, out := &in.Dashboard, &out.Dashboard
		*out = new(int)
		**out = **in
	}
	if in.Router != nil {
		in, out := &in.Router, &out.Router
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaResources.
func (in *GrafanaResources) DeepCopy() *GrafanaResources {
	if in == nil {
		return nil
	}
	out := new(GrafanaResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaService) DeepCopyInto(out *GrafanaService) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ServicePort, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaService.
func (in *GrafanaService) DeepCopy() *GrafanaService {
	if in == nil {
		return nil
	}
	out := new(GrafanaService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaSpec) DeepCopyInto(out *GrafanaSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(GrafanaConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Datasource != nil {
		in, out := &in.Datasource, &out.Datasource
		*out = new(GrafanaDatasource)
		(*in).DeepCopyInto(*out)
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(GrafanaService)
		(*in).DeepCopyInto(*out)
	}
	if in.MetaData != nil {
		in, out := &in.MetaData, &out.MetaData
		*out = new(MetaData)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(GrafanaIngress)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(GrafanaResources)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaSpec.
func (in *GrafanaSpec) DeepCopy() *GrafanaSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaStatus) DeepCopyInto(out *GrafanaStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaStatus.
func (in *GrafanaStatus) DeepCopy() *GrafanaStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaData) DeepCopyInto(out *MetaData) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaData.
func (in *MetaData) DeepCopy() *MetaData {
	if in == nil {
		return nil
	}
	out := new(MetaData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSAuth) DeepCopyInto(out *TLSAuth) {
	*out = *in
	if in.KeepCookies != nil {
		in, out := &in.KeepCookies, &out.KeepCookies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLSAuth != nil {
		in, out := &in.TLSAuth, &out.TLSAuth
		*out = new(bool)
		**out = **in
	}
	if in.TLSAuthWithCACert != nil {
		in, out := &in.TLSAuthWithCACert, &out.TLSAuthWithCACert
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSAuth.
func (in *TLSAuth) DeepCopy() *TLSAuth {
	if in == nil {
		return nil
	}
	out := new(TLSAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}