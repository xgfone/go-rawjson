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

package rawjson

import (
	"encoding"
	"encoding/json"
	"io"
)

var (
	_ io.WriterTo            = Bytes(nil)
	_ json.Marshaler         = Bytes(nil)
	_ encoding.TextMarshaler = Bytes(nil)
)

// Bytes represent a raw byte json.
type Bytes []byte

// MarshalText implements the interface encoding.TextMarshaler,
// which is equal to bs.MarshalJSON().
func (bs Bytes) MarshalText() ([]byte, error) { return bs.marshal() }

// MarshalJSON implements the interface json.Marshaler.
//
// If Compact is true, compact the bytes.
// If the bytes is empty, return []byte(`""`) instead.
func (bs Bytes) MarshalJSON() (b []byte, err error) { return bs.marshal() }

func (bs Bytes) marshal() (b []byte, err error) {
	if len(bs) == 0 {
		return emptyString, nil
	}

	if !Compact {
		return bs, nil
	}

	buf := getbuffer()
	if err = json.Compact(buf, bs); err == nil {
		b = make([]byte, buf.Len())
		copy(b, buf.Bytes())
	}
	putbuffer(buf)

	return
}

// WriteTo implements the interface io.WriterTo.
func (bs Bytes) WriteTo(w io.Writer) (n int64, err error) {
	if len(bs) == 0 {
		m, err := w.Write(emptyString)
		if m != len(bs) {
			err = io.ErrShortWrite
		}
		return int64(m), err
	}

	if !Compact {
		m, err := w.Write(bs)
		if m != len(bs) {
			err = io.ErrShortWrite
		}
		return int64(m), err
	}

	buf := getbuffer()
	if err = json.Compact(buf, bs); err == nil {
		n, err = buf.WriteTo(w)
	}
	putbuffer(buf)

	return
}
