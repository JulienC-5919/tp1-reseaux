package cartes

import "fmt"

type Couleur byte

const (
	Pique Couleur = iota
	Coeur
	Carreau
	Trefle
	JokerCouleur
	JokerNoir
)

var symbolesCouleur = [...]string{"S", "H", "D", "C"}

var nomsRang = [...]string{"", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func (c Couleur) estValide() bool {
	return c <= Trefle
}

// Carte encode la couleur (bits 4-6) et le rang (bits 0-3) dans un octet.
type Carte byte

func NouvelleCarte(couleur Couleur, rang byte) (Carte, error) {
	switch {
	case couleur == JokerCouleur || couleur == JokerNoir:
		rang = 0
	case !couleur.estValide():
		return 0, fmt.Errorf("couleur invalide: %d", couleur)
	case rang < 1 || rang > 13:
		return 0, fmt.Errorf("rang invalide: %d", rang)
	}
	return Carte(byte(couleur)<<4 | rang), nil
}

func (c Carte) Rang() byte       { return byte(c) & 0x0F }
func (c Carte) Couleur() Couleur { return Couleur(byte(c) >> 4 & 0x07) }

func (c Carte) ToString() (string, error) {
	switch c.Couleur() {
	case JokerCouleur:
		return "JOKER", nil
	case JokerNoir:
		return "joker", nil
	}
	if !c.Couleur().estValide() {
		return "", fmt.Errorf("couleur invalide: %d", c.Couleur())
	}
	if c.Rang() < 1 || c.Rang() > 13 {
		return "", fmt.Errorf("rang invalide: %d", c.Rang())
	}
	return tableCartes[c], nil
}

func (c Carte) String() string {
	s, err := c.ToString()
	if err != nil {
		return "?"
	}
	return s
}

// tableCartes contient les chaînes partagées (poids mouche), indexées par la valeur de la carte.
var tableCartes = func() (t [128]string) {
	for coul := Pique; coul <= Trefle; coul++ {
		for r := byte(1); r <= 13; r++ {
			t[byte(coul)<<4|r] = nomsRang[r] + symbolesCouleur[coul]
		}
	}
	return
}()

// CarteDepuisString est l'inverse de String.
func CarteDepuisString(s string) (Carte, error) {
	switch s {
	case "JOKER":
		return NouvelleCarte(JokerCouleur, 0)
	case "joker":
		return NouvelleCarte(JokerNoir, 0)
	}
	for c, nom := range tableCartes {
		if nom == s && nom != "" {
			return Carte(c), nil
		}
	}
	return 0, fmt.Errorf("carte invalide: %q", s)
}
