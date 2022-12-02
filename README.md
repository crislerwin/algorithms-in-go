Roteiro

Sequência de Fibonacci"

Em matemática, os números de Fibonacci, comumente denominados Fn , formam uma sequência, a sequência de Fibonacci, na qual cada número é a soma dos dois anteriores. A sequência geralmente começa de 0 e 1, embora alguns autores comecem a sequência de 1 e 1 ou às vezes (como Fibonacci) de 1 e 2. Começando de 0 e 1, os primeiros valores na sequência são: [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144]

- Matriz Esta função calcula o n-ésimo número de Fibonacci usando o método da matriz

```go
    func Matrix(n uint) uint {
	a, b := 1, 1
	c, rc, tc := 1, 0, 0
	d, rd := 0, 1

	for n != 0 {
		if n&1 == 1 {
			tc = rc
			rc = rc*a + rd*c
			rd = tc*b + rd*d
		}

		ta := a
		tb := b
		tc = c
		a = a*a + b*c
		b = ta*b + b*d
		c = c*ta + d*c
		d = tc*tb + d*d

		n >>= 1
	}
	return uint(rc)
}
```

Fórmula Esta função calcula o n-ésimo número de Fibonacci usando a [fórmula](https://en.wikipedia.org/wiki/Fibonacci_number#Relation_to_the_golden_ratio)
// Atenção! Testes para valores grandes caem devido ao erro de arredondamento de números de ponto flutuante, funciona bem, apenas em números pequenos

```go
func Formula(n uint) uint {
	sqrt5 := math.Sqrt(5)
	phi := (sqrt5 + 1) / 2
	powPhi := math.Pow(phi, float64(n))
	return uint(powPhi/sqrt5 + 0.5)
}
```
