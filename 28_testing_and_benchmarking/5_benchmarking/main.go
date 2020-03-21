package saying

import "fmt"

/*
- Benchmarking allows you to measure the performance of your code
*/

func Greet(s string) string {
	return fmt.Sprint("welcome my dear ", s)
}
