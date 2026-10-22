package cartes

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

// ErrPaquetVide est retournée quand on pioche dans un paquet vide.
var ErrPaquetVide = errors.New("le paquet est vide")

// ErrCarteIntrouvable est retournée quand la carte demandée n'est pas dans la main.
var ErrCarteIntrouvable = errors.New("carte introuvable dans la main")

type Paquet struct {
	cartes  []Carte
	paquets byte
	jokers  bool
	main    []Carte
	piles   map[string][]Carte
}

func NouveauPaquet(jokers bool) *Paquet {
	return NouveauMPaquet(1, jokers)
}

// NouveauMPaquet crée un paquet composé de plusieurs jeux de 52 cartes (plus 2 jokers si jokers est vrai).
func NouveauMPaquet(paquets byte, jokers bool) *Paquet {
	return &Paquet{
		cartes:  genererCartes(paquets, jokers),
		paquets: paquets,
		jokers:  jokers,
		main:    []Carte{},
		piles:   make(map[string][]Carte),
	}
}

// melanger mélange les cartes en place (Fisher-Yates).
func melanger(cartes []Carte) {
	rand.Shuffle(len(cartes), func(i, j int) {
		cartes[i], cartes[j] = cartes[j], cartes[i]
	})
}

// genererCartes génère l'ensemble des cartes d'un ou plusieurs paquets.
//
// Chaque paquet contient 52 cartes, plus 2 jokers si jokers est vrai.
// Le slice retourné n'est pas mélangé.
func genererCartes(paquets byte, jokers bool) []Carte {
	taille := 52
	if jokers {
		taille = 54
	}

	taille *= int(paquets)

	cartes := make([]Carte, 0, taille) // une seule allocation

	for s := byte(0); s < paquets; s++ {
		for coul := Pique; coul <= Trefle; coul++ {
			for r := byte(1); r <= 13; r++ {
				c, _ := NouvelleCarte(coul, r)
				cartes = append(cartes, c)
			}
		}

		if jokers {
			j1, _ := NouvelleCarte(JokerCouleur, 0)
			j2, _ := NouvelleCarte(JokerNoir, 0)
			cartes = append(cartes, j1, j2)
		}
	}

	return cartes
}

func (p *Paquet) Restantes() int {
	return len(p.cartes)
}

func (p *Paquet) Melanger() {
	melanger(p.cartes)
}

func (p *Paquet) Reinitialiser() {
	p.cartes = genererCartes(p.paquets, p.jokers)
	p.main = []Carte{}
	clear(p.piles)
}

func (p *Paquet) AjouterPile(nom string) {
	p.piles[nom] = []Carte{}
}

// AjouterDansPile retire la carte de la main et l'ajoute à la pile nommée.
// Si la pile n'existe pas ou si la carte n'est pas dans la main, rien n'est modifié.
func (p *Paquet) AjouterDansPile(nom string, c Carte) error {
	pile, ok := p.piles[nom]
	if !ok {
		return fmt.Errorf("la pile %s n'existe pas", nom)
	}
	i := slices.Index(p.main, c)
	if i < 0 {
		return fmt.Errorf("%w: %v", ErrCarteIntrouvable, c)
	}
	p.main = slices.Delete(p.main, i, i+1)
	p.piles[nom] = append(pile, c)
	return nil
}

// Piocher retire la carte du dessus du paquet, l'ajoute à la main et la retourne.
func (p *Paquet) Piocher() (Carte, error) {
	if len(p.cartes) == 0 {
		return 0, ErrPaquetVide
	}
	dernier := len(p.cartes) - 1
	c := p.cartes[dernier]
	p.cartes = p.cartes[:dernier]
	p.main = append(p.main, c)
	return c, nil
}
