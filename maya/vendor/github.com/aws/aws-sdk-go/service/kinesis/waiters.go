// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package kinesis

import (
	"github.com/aws/aws-sdk-go/private/waiter"
)

// WaitUntilStreamExists uses the Kinesis API operation
// DescribeStream to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *Kinesis) WaitUntilStreamExists(input *DescribeStreamInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeStream",
		Delay:       10,
		MaxAttempts: 18,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "StreamDescription.StreamStatus",
				Expected: "ACTIVE",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
