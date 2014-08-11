//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.
package bleve

import (
	"github.com/couchbaselabs/bleve/document"
)

type Batch map[string]interface{}

func NewBatch() Batch {
	return make(Batch, 0)
}

func (b Batch) Index(id string, data interface{}) {
	b[id] = data
}

func (b Batch) Delete(id string) {
	b[id] = nil
}

type Index interface {
	Index(id string, data interface{}) error
	Delete(id string) error

	Batch(b Batch) error

	Document(id string) (*document.Document, error)
	DocCount() uint64

	Search(req *SearchRequest) (*SearchResult, error)

	Fields() ([]string, error)

	Dump()
	DumpDoc(id string) ([]interface{}, error)
	DumpFields()

	Close()
}

type Classifier interface {
	Type() string
}

// Open the index at the specified path, and create it if it does not exist.
// The provided mapping will be used for all Index/Search operations.
func Open(path string, mapping *IndexMapping) (Index, error) {
	return newIndex(path, mapping)
}
