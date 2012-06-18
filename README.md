zeromqBroker
============

A message broker implementation using zeromq in go programming language. 
Producers read from standard input and send messages to a Broker. 
All the consumers receive messages from the Broker.

To run:
start Broker by : go run broker.go
start Consumers by : go run bonsumer.go
start Producers by: go run producer.go

Typing anything for standard input for producer will be sent to all the consumers.
