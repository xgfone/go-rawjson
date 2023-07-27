// Copyright 2023 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package rawjson provides the raw bytes json.
package rawjson

import (
	"bytes"
	"sync"
)

var (
	// The capacity size of the buffer.
	BufCapSize = 512

	// If true, compact the raw bytes json。
	Compact = true
)

var (
	emptyString = []byte(`""`)

	bufpool = sync.Pool{New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, BufCapSize))
	}}
)

func getbuffer() *bytes.Buffer  { return bufpool.Get().(*bytes.Buffer) }
func putbuffer(b *bytes.Buffer) { b.Reset(); bufpool.Put(b) }
