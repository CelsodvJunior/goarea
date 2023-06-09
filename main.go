package goarea

import "math"

// Pi é uma proporção numérica definida pela relação entre
// O perímetro de um circunferência e seu diâmetro
const Pi = 3.1416

// Circ calcula a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
