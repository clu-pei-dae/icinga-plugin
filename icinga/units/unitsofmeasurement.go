package units

// UnitOfMeasurement is a unit as specified in <http://demo1.dtnet.de/icinga/docs/en/perfdata.html#perfdata-format>
type UnitOfMeasurement string

const (
	NotSpecified = UnitOfMeasurement("")

	Seconds      = UnitOfMeasurement("s")
	MilliSeconds = UnitOfMeasurement("ms")
	MicroSeconds = UnitOfMeasurement("us")

	Percentage = UnitOfMeasurement("%")

	Bytes      = UnitOfMeasurement("B")
	KiloBytes  = UnitOfMeasurement("KB")
	MegaBytes  = UnitOfMeasurement("MB")
	GigaBytes  = UnitOfMeasurement("GB")
	TerraBytes = UnitOfMeasurement("TB")

	Continuous = UnitOfMeasurement("c")
)
