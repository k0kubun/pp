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

### TODO

- support private field of struct

## License

MIT License
