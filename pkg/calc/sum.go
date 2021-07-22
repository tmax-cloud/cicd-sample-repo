//test/pkg/calc/sum.go
package calc

//Sum is to add numbers
func Sum(a ...int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}
