package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Misora000/gobaidu/search"
)

func main() {
	res, err := search.Search(context.Background(), "鋼鐵人", 4)
	if err != nil {
		panic(err)
	}

	result, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}
