package example

//go:generate  accessor_engine Station

//go:generate go run ../generate_specail.go Gs
type Gs struct {
	Name string
}

//go:generate go run ../generate_specail.go Station
type Station struct {
	Exp  int
	Name string
}
