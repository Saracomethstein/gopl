package lenghtconv

func FToM(f Foot) Meter { return Meter(f * 0.3048) }
func MToF(m Meter) Foot { return Foot(m * 3.28084) }
