package printer

// Printer represents a generic printer
type JsonPrintable interface {
	PrintJson(printable interface{}, options PrintOptions) error
}
type PlainPrintable interface {
	PrintPlain(printable interface{}, options PrintOptions) error
}
type YamlPrintable interface {
	PrintYaml(printable interface{}, options PrintOptions) error
}
type Printable interface {
	JsonPrintable
	PlainPrintable
	YamlPrintable
}
