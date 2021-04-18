package lengthconv

import "fmt"

//Feet represents feet
type Feet float64

//Meter represents Meter
type Meter float64

//CFtM convert feet to meter
func CFtM(f Feet) Meter { return Meter(f / 3.281) }

//CMtF convert meter to feet
func CMtF(m Meter) Feet { return Feet(m * 3.281) }

func (m Meter) String() string { return fmt.Sprintf("%g m", m) }
func (f Feet) String() string  { return fmt.Sprintf("%g f", f) }
