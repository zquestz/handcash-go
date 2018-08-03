# handcash-go

Go library for handcash's API @ http://handcash.io/api-docs

## Installation

```
go get -u github.com/zquestz/handcash-go
```

## CLI Usage

```
handcash-go someuser
```

This returns:

```
{
    "receivingAddress": "mxszqDyaNGFcmTkPjJ2BGRpSTChdVWaNPZ",
    "publicKey": "03d193439a2f06ed1121be5b4e61381386ffee5ec5bec33daf17e33ccb34622753"
}
```

## Library Usage

```
import (
    "fmt"

    "github.com/zquestz/handcash-go/api"
)

func main() {
    handle := "someuser"
    resp, err := api.Receive(handle)
    if err != nil {
        fmt.Println("request failed")
        return
    }

    fmt.Println(resp)
}
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/zquestz/handcash-go.

## License

This library is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
