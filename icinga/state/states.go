package state

// State represents the icinga plugin states as documented in <http://demo1.dtnet.de/icinga/docs/en/pluginapi.html>
type State int

const (
	// Ok represents the state 0: "OK" / "UP"
	Ok = State(0)

	// Warning represents the state 1: "WARNING" / "UP or DOWN/UNREACHABLE"
	Warning = State(1)

	// Critical represents the state 2: "CRITICAL" / "DOWN/UNREACHABLE"
	Critical = State(2)

	// Unknown represents the state 3: "UNKNOWN" / "DOWN/UNREACHABLE"
	Unknown = State(3)
)

// ServiceState returns the description of the state.
func (s State) ServiceState() string {
	switch s {
	case Ok:
		return "Ok"
	case Warning:
		return "Warning"
	case Critical:
		return "Critical"
	case Unknown:
		return "UNKNOWN"
	default:
		return "UNKNOWN"
	}
}
