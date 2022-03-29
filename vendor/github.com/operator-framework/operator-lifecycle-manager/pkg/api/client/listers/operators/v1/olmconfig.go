/*
Copyright Red Hat, Inc.

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
	v1 "github.com/operator-framework/api/pkg/operators/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OLMConfigLister helps list OLMConfigs.
// All objects returned here must be treated as read-only.
type OLMConfigLister interface {
	// List lists all OLMConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.OLMConfig, err error)
	// Get retrieves the OLMConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.OLMConfig, error)
	OLMConfigListerExpansion
}

// oLMConfigLister implements the OLMConfigLister interface.
type oLMConfigLister struct {
	indexer cache.Indexer
}

// NewOLMConfigLister returns a new OLMConfigLister.
func NewOLMConfigLister(indexer cache.Indexer) OLMConfigLister {
	return &oLMConfigLister{indexer: indexer}
}

// List lists all OLMConfigs in the indexer.
func (s *oLMConfigLister) List(selector labels.Selector) (ret []*v1.OLMConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.OLMConfig))
	})
	return ret, err
}

// Get retrieves the OLMConfig from the index for a given name.
func (s *oLMConfigLister) Get(name string) (*v1.OLMConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("olmconfig"), name)
	}
	return obj.(*v1.OLMConfig), nil
}