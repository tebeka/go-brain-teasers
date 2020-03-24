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

type Job struct {
	State string
	done  chan struct{}
}

func (j *Job) Wait() {
	<-j.done
}

func (j *Job) Done() {
	j.State = "done"
	close(j.done)
}

func main() {
	ch := make(chan Job)
	go func() {
		j := <-ch
		j.Done()
	}()

	job := Job{"ready", make(chan struct{})}
	ch <- job
	job.Wait()
	fmt.Println(job.State)
}
