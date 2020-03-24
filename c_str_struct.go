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
	"unsafe"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef long long GoInt; // A Go integer is 64 bit on my machine
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

void show(char *msg) {
	GoSlice *s = (GoSlice*)(msg);
	// Create space for terminating NULL
	char *buf = (char *)malloc(s->len + 1);
	// Copy over original data
	memcpy(buf, s->data, s->len);
	// Add terminating NULL
	buf[s->len] = 0;

	printf("data = %s\n", buf);
}
*/
import "C"

func main() {
	msg := "srsly?"
	C.show((*C.char)(unsafe.Pointer(&msg)))
}
