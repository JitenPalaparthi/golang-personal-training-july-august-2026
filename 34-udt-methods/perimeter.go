package main

func (re Rect) Perimeter() float64 { // re is a receiver
	return 2 * float64(re.L+re.B)
}
