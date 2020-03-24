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
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	t1 := time.Now()
	data, err := json.Marshal(t1)
	if err != nil {
		log.Fatal(err)
	}

	var t2 time.Time
	if err := json.Unmarshal(data, &t2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(t1 == t2)
}
