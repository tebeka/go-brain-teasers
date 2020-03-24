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
	a := []int{1, 2, 3}
	a = Append(a, 4)
	fmt.Println(a)
	b := a[:3]
	b = append(b, 5)
	fmt.Println(b)
}

func Append(items []int, i int) []int {
	if len(items) == cap(items) { // No more space in underlying array
		// Go has a better growth heuristic than adding 1 every append
		newItems := make([]int, len(items)+1)
		copy(newItems, items)
		items = newItems
	} else {
		items = items[:len(items)+1]
	}

	items[len(items)-1] = i
	return items
}
