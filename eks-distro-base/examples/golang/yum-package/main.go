// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"

	libdeflate "github.com/4kills/go-libdeflate/v2"
)

func main() {
	// Compressor with default compression level. Errors if out of memory
	c, err := libdeflate.NewCompressor()
	if err != nil {
		log.Panic(err)
	}

	hello := []byte(`Hello World`)

	_, comp, err := c.Compress(hello, nil, libdeflate.ModeZlib)
	if err != nil {
		log.Panic(err)
	}

	c.Close()

	fmt.Println("Compressed hello world: %s", comp)
}
