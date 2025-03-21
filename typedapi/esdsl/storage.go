// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/storagetype"
)

type _storage struct {
	v *types.Storage
}

func NewStorage(type_ storagetype.StorageType) *_storage {

	tmp := &_storage{v: types.NewStorage()}

	tmp.Type(type_)

	return tmp

}

// You can restrict the use of the mmapfs and the related hybridfs store type
// via the setting node.store.allow_mmap.
// This is a boolean setting indicating whether or not memory-mapping is
// allowed. The default is to allow it. This
// setting is useful, for example, if you are in an environment where you can
// not control the ability to create a lot
// of memory maps so you need disable the ability to use memory-mapping.
func (s *_storage) AllowMmap(allowmmap bool) *_storage {

	s.v.AllowMmap = &allowmmap

	return s
}

func (s *_storage) Type(type_ storagetype.StorageType) *_storage {

	s.v.Type = type_
	return s
}

func (s *_storage) StorageCaster() *types.Storage {
	return s.v
}
