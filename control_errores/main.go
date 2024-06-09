package main

import (
	"errors"
	"fmt"
	"strconv"
)

//Manejo de errores

func dividev1(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func dividev2(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("No se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func main() {
	str := "123"
	num, err := strconv.Atoi(str)

	//error puede devolver 2 cosas: nil / mensaje de error
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("num :: ", num)

	//Usar error personalizado
	resultado1, err1 := dividev1(10, 0)
	if err1 != nil {
		fmt.Println("Error1: ", err1)
		//	return
	}
	fmt.Println("resultado1 :: ", resultado1)

	resultado2, err2 := dividev2(10, 0)
	if err2 != nil {
		fmt.Println("Error2: ", err2)
		//	return
	}
	fmt.Println("resultado2 :: ", resultado2)
}

//Error, tipo de dato que implementa un tipo interface

//Se pueden crear errores personalizados
