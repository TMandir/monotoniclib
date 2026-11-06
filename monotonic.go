// Package monotonic provides primitives for evaluating monotonic entity state
// changes.
package monotonic

// EntityStabilityStateValue contains the previous and proposed current values
// for an entity's state.
type EntityStabilityStateValue struct {
	Previous float64
	Current  float64
}

// CheckUpgradeable reports whether the current value is greater than or equal
// to the previous value.
//
// It compares values only; it does not validate the entity or authorize a
// state transition.
func CheckUpgradeable(state EntityStabilityStateValue) bool {
	return state.Current >= state.Previous
}
