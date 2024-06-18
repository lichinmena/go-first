package main

import (
	"fmt"
	"net/http"
	"time"
)

//Concurrencia
//go runtimes
//Son independientes entre si
//Se ejecuta de forma simultanea en orden de forma determinista

/* Sin concurrencia
func main() {
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://google.com.mx",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://api.some.com",
	}

	for _, api := range apis {
		checkAPI(api)
	}

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomo %v segundos!\n", elapsed.Seconds())
}

func checkAPI(api string) {
	if _, err := http.Get(api); err != nil {
		fmt.Printf("ERROR: ¡%s esta caido\n", api)
	}

	fmt.Printf("SUCCESS: ¡%s esta en funcionamiento!\n", api)
}
*/

/*
// Con concurrencia normal
func main() {
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://google.com.mx",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://api.some.com",
	}

	for _, api := range apis {
		go checkAPI(api)
	}

	//Este Sleep, es solo para comprobar que si lo hace.
	//NO se percibe por que lo hace muy rapido
	time.Sleep(5 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomo %v segundos!\n", elapsed.Seconds())
}

func checkAPI(api string) {
	if _, err := http.Get(api); err != nil {
		fmt.Printf("ERROR: ¡%s esta caido\n", api)
	}

	fmt.Printf("SUCCESS: ¡%s esta en funcionamiento!\n", api)
}
*/

// Con concurrencia con canales
func main() {
	//EJEMplo de como crear un canal
	/*
		canal := make(chan int)
		canal <- 15
		valor := <- canal
	*/
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://google.com.mx",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://api.some.com",
	}

	ch := make(chan string)
	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomo %v segundos!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("ERROR: ¡%s esta caido\n", api)
	}

	ch <- fmt.Sprintf("SUCCESS: ¡%s esta en funcionamiento!\n", api)
}
