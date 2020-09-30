package icinga_plugin

import (
	"testing"

	"clu-pei-dae/icinga-plugin/units"
)

func TestPerformanceDataElement_ToString(t *testing.T) {
	type fields struct {
		Label         string
		Value         string
		Unit          units.UnitOfMeasurement
		WarningValue  string
		CriticalValue string
		MinimumValue  string
		MaximumValue  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "percentage", fields: fields{
			Label:         "CPU",
			Value:         "2.3",
			Unit:          units.Percentage,
			WarningValue:  "2.5",
			CriticalValue: "3",
			MinimumValue:  "1",
			MaximumValue:  "4",
		}, want: "'CPU'=2.3%;2.5;3;1;4"},
		{name: "miliseconds", fields: fields{
			Label:         "time",
			Value:         "5",
			Unit:          units.MilliSeconds,
			WarningValue:  "2.5",
			CriticalValue: "3",
			MinimumValue:  "1",
			MaximumValue:  "4",
		}, want: "'time'=5ms;2.5;3;1;4"},
		{name: "miliseconds without min/max", fields: fields{
			Label:         "time",
			Value:         "5",
			Unit:          units.MilliSeconds,
			WarningValue:  "2.5",
			CriticalValue: "3",
		}, want: "'time'=5ms;2.5;3;;"},
		{name: "just value", fields: fields{
			Label: "time",
			Value: "5",
		}, want: "'time'=5;;;;"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PerformanceDataElement{
				Label:         tt.fields.Label,
				Value:         tt.fields.Value,
				Unit:          tt.fields.Unit,
				WarningValue:  tt.fields.WarningValue,
				CriticalValue: tt.fields.CriticalValue,
				MinimumValue:  tt.fields.MinimumValue,
				MaximumValue:  tt.fields.MaximumValue,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerformanceData_Add(t *testing.T) {
	type args struct {
		element PerformanceDataElement
	}
	tests := []struct {
		name string
		p    PerformanceData
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPerformanceData_ToString(t *testing.T) {
	tests := []struct {
		name string
		p    PerformanceData
		want string
	}{
		{name: "single value", p: PerformanceData{
			PerformanceDataElement{
				Label: "time",
				Value: "5",
			}}, want: "'time'=5;;;;"},
		{name: "multiple values", p: PerformanceData{
			PerformanceDataElement{
				Label: "time",
				Value: "5",
			},
			PerformanceDataElement{
				Label:         "startup",
				Value:         "5",
				Unit:          units.MilliSeconds,
				WarningValue:  "2.5",
				CriticalValue: "3"},
		}, want: "'time'=5;;;; 'startup'=5ms;2.5;3;;"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerformanceData_Add1(t *testing.T) {
	type args struct {
		element PerformanceDataElement
	}
	tests := []struct {
		name string
		p    PerformanceData
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
