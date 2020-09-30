package states

// State represents the icinga plugin states as documented in <http://demo1.dtnet.de/icinga/docs/en/pluginapi.html>
type State int

const (
	// Ok represents the states 0: "OK" / "UP"
	Ok = State(0)

	// Warning represents the states 1: "WARNING" / "UP or DOWN/UNREACHABLE"
	Warning = State(1)

	// Critical represents the states 2: "CRITICAL" / "DOWN/UNREACHABLE"
	Critical = State(2)

	// Unknown represents the states 3: "UNKNOWN" / "DOWN/UNREACHABLE"
	Unknown = State(3)
)

// ServiceState returns the description of the states.
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
