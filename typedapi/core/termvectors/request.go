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

package termvectors

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package termvectors
//
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/_global/termvectors/TermVectorsRequest.ts#L33-L187
type Request struct {

	// Doc An artificial document (a document not present in the index) for which you
	// want to retrieve term vectors.
	Doc json.RawMessage `json:"doc,omitempty"`
	// Filter Filter terms based on their tf-idf scores.
	// This could be useful in order find out a good characteristic vector of a
	// document.
	// This feature works in a similar manner to the second phase of the More Like
	// This Query.
	Filter *types.TermVectorsFilter `json:"filter,omitempty"`
	// PerFieldAnalyzer Override the default per-field analyzer.
	// This is useful in order to generate term vectors in any fashion, especially
	// when using artificial documents.
	// When providing an analyzer for a field that already stores term vectors, the
	// term vectors will be regenerated.
	PerFieldAnalyzer map[string]string `json:"per_field_analyzer,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		PerFieldAnalyzer: make(map[string]string, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Termvectors request: %w", err)
	}

	return &req, nil
}
