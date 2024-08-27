package main

import (
	"fmt"
	"intermedio/internal"
)

func main() {
	//p1 := internal.Persona{
	//	Nombre:   "Pavarotti",
	//	Apellido: "Luciano",
	//	Edad:     54,
	//}
	//
	//minor := 10
	//mayor := 5
	//
	//internal.IntercambiarValores(&minor, &mayor)
	//internal.IntercambiarNombre(&p1)

	// Crear una nueva instancia de SequenceGenerator
	//sequence := internal.SequenceGenerator()
	//
	//// Llama a sequence varias veces
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//
	//// Crear otra instancia de contador independiente
	//otherSequence := internal.SequenceGenerator()
	//
	//// Llamar al segundo contador
	//fmt.Println(otherSequence())
	//fmt.Println(otherSequence())
	//fmt.Println(otherSequence())
	//
	////fmt.Println("El valor de minor es:", minor)
	////fmt.Println("El valor de mayor es:", mayor)
	////fmt.Println("El nombre es:", p1.Nombre)
	////fmt.Println("El apellido es:", p1.Apellido)
	//
	//sequenceWithJump := internal.SequecenJumperGenerator(2, 6)
	//fmt.Println(sequenceWithJump())
	//fmt.Println(sequenceWithJump())
	//fmt.Println(sequenceWithJump())
	//fmt.Println(sequenceWithJump())
	//fmt.Println(sequenceWithJump())
	//
	//seq, restart := internal.SequecenJumperGeneratorWithRestart(10, 5)
	//
	//fmt.Println(seq())
	//fmt.Println(seq())
	//fmt.Println(seq())
	//fmt.Println(seq())
	//
	//restart()
	//
	//fmt.Println(seq())
	//fmt.Println(seq())
	//fmt.Println(seq())
	//
	//split := func(a, b int) int {
	//	return a / b
	//}
	//
	//add := func(a, b int) int {
	//	return a + b
	//}
	//
	//subtract := func(a, b int) int {
	//	return a - b
	//}
	//
	//resultSplit := split(30, 5)
	//resultAdd := add(3, 5)
	//resultSubtract := subtract(10, 8)
	//
	//fmt.Println("Result the split function: ", resultSplit)
	//fmt.Println("Result the add function: ", resultAdd)
	//fmt.Println("Result the subtract function: ", resultSubtract)

	//sliceInt := []internal.TypeInteger{{1}, {2}, {3}, {4}, {5}}
	//sliceString := []internal.TypeString{{"seis"}, {"siete"}, {"ocho"}, {"nueve"}, {"diez"}}
	//
	//fmt.Println(internal.JoinSlice(sliceInt, sliceString))

	//library := map[string]internal.Book{
	//	"Deep Work": {
	//		Title:  "Deep Work",
	//		Author: "Pepito",
	//		Year:   2020,
	//	},
	//	"Habitos atomicos": {
	//		Title:  "Habitos atomicos",
	//		Author: "Roberto",
	//		Year:   2019,
	//	},
	//	"Besame mucho": {
	//		Title:  "Besame mucho",
	//		Author: "CarlOs",
	//		Year:   2015,
	//	},
	//}
	//
	//bookAdd := internal.Book{
	//	Title:  "Alto libro",
	//	Author: "Milanesio",
	//	Year:   1985,
	//}
	//
	//bookEdit := internal.Book{
	//	Title:  "Besame Mucho",
	//	Author: "Besame Mucho",
	//	Year:   1775,
	//}
	//
	//err := internal.AddBook(library, bookAdd)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(library)
	//}
	//
	//err = internal.DeleteBook(library, "Alto libro")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(library)
	//}
	//
	//err, b := internal.GetBook(library, "Besame mucho")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(b)
	//}
	//
	//err, _ = internal.EditBook(library, "Besame mucho", bookEdit)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(library)
	//}
	//
	//err, b = internal.GetBook(library, "Besame mucho")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(b)
	//}

	//internal.SumChannel()

	//pages := []string{
	//	"https://google.com",
	//	"https://go.dev/tour/welcome/1",
	//	"https://translate.google.com.ar/",
	//}
	//internal.OpenPages(pages)

	//numeros := []int{8, 2, 20, 4, 56, 6}
	//
	//internal.IsPairChannel(numeros)

	//circulo := internal.Circulo{
	//	Nombre: "Circulo",
	//	Radio:  2.5,
	//}
	//
	//cuadrado := internal.Cuadrado{
	//	Nombre: "Cuadrado",
	//	Lado:   5.0,
	//}
	//
	//rectangulo := internal.Rectangulo{
	//	Nombre:    "Rectangulo",
	//	LadoLargo: 3.5,
	//	LadoCorto: 2.7,
	//}
	//
	//internal.NombreFigura(circulo)
	//internal.Draw(circulo)
	//internal.NombreFigura(cuadrado)
	//internal.Draw(cuadrado)
	//internal.NombreFigura(rectangulo)
	//internal.Draw(rectangulo)

	//content, err := internal.ReadFile("textPrueba.txt")
	//if err != nil {
	//	fmt.Println("Error reading file: ", err)
	//} else {
	//	fmt.Println(content)
	//}

	//result, err := internal.Divide(8.0, 0)
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//} else {
	//	fmt.Println("Result: ", result)
	//}

	//err := internal.Authentication("Luciano", "contra")
	//if err != nil {
	//	if errors.As(err, &internal.UserNotFoundError{}) {
	//		fmt.Println("Error authentication by User. Error: ", err)
	//	} else if errors.Is(err, internal.PasswordIncorrectError{}) {
	//		fmt.Println("Error authentication by Password: ", err)
	//	} else {
	//		fmt.Println("An unexpected error ocurred.")
	//	}
	//} else {
	//	fmt.Println("Authentication Successful.")
	//}

	//a := 10
	//b := 5
	//
	//resultSum := mathutil.Sum(a, b)
	//resultRest := mathutil.Rest(a, b)
	//resultMultiplication := mathutil.Multiplication(a, b)
	//resultDivisionInt := mathutil.DivisionInt(a, b)
	//
	//fmt.Println("Suma =", resultSum)
	//fmt.Println("Resta =", resultRest)
	//fmt.Println("Multiplicacion =", resultMultiplication)
	//fmt.Println("Division =", resultDivisionInt)

	//uuid := uuidutil.GenerateUUID()
	//fmt.Println(uuid)

	//a := 5
	//b := 3
	//product := mathutil.Multiplication(a, b)
	//uuid := stringutil.GenerateUUID()
	//
	//fmt.Printf("The product of %d and %d is %d\n", a, b, product)
	//fmt.Printf("Generated UUID: %s\n", uuid)

	//http.HandleFunc("/hello", internal.HelloHandler)
	//http.HandleFunc("/echo", internal.EchoHandler)
	//
	//fmt.Println("Starting server on :8080")
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	fmt.Println("Error starting server:", err)
	//}
	//
	//resp, err := http.Get("www.lalal.com")
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//fmt.Println("response code: ", resp.StatusCode)
	//
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//fmt.Println("Response Body:", string(body))
	//
	//data := map[string]interface{}{
	//	"title":  "foo",
	//	"body":   "bar",
	//	"userId": 1,
	//}
	//jsonData, err := json.Marshal(data)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//
	//response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//defer response.Body.Close()
	//
	//fmt.Println("Status Code:", response.StatusCode)
	//
	//body, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//fmt.Println("Response Body:", string(body))

	//file, err := internal.OpenFile("archivoNuevo.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(file)

	//err := internal.CreateFile("archivoCaca.txt")
	//if err != nil {
	//	fmt.Println("error: ", err)
	//}
	//data, err := internal.ReaderFile("archivoCaca.txt")
	//fmt.Println(data)

	persona := internal.Person{
		Name:  "Luciano",
		Age:   32,
		Email: "luciano@gonzalez",
	}
	//
	//resultJson, err := internal.DecodeJson(persona)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(string(resultJson))
	//
	//json := `{"Name": "John Doe", "Age": 30, "Email": "john@example.com"}`
	//p, err := internal.UncodeJson([]byte(json))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(p)

	//archivo, err := internal.ReadArchive("data.csv")
	//if err != nil {
	//	return
	//}
	//
	//data, err := internal.ConvertToPersonStruct(archivo)
	//
	//for _, p := range data {
	//	fmt.Println(p.Name)
	//}
	//
	//humansFilter := internal.FilterHumanByAge(data, 25)
	//
	//for _, p := range humansFilter {
	//	fmt.Println(p.Name)
	//}

	var p1 *internal.Person

	p1 = &persona
	p2 := p1

	p2.Name = "NuevoNombre"

	var lala []internal.Person

	lala = append(lala, *p2)

	lala[0].Name = "Ricardo"

	fmt.Println("p1: ", p1.Name)
	fmt.Println("p2: ", p2.Name)
	fmt.Println("p2: ", &lala[0].Name)
	fmt.Println("p2: ", &p2.Name)

	//pruebaaa
}
