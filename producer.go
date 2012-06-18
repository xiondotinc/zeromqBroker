/*
  Copyright 2010 Xion Inc
                 xion.inc@gmail.com
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

import "fmt"
import zmq "github.com/alecthomas/gozmq"

const (
	PULLSOCKET = "5000"
)
func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.PUSH)
	socket.Connect("tcp://127.0.0.1:" + PULLSOCKET)

	for {
		var msg string;
		_, error := fmt.Scanln(&msg)
		if error != nil {
			println("Error: Reading standard input.")
			return
		}
		socket.Send([]byte(msg), 0)
		println("Sending -> ", msg)
	}
}
