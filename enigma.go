package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// extract real values
	va := ***a
	vb := *b
	vc := *******c
	vd := ****d

	// perform rotation
	// a → c
	*******c = va

	// c → d
	****d = vc

	// d → b
	*b = vd

	// b → a
	***a = vb
}
