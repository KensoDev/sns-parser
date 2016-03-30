# Amazon SNS Lambda Parser

[![Circle CI](https://circleci.com/gh/KensoDev/sns-parser.svg?style=svg)](https://circleci.com/gh/KensoDev/sns-parser)

![image](http://assets.avi.io/SimpleIcon_SNS.png)

Amazon sends a lot of events to SNS. You can "subscribe" to those events using email/https/sms and of course now Lambda functions.

sns-parser is meant to be paired with a lambda function.

The lambda function sends the JSON event to the Go application, it will parse it and then you can check your message.

## Code example

`main.js`

```javascript
var child_process = require('child_process');

exports.handler = function(event, context) {
  var proc = child_process.spawn('./notifier', [ JSON.stringify(event) ], { stdio: 'inherit' });

  proc.on('close', function(code) {
    if(code !== 0) {
      return context.done(new Error("Process exited with non-zero status code"));
    }

    context.done(null);
  });
}
```

`notifier.go`

```
package main

import (
	"fmt"
	"github.com/kensodev/sns-parser"
)

func main() {
	m := os.Args[1]
	parser := snsparser.NewSNSParser([]byte(m))

	failed, message := parser.IncludesMessage("Failed to deploy application")

	if failed {
		// DO SOMETHING With the message
		// it will return a SNS object (check models.go for more details)

	} else {
		fmt.Printf("Everything is OK, nothing to report in this message")
	}
}
```

In this sample, I am checking if the message includes `Failed to deploy application`. if it does, I can do whatever I want with it.

## Use cases

1. Send events to slack only if they match a specific condition
2. Send HTTP post if a certain event matched

etc...

## Production use

This code is deployed and working in production, consumed by a lambda function and sending messages to Slack using a webhook.