//Punto Numero 1 "Crear una lista"
//Punto Numero 2 "Diferencia entre la mayor y la menor cantidad"
//Punto numero 3 "Media y mediana"
//Punto numero 4 "la cantidad de semanas consecutivas en las que la presión promedio semanal supera o está por debajo del promedio anual de presiones semanales."
//Punto numero 6 "Temperaturas"

package main

import (
	"fmt"
)

func main() {

	listaTemperatura := []float64{200.783333, 196.111, 197.583333, 197.66667, 198.738889, 200.883333, 200.405556, 219.288889, 219.227778, 213.255556, 216.011111, 220.538889, 216.927778, 219.588889, 214.444, 216.111, 217.778, 215.494444, 215.538889, 199.966667, 201.538889, 201.677778, 195.461111, 199.194444, 193.077778, 199.38889, 193.933333, 191.05, 200.883333, 191.688889, 197.427778, 195.877778, 198.5, 197.383333, 199.811111, 225.605556, 229.338889, 232.377778, 222.566667222, 222.94444, 225.305556, 224.788889, 229.494444, 227.172222, 231.961111, 222.66667, 225.066667, 191.566667, 201.994444, 196.155556, 192.244444, 196.394444}

	fmt.Println(listaTemperatura)
	fmt.Print(len(listaTemperatura))

}
