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

func main() {
	// Integer
	printNum(10)    // 10 of type int
	printNum(010)   // 8 of type int
	printNum(0x10)  // 16 of type int
	printNum(0b10)  // 2 of type int
	printNum(1_000) // 1000 of type int <1>

	// Float
	printNum(3.14)   // 3.14 of type float64
	printNum(.2)     // 0.2 of type float64
	printNum(1e3)    // 1000 of type float64 <!--2-->
	printNum(0x1p-2) // 0.25 of type float64 <3>

	// Complex
	printNum(1i)     // (0+1i) of type complex128
	printNum(3 + 7i) // (3+7i) of type complex128
	printNum(1 + 0i) // (1+0i) of type complex128
}

func printNum(n interface{}) {
	fmt.Printf("%v of type %T\n", n, n)
}
