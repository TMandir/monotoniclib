# Monotonic Protocol

`monotoniclib` is the foundation for a Go library exploring monotonic changes
to entity state. Its current primitive checks that a proposed current value is
not less than the previous value.

This is a distinct protocol project. It does not depend on or define the rules
of Stability Calculus.

## Status and scope

The library currently provides a state value type and a comparison function.
It is an early foundation, not a complete protocol implementation. In
particular, the library does not yet validate entities, authorize transitions,
persist state, or define the full protocol rules.

`CheckUpgradeable` compares the two numeric values only. It accepts equal
values and does not impose a range on values; entity acceptance and other
protocol validation are outside its current scope.

## Installation

The module path is currently `monotoniclib`. Once published at its canonical
repository URL, use that URL as the import path in client projects.

## Usage

```go
package main

import (
	"fmt"

	monotonic "monotoniclib"
)

func main() {
	state := monotonic.EntityStabilityStateValue{
		Previous: 0.50,
		Current:  0.70,
	}

	fmt.Println("Is upgradeable:", monotonic.CheckUpgradeable(state))
}
```

Output:

```text
Is upgradeable: true
```

## API

- `EntityStabilityStateValue` contains exported `Previous` and `Current`
  values so callers can construct and inspect state comparisons.
- `CheckUpgradeable` returns whether `Current >= Previous`.

## Tests

Run the test suite from the module directory:

```sh
go test ./...
```
---
##More
Further more see: https://github.com/letbecool/Sunya
---

## License

All rights are reserved by Trustless Mandir Invention and Innovation Center.
See [LICENSE](./LICENSE).

