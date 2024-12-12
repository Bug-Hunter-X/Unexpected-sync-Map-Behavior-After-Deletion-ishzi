```go
func main() {
    m := make(map[string]string)
    m["key"] = "value"
    val, ok := m["key"]
    fmt.Println(val, ok)
    delete(m, "key")
    val, ok = m["key"]
    fmt.Println(val, ok) // This will print <nil> false consistently after deletion
}
```