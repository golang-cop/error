<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/go-composites.png" alt="go-composites/error" width="720"></p>

# error

[![ci](https://github.com/go-composites/error/actions/workflows/ci.yml/badge.svg)](https://github.com/go-composites/error/actions/workflows/ci.yml)

The error composite of [go-composites](https://github.com/go-composites). An
`Error` is the value carried by a
[`Result`](https://github.com/go-composites/result): it follows the
Null-Object grammar, so an "absence of error" is itself a real object
(`NullError`) rather than a bare Go `nil`. `Result.HasError()` is simply
`!Error.IsNull()`, which is why a fresh `Result` — whose error is a
`NullError` — reports no error.

## Install

```sh
go get github.com/go-composites/error
```

## API

`Error.Interface` embeds the reflective `Object.Interface` (`Kind`, `RespondTo`,
`Methods`) and adds:

| method | returns | notes |
| --- | --- | --- |
| `Message()` | `string` | the error text |
| `IsNull()` | `bool` | `true` only for `NullError`; `false` for a real error |
| `Kind()` | `string` | the interface kind, e.g. `Error.Interface` (from `Object`) |
| `RespondTo(name)` | `bool` | whether the value exposes a method named `name` |
| `Methods()` | `[]string` | the value's method names, via reflection |

`New(message string) Interface` builds a concrete error with the given message
(its `IsNull()` is `false`).

### Subpackages

- **`error/src/object`** (`Object`) — the reflective base shared by every error
  type. Provides `Kind`, `RespondTo` and `Methods`, both as package functions
  over any `interface{}` and as the methods of the embedded `Object.Interface`.
- **`error/src/null`** (`NullError`) — the Null-Object error. `New()` yields an
  error whose `Message()` is `""` and whose `IsNull()` is `true`; it is the
  default error of a fresh `Result`.
- **`error/src/method_not_implemented`** (`MethodNotImplementedError`) — a
  sentinel error. `New(name string)` formats `"The method <name> is not
  implemented"`; its `IsNull()` is `false`.

## Usage

```go
package main

import (
	"fmt"

	Error "github.com/go-composites/error/src"
	NullError "github.com/go-composites/error/src/null"
	NotImplemented "github.com/go-composites/error/src/method_not_implemented"
)

func main() {
	e := Error.New("key not found")
	fmt.Println(e.Message())  // key not found
	fmt.Println(e.IsNull())   // false

	none := NullError.New()
	fmt.Println(none.IsNull()) // true  -> a Result holding this HasError() == false

	nim := NotImplemented.New("Frobnicate")
	fmt.Println(nim.Message()) // The method Frobnicate is not implemented
}
```

## License

BSD-3-Clause © the go-composites/error authors.
