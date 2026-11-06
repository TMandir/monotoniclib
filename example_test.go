package monotonic_test

import (
	"fmt"

	monotonic "github.com/TMandir/monotoniclib"
)

func ExampleCheckUpgradeable() {
	state := monotonic.EntityStabilityStateValue{
		Previous: 0.50,
		Current:  0.70,
	}

	fmt.Println("Is upgradeable:", monotonic.CheckUpgradeable(state))
	// Output: Is upgradeable: true
}
