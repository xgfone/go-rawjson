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
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkBytes_MarshalJSON_Parallel(b *testing.B) {
	bs := []byte(` {"a" : 123 , "b" : " a b c "} `)

	b.ReportAllocs()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Bytes(bs).MarshalJSON()
		}
	})
}

func BenchmarkBytes_MarshalJSON_For(b *testing.B) {
	bs := []byte(` {"a" : 123 , "b" : " a b c "} `)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Bytes(bs).MarshalJSON()
	}
}

func BenchmarkBytes_WriterTo_Parallel(b *testing.B) {
	bs := []byte(` {"a" : 123 , "b" : " a b c "} `)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := getbuffer()
		Bytes(bs).WriteTo(buf)
		putbuffer(buf)
	}
}

func BenchmarkBytes_WriterTo_For(b *testing.B) {
	bs := []byte(` {"a" : 123 , "b" : " a b c "} `)
	buf := bytes.NewBuffer(nil)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		Bytes(bs).WriteTo(buf)
	}
}

func ExampleBytes() {
	/// Empty
	fmt.Println("\nEmpty")

	bs, _ := Bytes(nil).MarshalJSON()
	fmt.Printf("1: [%s]\n", string(bs))

	bs, _ = Bytes("").MarshalJSON()
	fmt.Printf("2: [%s]\n", string(bs))

	buf := bytes.NewBuffer(nil)
	Bytes(nil).WriteTo(buf)
	fmt.Printf("3: [%s]\n", buf.String())

	buf.Reset()
	Bytes(nil).WriteTo(buf)
	fmt.Printf("4: [%s]\n", buf.String())

	/// No Compact
	fmt.Println("\nNo Compact")
	Compact = false

	bs, _ = Bytes(` {"a" : 123 , "b" : " a b c "} `).MarshalJSON()
	fmt.Printf("5: [%s]\n", string(bs))

	/// Compact
	fmt.Println("\nCompact")
	Compact = true

	bs, _ = Bytes(` {"a" : 123 , "b" : " a b c "} `).MarshalJSON()
	fmt.Printf("6: [%s]\n", string(bs))

	bs, _ = Bytes(` 123 `).MarshalJSON()
	fmt.Printf("7: [%s]\n", string(bs))

	bs, _ = Bytes(` " 123 " `).MarshalJSON()
	fmt.Printf("8: [%s]\n", string(bs))

	// Output:
	//
	// Empty
	// 1: [""]
	// 2: [""]
	// 3: [""]
	// 4: [""]
	//
	// No Compact
	// 5: [ {"a" : 123 , "b" : " a b c "} ]
	//
	// Compact
	// 6: [{"a":123,"b":" a b c "}]
	// 7: [123]
	// 8: [" 123 "]
}
