//Punto Numero 1 "Crear una lista"
//Punto Numero 2 "Diferencia entre la mayor y la menor cantidad"
//Punto numero 3 "Media y mediana"
//Punto numero 4 "la cantidad de semanas consecutivas en las que la presión promedio semanal supera o está por debajo del promedio anual de presiones semanales."
//Punto numero 6 "Temperaturas"
//punto numero 6.1 "Desviaciones"
//Punto numero 6.2 "Clasificación"
//punto numero 6.3 "Promedio de desviaciones"

package main

import "fmt"

func main() {

	Desviaciongeneral := 13.08
	DesviacionB := 5.37
	DesviacionM := 3.32

	promedio := (DesviacionB + DesviacionM) / 2.0

	fmt.Println("El promedio de las desviaciones fue de", promedio, "F°")

	Diferencia := Desviaciongeneral - promedio

	fmt.Println("La diferencia entre el la desviación general y el promedio entre la buena y la mala es de", Diferencia)

}
