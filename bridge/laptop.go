package bridge

type computer interface {
	printer
	setPrinter(printer)
}

type mac struct {
	name string
	printer
}

func (m *mac) print() string {
	return m.name + ": " + m.printer.print()
}

// setPrinter sets the printer for the mac
func (m *mac) setPrinter(p printer) {
	m.printer = p
}

// windows
type windows struct {
	name string
	printer
}

// setPrinter sets the printer for the windows
func (w *windows) setPrinter(p printer) {
	w.printer = p
}

func (w *windows) print() string {
	return w.name + ": " + w.printer.print()
}
