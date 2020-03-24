/*
Copyright 2020 Miki Tebeka <miki@353solutions.com>

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

package main

import (
	"fmt"
)

type FilePerm uint16 // 16 flags are enough

const (
	Read FilePerm = 1 << iota
	Write
	Execute
)

// String implements fmt.Stringer interface
func (p FilePerm) String() string {
	switch p {
	case Read:
		return "read"
	case Write:
		return "write"
	case Execute:
		return "execute"
	}

	return fmt.Sprintf("unknown FilePerm: %d", p) // don't use %s here :)
}

func main() {
	fmt.Println(Execute)        // execute
	fmt.Printf("%d\n", Execute) // 4
}
