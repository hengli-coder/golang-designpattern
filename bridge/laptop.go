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

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

type windows struct {
	name string
	printer
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

func (w *windows) print() string {
	return w.name + ": " + w.printer.print()
}
