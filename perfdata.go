package icinga_plugin

import (
	"fmt"
	"strings"

	"clu-pei-dae/icinga-plugin/units"
)

// PerformanceDataElement is one element of performance data that can be returned by a icinga plugin
type PerformanceDataElement struct {
	Label         string
	Value         string
	Unit          units.UnitOfMeasurement
	WarningValue  string
	CriticalValue string
	MinimumValue  string
	MaximumValue  string
}

// String converts PerformanceDataElement to a string that conforms the performance data specification
func (p PerformanceDataElement) String() string {
	perfString := fmt.Sprintf("'%s'=%s%s;%s;%s;%s;%s", p.Label, p.Value, p.Unit, p.WarningValue, p.CriticalValue, p.MinimumValue, p.MaximumValue)
	return perfString
}

// PerformanceData is a slice of PerformanceDataElements that can be easily converted to a performance data string
type PerformanceData []PerformanceDataElement

// String converts PerformanceData to a string that conforms the performance data specification
func (p PerformanceData) String() string {
	elementStrings := make([]string, len(p))

	for i, element := range p {
		elementStrings[i] = element.String()
	}

	return strings.Join(elementStrings, " ")
}

// Add adds an PerformanceDataElement to PerformanceData
func (p *PerformanceData) Add(element PerformanceDataElement) {
	*p = append(*p, element)
}
