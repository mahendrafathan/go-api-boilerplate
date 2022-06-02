package helpers

import "fmt"

// fungsi untuk membagi 2 bilangan
// input a penyebut, b pembagi (tidak boleh 0)
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("pembagi tidak boleh 0")
	}

	return a / b, nil
}
