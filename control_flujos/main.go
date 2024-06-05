package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	t := time.Now()
	hora := t.Hour()
	fmt.Println(t)
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())

	fmt.Println("=====> Forma larga if else")
	//Version larga
	if hora < 12 {
		fmt.Println("!Mañana")
	} else if hora < 17 {
		fmt.Println("!Tarde")
	} else {
		fmt.Println("!Noche")
	}

	fmt.Println("=====> Forma corta if else")
	//Simplificado: forma de go
	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("!Mañana")
	} else if t.Hour() < 17 {
		fmt.Println("!Tarde")
	} else {
		fmt.Println("!Noche")
	}

	fmt.Println("=====> switch forma antigua")

	os := runtime.GOOS
	switch os {
	case "Windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}

	fmt.Println("=====> switch forma nueva")

	switch osss := runtime.GOOS; osss {
	case "Windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}

	fmt.Println("=====> switch forma nueva con tiempo")

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("!Mañana")
	case t.Hour() < 17:
		fmt.Println("!Tarde")
	default:
		fmt.Println("!Noche")
	}

	fmt.Println("=====> FOR infinito")
	//Se comenta
	/*
		for {
			fmt.Println("heyyy")
		}
	*/

	fmt.Println("=====> FOR condicion")
	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("=====> FOR bucle")
	for x := 1; x <= 10; x++ {
		fmt.Println(x)
	}

	fmt.Println("=====> FOR bucle Break")
	for x := 1; x <= 10; x++ {
		fmt.Println(x)
		if x == 5 {
			break
		}
	}

	fmt.Println("=====> FOR bucle continue")
	for x := 1; x <= 10; x++ {
		if x == 5 {
			continue
		}
		fmt.Println(x)
	}

}
