// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"

	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// CiliumBGPPeerConfigsGetter has a method to return a CiliumBGPPeerConfigInterface.
// A group's client should implement this interface.
type CiliumBGPPeerConfigsGetter interface {
	CiliumBGPPeerConfigs() CiliumBGPPeerConfigInterface
}

// CiliumBGPPeerConfigInterface has methods to work with CiliumBGPPeerConfig resources.
type CiliumBGPPeerConfigInterface interface {
	Create(ctx context.Context, ciliumBGPPeerConfig *v2alpha1.CiliumBGPPeerConfig, opts v1.CreateOptions) (*v2alpha1.CiliumBGPPeerConfig, error)
	Update(ctx context.Context, ciliumBGPPeerConfig *v2alpha1.CiliumBGPPeerConfig, opts v1.UpdateOptions) (*v2alpha1.CiliumBGPPeerConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.CiliumBGPPeerConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.CiliumBGPPeerConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumBGPPeerConfig, err error)
	CiliumBGPPeerConfigExpansion
}

// ciliumBGPPeerConfigs implements CiliumBGPPeerConfigInterface
type ciliumBGPPeerConfigs struct {
	*gentype.ClientWithList[*v2alpha1.CiliumBGPPeerConfig, *v2alpha1.CiliumBGPPeerConfigList]
}

// newCiliumBGPPeerConfigs returns a CiliumBGPPeerConfigs
func newCiliumBGPPeerConfigs(c *CiliumV2alpha1Client) *ciliumBGPPeerConfigs {
	return &ciliumBGPPeerConfigs{
		gentype.NewClientWithList[*v2alpha1.CiliumBGPPeerConfig, *v2alpha1.CiliumBGPPeerConfigList](
			"ciliumbgppeerconfigs",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v2alpha1.CiliumBGPPeerConfig { return &v2alpha1.CiliumBGPPeerConfig{} },
			func() *v2alpha1.CiliumBGPPeerConfigList { return &v2alpha1.CiliumBGPPeerConfigList{} }),
	}
}
