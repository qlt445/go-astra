package main

import (
	"encoding/json"
	"fmt"
	"github.com/vetcher/go-astra"
)

func main() {
	path := "/Users/qlt/git_repo/gameserver/internal/game/constants.go"
	file, err := astra.ParseFile(path, astra.IgnoreVariables)
	if err != nil {
		fmt.Println(err)
	}
	t, err := json.MarshalIndent(file, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(t))
}
