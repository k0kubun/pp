# pp

Colored pretty printer for Go language

## Usage

```go
import "github.com/k0kubun/pp"

func main() {
  m := map[string]string{
    "foo": "bar",
  }
  pp.Print(m)
}
```

## Supported Types

- [x] Bool
- [x] Int
- [x] Int8
- [x] Int16
- [x] Int32
- [x] Int64
- [x] Uint
- [x] Uint8
- [x] Uint16
- [x] Uint32
- [x] Uint64
- [x] Uintptr
- [x] Float32
- [x] Float64
- [x] Complex64
- [x] Complex128
- [ ] Array
- [ ] Chan
- [ ] Func
- [ ] Interface
- [x] Map
- [ ] Ptr
- [ ] Slice
- [x] String
- [ ] Struct
- [ ] UnsafePointer

### TODO

support private field of struct

## License

MIT License
