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
	"context"
	"fmt"
)

func fibs(ctx context.Context, n int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		a, b := 1, 1
		for i := 0; i < n; i++ {
			select {
			case ch <- a:
				a, b = b, a+b
			case <-ctx.Done():
				fmt.Println("cancelled")
			}
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := fibs(ctx, 5)
	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	cancel()
}
