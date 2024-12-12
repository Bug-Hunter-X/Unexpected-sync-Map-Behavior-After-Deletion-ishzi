```go
func main() {
    var m sync.Map
    m.Store("key", "value")
    val, ok := m.Load("key")
    fmt.Println(val, ok)
    m.Delete("key")
    val, ok = m.Load("key")
    fmt.Println(val, ok) // this will print <nil> false, even if the value was loaded previously
}
```