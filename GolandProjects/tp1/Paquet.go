package main

import "math/rand"

type Paquet struct {
	cartes []Carte
	jokers bool
}

func NouveauPaquet(jokers bool) *Paquet {

	taille := 52
	if jokers {
		taille = 54
	}
	p := &Paquet{cartes: make([]Carte, 0, taille), jokers: jokers} // une seule allocation

	p.reinitialiser()

	return p
}

func (p *Paquet) melanger() {
	rand.Shuffle(len(p.cartes), func(i, j int) {
		p.cartes[i], p.cartes[j] = p.cartes[j], p.cartes[i]
	})
}

func (p *Paquet) reinitialiser() {
	for coul := Pique; coul <= Trefle; coul++ {
		for r := byte(1); r <= 13; r++ {
			c, _ := NouvelleCarte(coul, r)
			p.cartes = append(p.cartes, c)
		}
	}
	if p.jokers {
		j1, _ := NouvelleCarte(JokerCouleur, 0)
		j2, _ := NouvelleCarte(JokerNoir, 0)
		p.cartes = append(p.cartes, j1, j2)
	}

	rand.Shuffle(len(p.cartes), func(i, j int) {
		p.cartes[i], p.cartes[j] = p.cartes[j], p.cartes[i]
	})
}
