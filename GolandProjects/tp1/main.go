package main

import "tp1/cartes"

func main() {
	p := cartes.NouveauMPaquet(2, true)
	//p.Melanger()
	for i := 0; i < 54; i++ {
		c, _ := p.Piocher()
		s, _ := c.ToString()
		println(s)
	}

	println(p.Restantes())
}
