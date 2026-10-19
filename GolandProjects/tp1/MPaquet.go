package main

import "math/rand"

type MPaquet struct {
	cartes  []Carte
	paquets byte
	jokers  bool
}

func nouveauMPaquet(paquets byte, jokers bool) *MPaquet {
	taille := 52
	if jokers {
		taille = 54
	}

	taille *= int(paquets)

	p := &MPaquet{cartes: make([]Carte, 0, taille), paquets: paquets, jokers: jokers} // une seule allocation

	p.reinitialiser()

	return p
}

func (p *MPaquet) melanger() {
	rand.Shuffle(len(p.cartes), func(i, j int) {
		p.cartes[i], p.cartes[j] = p.cartes[j], p.cartes[i]
	})
}

func (p *MPaquet) reinitialiser() {
	base := NouveauPaquet(p.jokers).cartes // un paquet complet, réutilisé pour chaque copie

	p.cartes = p.cartes[:0] // conserve la capacité allouée
	for s := byte(0); s < p.paquets; s++ {
		p.cartes = append(p.cartes, base...)
	}

	p.melanger()
}
