# At least once Delivery

In case of failures that lead to message loss or take too long to recover from, messages are retransmitted to assure they are delivered at least once.

## Context

Sometimes, message duplicity can be coped with by the application using a Message-oriented Middleware. Therefore, for scenarios where message duplicates are uncritical, it shall will be ensured that messages are received.

## Solution

For each message retrieved by a receiver an ack is sent back to the message sender. In case this ack is not retrieved after a certain time frame, the messsage is resend.


