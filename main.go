package goarea

import "math"

// Pi é uma proporção númerica definida pela relação entre
// o perímetro de uma circuferencia e seu diametro
const Pi = 3.1416

// Circ é responsável por calcular a area da circuferencia.
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsavel por calcular a area de retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// não é uma função visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
