/*
  Copyright 2012 Xion Inc  
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

import zmq "github.com/alecthomas/gozmq"

const (
	PUBSOCKET = "6000"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.SUB)
	socket.Connect("tcp://127.0.0.1:" + PUBSOCKET)
	socket.SetSockOptString(zmq.SUBSCRIBE, "")
	for {
		msg, _ := socket.Recv(0)
		println("Received : ", string(msg))
	}
}