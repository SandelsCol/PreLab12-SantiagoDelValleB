//Punto Numero 1 "Crear una lista"
//Punto Numero 2 "Diferencia entre la mayor y la menor cantidad"
//Punto numero 3 "Media y mediana"

package main

import (
	"fmt"
	"sort"
)

func main() {

	lista := []float64{110.06, 107.89, 108.45, 108.49, 109.03, 110.11, 109.87, 119.38, 119.35, 116.34, 117.73, 120.01, 118.19, 119.53, 117.09, 118.03, 118.65, 117.47, 117.49, 109.65, 110.44, 110.51, 107.38, 109.26, 106.18, 109.36, 106.61, 105.16, 110.11, 105.48, 108.37, 107.59, 108.91, 108.35, 109.57, 122.56, 124.44, 125.97, 121.03, 121.22, 122.41, 122.15, 124.52, 123.35, 125.76, 121.08, 122.29, 105.42, 110.67, 107.73, 105.76, 107.85}

	x := 5930
	y := len(lista)
	Media := x / y

	fmt.Println("La media es", Media, "Kpa aproximadamente")

	sort.Float64s(lista)

	Promedio1 := lista[:28]
	Promedio2 := lista[28:]
	Longitud1 := len(Promedio1)
	Longitud2 := len(Promedio2)

	fmt.Println("Los siguientes valores corresponden a las semanas malas por debajo del promedio", Promedio1, "Siendo un total de", Longitud1, "Semanas")
	fmt.Println("Los siguientes valores corresponden a las semanas buenas por encima del promedio", Promedio2, "Siendo un total de", Longitud2, "Semanales")
}
