package unluckMain

import "fmt"

func ClosureUnluck(nameVal string) func(src interface{}) {
	name := nameVal
	unluckCnt := 0

	return func(src interface{}) {
		switch val := src.(type) {
		case int:
			unluckCnt++
			fmt.Printf("%s dropped %d yen.\n", name, val)
			fmt.Printf("Total unluck count=%d.\n", unluckCnt)

		case string:
			unluckCnt++
			fmt.Printf("%s dropped %s.\n", name, val)
			fmt.Printf("Total unluck count=%d.\n", unluckCnt)
		default:
			fmt.Printf("type unsupported.\n")
			fmt.Printf("Total unluck count=%d.\n", unluckCnt)
		}
	}
}
