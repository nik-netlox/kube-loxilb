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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	loxiurlv1 "github.com/loxilb-io/kube-loxilb/pkg/crds/loxiurl/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// LoxiURLLister helps list LoxiURLs.
// All objects returned here must be treated as read-only.
type LoxiURLLister interface {
	// List lists all LoxiURLs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*loxiurlv1.LoxiURL, err error)
	// Get retrieves the LoxiURL from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*loxiurlv1.LoxiURL, error)
	LoxiURLListerExpansion
}

// loxiURLLister implements the LoxiURLLister interface.
type loxiURLLister struct {
	listers.ResourceIndexer[*loxiurlv1.LoxiURL]
}

// NewLoxiURLLister returns a new LoxiURLLister.
func NewLoxiURLLister(indexer cache.Indexer) LoxiURLLister {
	return &loxiURLLister{listers.New[*loxiurlv1.LoxiURL](indexer, loxiurlv1.Resource("loxiurl"))}
}
