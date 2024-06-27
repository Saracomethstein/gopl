package lenghtconv

import "fmt"

type Foot float64
type Meter float64

func (f Foot) String() string  { return fmt.Sprintf("%gF", f) }
func (m Meter) String() string { return fmt.Sprintf("%gM", m) }
