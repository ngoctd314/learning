# Go 100 mistakes

[#1. Unintended variable shadowing]("")

[#2. Unnecessary nested code]("")

[#3. Misusing unit functions]("")

[#4. Overusing getters and setters]("")

[#5. Interface pollution]("")

[#6. Interface on the producer side]("")

[#7. Returning interfaces]("")

[#8. Any says nothing]("")

[#9. Being confused about when to use generice]("")

[#10. Not being aware of the possible problems with type embedding]("")

[#36. Not understanding the concept of a rune](./5.%20Strings.md#36-not-understanding-the-concept-of-a-rune)

- A charset is a set of characters(a,b,c...), whereas an encoding describes how to translate a charset into binary(97, 98, 99 -> base2).
- In Go, a string references an immutable slice of arbitrary bytes.
- Go source code is encoded using UTF-8. Hence, all string literals are UTF-8 strings. But because a string can contain arbitrary bytes, if it's obtained from somewhere else (not from the source code). It isn't guaranteed to be based on the UTF-8 encoding.
- A rune corresponds to the concept of a Unicode code point, meaning an item represented by a single value.
- Using UTF-8, a Unicode code point can be encoded into 1 to 4 bytes.
- Using len on a string in Go returns the number of bytes, not the number of runes.

[#37. Inaccurate string iteration]("./5.%20Strings.md#37-inaccurate-string-iteration")

- If we want to iterate over a string's runes, we can use the range loop on the string directly. But we have to recall that the index corresponds not to the rune index but rather to the starting index of the byte sequence of the rune.
- Rune can be composed of multiple bytes, if we want to access the rune itself, we should use the value variable of range, not the index in the string. Meanwhile, if we are interested in getting the ith rune of a string, we should convert the string into a slice of runes in most cases.

[#40. Useless string conversions]("./5.%20Strings.md#40-useless-string-conversions")

[#41. Substrings and memory leaks]("./5.%20Strings.md#41-substrings-and-memory-leaks")

- Understanding that a rune corresponds to the concept of a Unicode code point and that it can be composed of multiple bytes should be part of the Go developer's core knowledge to work accurately with strings.
- Iterating on a string with the range operator iterates on the runes with the index corresponding to the starting index of the rune's byte sequence. To access a specific rune index (such as the third rune), convert the string into a []rune.
- Strings.TrimRight/strings. TrimLeft removes all the trailing/leading runes contained in a given set, whereas strings. Trimsuffix/strings TrimPrefix returns a string without a provided suffix/prefix.
- Concatenating a list of strings should be done with strings.Builder to prevent allocating a new string during each iteration.
- Remembering that the bytes package offers the same operations as the strings package can help avoid extra byte/string conversions.
- Using copies instead of subtrings can prevent memory leaks, as the string returned by a substring operation will be backed by the same byte array.

[#42. Not knowing which type of receiver to use]("./6.%20Functions%20and%20methods.md#42-not-knowning-which-type-of-receiver-to-use")

- The decision whether to use a value or a pointer receiver should be based on factors such as the type, whether it has to be mutated, whether it contains a field that can't be copied, and how large the object it. When in doubt, use a pointer receiver.

[#43. Never using named result parameters]("./6.%20Functions%20and%20methods.md#43-never-using-named-result-parameters")

- Using named result parameters can be an efficient way to improve the readability of a function/method, especially if multiple result parameters have the same type. In some cases, this approach can also be convenient  because named result parameters are initialized to their zero value. But be cautions about potential side effects. 

[#44. Unintended side effects with named result parameters]("./6.%20Functions%20and%20methods.md#44-unintended-side-effects-with-named-result-parameters")

[#45. Returning a nil receiver]("./6.%20Functions%20and%20methods.md#45-returning-a-nil-receiver")

- A nil receiver is allowed, and an interface converted from a nil pointer isn't a nil interface. When we have to return an interface, we should return not a nil pointer but a nil value directly. Generally, having a nil pointer isn't a desirable state and means a probable bug.
- A nil pointer is different a nil interface.
- When returning an interface, be cautions about returning not a nil pointer but an explicit nil value. Otherwise, unintended consequences may result because the caller will receive a non-nil value.

[#46. Using a filename as a function input]("./6.%20Functions%20and%20methods.md#46-using-file-name-as-a-function-input")

- Designing functions to receive io.Reader types instead of filenames improves the reusability of a function and makes testing easier.

[#47. Ignoring how defer arguments and receivers are evaluated]("./6.%20Functions%20and%20methods.md#47-ignoring-how-defer-arguments-and-receivers-are-evaluated")

- When we call defer on a function or method, the call's arguments are evaluated immediately. If we want to mutate the arguments provided to defer afterward, we can use pointers or closures. For a method, the receiver is also evaluated immediately; hence, the behavior depends on whether the receiver is a value or a pointer.
- Passing a pointer to a defer function and wrapping a call inside a closure are two possible solutions to overcome the immediately evaluation of arguments and receivers. 

[#48. Panicking]("./7.%20Error%20management.md#48-panicking")

- mysql.Register is called in init(), which limits error handling. For all these reasons, the designer made the function panic in case of an error.
- Another use case in which to panic is when our application requires a dependency but fails to initialize it.   
- Using panic is an option to deal with errors in Go. However, it should only be used sparingly in unrecoverable conditions: for example, to signal a programmer error or when you fail to load a mandatory dependency.

[#49. Ignoring when to wrap an error]("./7.%20Error%20management.md#49-ignore-when-to-wrap-an-error")

- When handling an error, we can decide to wrap it. Wrapping is about adding context to an error and/or marking an error as specific type. If we need to mark an error, we should create a custom error type. However, if we just want to add extra context, we should use fmt.Errorf() with the %w directive as it doesn't require creating a new error type. Yet, error wrapping creates potential coupling as it makes the source error available for the caller. If we want to prevent it, we shouldn't use error wrapping but error transformation, for example, use fmt.Errorf with the %v directive.

[#50. Checking an error type inaccurately]("./7.%20Error%20management.md#50-checking-an-error-type-inaccurately")

- If we rely on Go 1.13 error wrapping, we must use errors.As to check whether an error is a specific type. This way, regardless of whether the error is returned directly by the function we call or wrapped our main error and see if one of the errors is a specific type.

[#51. Checking an error value inaccurately]("./7.%20Error%20management.md#51-checking-an-error-value-inaccurately")

- If we use error wrapping in our application with the %w directive and fmt.Errorf, checking an error against a specific value should be done using errors.Is instead of ==. Thus, event if the sentinel error is wrapped, errors.Is can recursively unwrap it and compare each error in the chain against the provided value.

[#52. Handling an error twice]("./7.%20Error%20management.md#52-handling-an-error-twice")

- Handling an error should be done only once. 
- Logging an error is handling an error. Hence, we should either log or return an error.

[#53. Not handling an error]("./7.%20Error%20management.md#53-not-handling-an-error")

- Ignoring an error in Go should be the exception. In many cases, we may still favor logging them, even at a low log level. But if we are sure that an error can and should be ignored, we must do so explicit by assigning it to the bank identifier. This way, a future reader will understanding that we ignored the error intentionally.

[#54. Not handling defer errors]("./7.%20Error%20management.md#54-not-handling-defer-errors")

- Errors should always be handled. In the case of errors returned by defer calls, the very least we should do is ignore them explicitly. If this isn't enough, we can handle the error directly by logging it or propagating it up to the caller.


[#55. Mixing up concurrency and parallelism]

Even after years of concurrent programming, developers may not clearly understanding the differences between concurrency and parallelism.

In this coffee shop, one waiter is in charge of accepting orders and preparing them using a single coffee machine. Customers give their orders and then wait for their coffee.

```go
coffeMachine := func(order string) any{
    return nil
} 
waiter := func(order string) any {
    return coffeMachine(order)
}

for _, customer := range customers {
    waiter(customer.order)    
}
```

If the waiter is having a hard time serving all the customers and the coffee shop wants to speed up the overall process, one idea might be to have a second waiter and a second coffee machine. A customer in the queue would wait for a waiter to be available.

In this new process, every part of the system is independent. The coffee shop should serve consumers twice as fast. This is a parallel implementation of a coffee shop.

If we want to scale, we can keep duplicating waiters and coffee machines over and over. However, this isn't the onl possible coffee shop design. Another approach might be to split the work done by the waiters and have one in charge of accepting orders and another one who grinds the coffee beans, which are then brewed in a single machine. Also, instead of blocking the customer queue until a customer is served, we could introduce another queue for customers waiting for their orders. (two queues: order queue, waiting queue).

With this new design, we don't make things parallel. But the overall struct is affected: we split a given role into two roles, and we introduce another queue. Unlike parallelism, which is about doing the same thing multiple times at once, concurrency is about structure.

Assuming one thread represents the waiter accepting orders and another represents the coffee machine, we have introduced yet another thread to grind the coffee beans. Each thread is independent but has to coordinate with others. Here, the waiter thread accepting orders has to communicate which coffee beans to grind. Meanwhile, the coffee-grinding threds must communicate with the coffee machine thread.

What if we want to increase throughput by serving more customers per hour? Because grinding beans takes longer than accepting orders, a possible change could be to hire another coffee-grinding waiter. 

[#82. Not categorizing tests]("")

Categorizing tests is a best practice for a successful testing strategy. In this section, we've seen three ways to categorize tests:

- Using build tags at the test file level.
- Using environment variables to make a specific test.
- Based on the test pace using short mode.

We can also combine approaches: for example, using build tags or environment variables to classify a test (for example, as a unit or integration test) and short mode if our project contains long-running unit tests. 

[#83. Not enabling the -race flag]("")

[#84. Not using test execution modes]("")

[#85. Not using table-driven tests]("")

[#86. Sleeping in unit tests]("")

[#91. Not understanding CPU caches]("")


