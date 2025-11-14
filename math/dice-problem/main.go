package main

import "fmt"

func DiceProblem(face int) (int, error) {
	if face < 1 || face > 6 {
		return 0, fmt.Errorf("face must be between 1 and 6")
	}
	return 7 - face, nil
}

func main() {
    face := 5
    result, err := DiceProblem(face)
    if err != nil {
        fmt.Println("error:", err)
        return
    }
    fmt.Println(result)
}


