```go
type T int

func (t T) M() { print(t) }

type S struct{ *T }

var t = new(T)
var s = S{T: t}

func main() {
	f := t.M
	g := s.M
	h := reflect.ValueOf(s).MethodByName("M").Interface().(func())
	*t = 5
	f()
	g()
	h()
}
```