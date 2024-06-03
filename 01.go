package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

// Constantes a nivel de paquete
// const Pi float32 = 3.1416
const Pi = 3.1416

const (
	x = 100
	y = 0b1010 // binario
	z = 0o12   // octal
	w = 0xFF   //hexadecimal
)

// Valor IOTA
const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main2() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	//Declaraciones de variables
	var firstName string

	var one, two, three string

	var age int

	firstName = "Luis"
	one = "Uno"
	two = "Dos"
	three = "Tres"
	age = 35

	fmt.Println(firstName, one, two, three, age)

	var (
		company string = "Chevrolet"
		model   string = "Sonic"
		year    int    = 2016
	)

	fmt.Println(company, model, year)

	var (
		person  = "Fiscal"
		address = "Merida"
	)

	fmt.Println(person, address)

	//Se pueden desclarar dentro y fuera de las funciones
	var color, size, long = "Blue", "Medium", 100
	fmt.Println(color, size, long)

	//Mas simple.- Se declaran dentro de las funciones
	num1, num2, isTrue := 1, 2, true
	fmt.Println(num1, num2, isTrue)

	num3 := 100
	num3 = 123
	fmt.Println(num3)

	//Constantes ===========>
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)

	//Tipo de datos basicos
	var m1 int16 = 12
	var m2 uint16 = 1
	fmt.Println(m1, m2)

	var m3 float32 = 100.12
	var m4 float64 = 120.33
	fmt.Println(m3, m4)

	//Paquete Mat para realizar operaciones mas complejas
	fmt.Println(math.MinInt64, math.MaxFloat64)

	var valOne bool = true
	var valTwo bool = false
	fmt.Println(valOne, valTwo)

	fullName := "Luis Mena \t(alias \"lichin\")"
	fmt.Println(fullName)

	//byte alias del tipo uint8, almacena numeros sin signo
	//0-255
	//Datos de caracteres ascii y byte
	var a byte = 'a'

	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0]) //IMprime el valor decimal de h

	//Running, es un alias de int32. unicode
	var r rune = '❤'
	fmt.Println(r)

	//Valores prederminados, por defecto
	var (
		defaultInt    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)
	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)

	//20.- Conversiones de tipos
	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(int32(integer16) + integer32)

	//String a tipo entero
	ss := "100"
	x1, _ := strconv.Atoi(ss) //Esta funcion devuelve dos valores, la conversion y el error.
	//el guion bajo representa que no vamos a utilizar el error, por el momento.
	fmt.Println("String a entero :: ", x1)

	//Entero a cadena
	n1 := 42
	ss = strconv.Itoa(n1)
	fmt.Println("Entero a String :: ", ss)

	//21.- Paquete fmt
	//Es uno de los mas utilzados
	fmt.Print("Otro Mensaje")

	fmt.Println()

	name1 := "Luis"
	age1 := 35
	fmt.Printf("Hola, me llamo %s y tengo %d años \n", name1, age1)

	gretting := fmt.Sprintf("Hola, me llamo %s y tengo %d años \n", name1, age1)
	fmt.Println(gretting)

	//Ingresar datos de usuario
	var name2 string
	var age2 int

	fmt.Print("Ingrese el valor de name2:")
	fmt.Scanln(&name2) //Referencia de la memoria

	fmt.Print("Ingrese el valor de age2:")
	fmt.Scanln(&age2) //Referencia de la memoria

	fmt.Printf("Hola, me llamo %s y tengo %d años \n", name2, age2)
	//%v -> cuando no sabes el tipo de dato a mostrar

	//Qu2 tipo de dato tiene una variable
	fmt.Printf("El tipo de dato de name2 es: %T\n", name2)
	fmt.Printf("El tipo de dato de age2 es: %T\n", age2)

}
