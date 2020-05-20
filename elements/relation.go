package elements

import "fmt"

type C4Relation struct {
	C4Printable
	From       C4NodeElement
	To         C4NodeElement
	Label      string
	Technology string
}

func (r *C4Relation) ToC4PlantUMLString() string {
	return fmt.Sprintf("Rel(%v, '%v', '%s', '%s')\n", r.From.Alias(), r.To.Alias(), r.Label, r.Technology)
}
