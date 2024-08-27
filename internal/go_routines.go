package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

//Escribe un programa que inicie varias goroutines que envíen números a un canal. En la función principal, recibe y suma estos números.

func SumChannel() {
	var wg sync.WaitGroup
	ch := make(chan int)
	numbers := []int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		wg.Add(1)                // Incrementar el contador de WaitGroup
		go sendNum(num, ch, &wg) // Iniciar una goroutine
	}

	go func() {
		wg.Wait() // Esperar a que todas las goroutines terminen
		close(ch) // Cerrar el canal
	}()

	sum := 0
	for n := range ch {
		fmt.Println(n)
		sum += n // Sumar los números recibidos
	}

	fmt.Println("Sum:", sum)
}

func sendNum(num int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Indicar que la goroutine ha terminado
	c <- num        // Enviar el número al canal
}

// Crea una función que use goroutines para realizar tareas concurrentes, como abrir varias páginas web al mismo tiempo.
func OpenPages(pages []string) {
	var wg sync.WaitGroup

	for _, page := range pages {
		wg.Add(1) // Incrementar el contador de WaitGroup
		go fetchURL(&wg, page)
	}

	wg.Wait() // Esperar a que todas las goroutines terminen
	fmt.Println("All pages fetched")

}

func fetchURL(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}

	defer resp.Body.Close() // se cierre correctamente después de que se haya terminado de utilizar.
	// Esta es una práctica importante en el manejo de recursos en Go, ya que ayuda a prevenir fugas de recursos (resource leaks).

	body, err := ioutil.ReadAll(resp.Body) //Lee el cuerpo de la respuesta
	if err != nil {
		fmt.Printf("Error reading response from %s: %v\n", url, err)
		return
	}
	fmt.Printf("Fetched %d bytes from %s\n", len(body), url)
}

// Implementa un programa que utilice channels para coordinar la ejecución de múltiples goroutines. La idea es enviar una cadena de numeros y evaluar si son todos pares
func IsPairChannel(numeros []int) {
	ch := make(chan bool)

	var wg sync.WaitGroup

	for _, numero := range numeros {
		wg.Add(1)
		go func() {
			sendPairValue(numero, ch, &wg)
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	isPair := true
	for numero := range ch {
		if !numero {
			isPair = false
			break
		}
	}

	if isPair {
		fmt.Println("Todos los numeros de la cadena son pares")
	} else {
		fmt.Println("En la cadena hay numeros que no son pares")
	}
}

func sendPairValue(num int, ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done() // Indicar que la goroutine ha terminado
	isPair := num%2 == 0
	ch <- isPair
}
