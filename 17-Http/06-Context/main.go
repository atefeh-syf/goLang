package main

import (
	"context-examples/examples"
	//"context-examples/management"
)

/*
Context
Creation
Send value
Cancellation from request
Internal cancellation
timeout
Deadline

*/

func main() {
	examples.InternalCancellationExample()
	//management.Run()
}
