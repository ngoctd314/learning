# Wire User Guide

## Basics

Wire has two core concepts: providers and injectors

### Defining Providers

The primary mechanism in Wire is the **provider:** a function that can produce a value. These functions are ordinary Go code

```go
package foobarbaz

type Foo struct {
    X int
}

// ProvideFoo return Foo
func ProvideFoo() Foo {
    return Foo{X: 42}
}
```

Provider functions must be exported in order to be used from other packages, just like ordinary functions.

Providers can specify dependencies with parameters:

```go
package foobarbaz

type Bar struct {
    X int
}

// ProvideBar returns a Bar: a negative Foo
func ProvideBar(foo Foo) Bar {
    return Bar{X: -foo.X}
}
```

Provider can also return errors:

```go
package foobarbaz

import (
    "context"
    "errors"
)

type Baz struct {
    X int
}

func ProvideBaz(ctx context.Context, bar Baz) (Baz, error) {
    if baz.X == 0 {
        return Baz{}, errors.New("cannot provide baz when bar is zero")
    }
    return Baz{X: bar.X}, nil
}
```

Provider can be grouped into **provider sets.** This is useful if several providers will frequently be used together. To add these providers to a new set called SuperSet, use the wire.NewSet function:

```go
package foobarbaz

var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)
```

You can also add other provider sets into a provider set 
```go
package foobarbaz

var MegaSet = wire.NewSet(SuperSet, pkg.OtherSet)
```

### Injectors

An application wires up these providers with an injector: a function that calls providers in dependency order. With Wire, you write the injector's signature, then Wire generates the function's body.

An injector is declared by writing a function declaration whose body is call to wire.Build. The return values don't matter as long as they are of the correct type. The values themselves will be ignored in the generated code.

```go
package main

func initializeBar(ctx context.Context) (foobarbaz.Baz, error) {
    wire.Build(foobarbaz.MegaSet)
    return foobarbaz.Baz{}, nil
}
```

Like providers, injectors can be parameterized on inputs (which then get sent to providers) and can return errors. Arguments to wire.Build are the same as wire.NewSet: the form a provider set. This is the provider set that gets used during code generation for that injector.

Any non-injector declarations found in a file with injectors will be copied into the genarted fild.

Wire will produce an implementation of the injector in a file called wire_gen.go that looks something like this:

```go
package main

func initializeBaz(ctx context.Context) (foobarbaz.Baz, error) {
    foo := foobarbaz.ProvideFoo()
    bar := foobarbaz.ProvideBar(foo)
    baz, err := foobarbaz.ProvideBaz(ctx, bar)
    if err != nil {
        return foobarbaz.Baz{}, err
    }
    return baz, nil
}
```

## Advanced Features

### Binding Interfaces

Frequently, dependency injection is used to bind a concrete implementation for an interface. Wire matches inputs to outputs via type identity, so the inclination might be to create a provider that returns an interface type. However, this would not be idiomatic, since the Go best practice is to return concrete types. Instead, you can declare an interface binding in a provider set:

```go
type Fooer interface {
    Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
    return string(*b)
}

func provideMyFooer() *MyFooer {
    b := new(MyFooer)
    *b = "Hello World!"
    return b
}

```