package internal

// SequenceGenerator retorna una función que incrementa y devuelve un contador
func SequenceGenerator() func() int {
	count := 1 // Variable capturada por el closure
	// Closure que incrementa y devuelve el contador
	return func() int {
		count++
		return count
	}
}

func SequecenJumperGenerator(init, jump int) func() int {
	count := init - jump
	return func() int {
		count += jump
		return count
	}
}

func SequecenJumperGeneratorWithRestart(init, jump int) (func() int, func()) {
	count := init - jump
	return func() int {
			count += jump
			return count
		}, func() {
			count = init - jump
		}
}

type Operation func(int, int) int

func ExecuteOperation(a, b int, operation Operation) int {
	return operation(a, b)
}
