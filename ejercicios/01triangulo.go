/*
Ejercicio: Calcule e imprima el área y el perímetro del triángulo
Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.

El programa debe:

Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.

Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.

Calcular el área del triángulo usando la fórmula base x altura / 2.

Calcular el perímetro del triángulo sumando los lados.

Imprimir el área y el perímetro del triángulo con dos decimales de precisión.

El programa debe usar variables para almacenar los lados del triángulo, la hipotenusa, el área y el perímetro. También debe usar constantes para representar el número de decimales de precisión que se desean en la salida.

Además, se deben utilizar funciones del paquete Math de Go para calcular la raíz cuadrada y cualquier otro cálculo matemático necesario.
*/
package main

import (
	"fmt"
	"math"
)

const TWO_DECIMALS string = "%.2f"
const PRECISION int16 = 2

func main() {
	fmt.Println("----> Calcular el area y perimetro de un triangulo rectangulo <----")

	var ladoLargo float64
	var ladoCorto float64

	fmt.Print("Ingresa la longitud de lado mas largo del triangulo: ")
	fmt.Scanln(&ladoLargo)

	fmt.Print("Ingresa la longitud de lado mas corto del triangulo: ")
	fmt.Scanln(&ladoCorto)

	fmt.Printf("Los datos capturados son: lado largo %v lado corto %v\n", ladoLargo, ladoCorto)

	var hipotenusa float64
	var area float64
	var perimetro float64

	//calculo de la hipotenusa
	//hipotenusa = math.Sqrt((ladoLargo * ladoLargo) + (ladoCorto * ladoCorto))
	hipotenusa = math.Sqrt(math.Pow(ladoLargo, 2) + math.Pow(ladoCorto, 2))
	area = (ladoLargo * ladoCorto) / 2
	perimetro = ladoLargo + ladoCorto + hipotenusa

	//hipotenusaFormateada := fmt.Sprintf(TWO_DECIMALS, hipotenusa)
	//areaFormateada := fmt.Sprintf(TWO_DECIMALS, area)
	//perimetroFormateado := fmt.Sprintf(TWO_DECIMALS, perimetro)

	//fmt.Println("La hipotenusa es: ", hipotenusaFormateada)
	//fmt.Println("El area es: ", areaFormateada)
	//fmt.Println("El perimetro es: ", perimetroFormateado)

	fmt.Printf("\nHipotenusa: %.*f \n", PRECISION, hipotenusa)
	fmt.Printf("\nArea: %.*f \n", PRECISION, area)
	fmt.Printf("\nPerimetro: %.*f \n", PRECISION, perimetro)
}
