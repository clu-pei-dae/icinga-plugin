package icinga

import "fmt"

func CreateOutputData(message, perfdata string) string {
	return fmt.Sprintf("%s | %s", message, perfdata)
}
