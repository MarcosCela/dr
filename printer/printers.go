package printer

import "fmt"
import "io"
import "github.com/hokaccha/go-prettyjson"

// Printer represents a generic printer
type Printer interface {
	Print(message interface{})
}
type SmartPrinter interface {
	Print(printable interface{})
}

// Plain prints text on a plaintext format
type Plain struct {
	writer io.Writer
}

// JSON prints text in formatted JSON
type JSON struct {
	writer io.Writer
}

// PlainPrinter returns an instance of the default printer
func PlainPrinter(writer io.Writer) Printer {
	return &Plain{writer}

}

// JSONPrinter returns an instance of a printer that formats output on a pretty JSON format
func JSONPrinter(writer io.Writer) Printer {
	return &JSON{writer}
}

// Print prints a message with the default formatter
func (printer Plain) Print(message interface{}) {
	switch x := message.(type) {
	case []string:
		for _, item := range x {
			_, _ = fmt.Fprintf(printer.writer, "%+v\n", item)
		}
	default:
		_, _ = fmt.Fprintf(printer.writer, "%+v\n", x)
	}
}

// Print prints a message with the JSON formatter
func (printer JSON) Print(message interface{}) {
	s, _ := prettyjson.Marshal(message)
	_, _ = fmt.Fprintf(printer.writer, "%v\n", string(s))
}
