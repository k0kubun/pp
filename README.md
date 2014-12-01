# pp [![wercker status](https://app.wercker.com/status/d68cd67ea8da91a05e4e61a79852dad9/s "wercker status")](https://app.wercker.com/project/bykey/d68cd67ea8da91a05e4e61a79852dad9)

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
- [x] Array
- [x] Chan
- [ ] Func
- [ ] Interface
- [x] Map
- [ ] Ptr
- [x] Slice
- [x] String
- [x] Struct
- [ ] UnsafePointer

### TODO

- support private field of struct
- better type parser

## License

MIT License
