package main

import "fmt"

// 错误示例
// func main() {
// 	doIt := func(arg int) interface{} {
// 		var result *struct{} = nil
// 		if arg > 0 {
// 			result = &struct{}{}
// 		}
// 		return result
// 	}

// 	if res := doIt(-1); res != nil {
// 		fmt.Println("Good result: ", res)	// Good result:  <nil>
// 		fmt.Printf("%T\n", res)			// *struct {}	// res 不是 nil，它的值为 nil
// 		fmt.Printf("%v\n", res)			// <nil>
// 	}
// }

// 正确示例
func main() {
	doIt := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		} else {
			return nil // 明确指明返回 nil
		}
		return result
	}

	if res := doIt(-1); res != nil {
		fmt.Println("Good result: ", res)
	} else {
		fmt.Println("Bad result: ", res) // Bad result:  <nil>
	}
}
