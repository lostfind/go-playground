package main

import "fmt"

type Robot struct {
	name  string
	birth int
}

func main() {
	dora := Robot{
		name:  "Doraemon",
		birth: 2112,
	}

	Anything := map[string]interface{}{
		"valString": "文字列",
		"valInt":    1234,
		"valBool":   true,
		"valStruct": dora,
	}

	var intValue int
	for i, v := range Anything {
		switch v := v.(type) { // 元の型キャストされる (switchのスコープ内のみ)
		case int:
			intValue = v
			fmt.Printf("%s:%d (%T)\n", i, v, v) // valInt:1234 (int)
		case string:
			fmt.Printf("%s:%s (%T)\n", i, v, v) // valString:文字列 (string)
		case bool:
			fmt.Printf("%s:%t (%T)\n", i, v, v) // valBool:true (bool)
		case Robot:
			fmt.Printf("%s: %+v (%T)\n", i, v, v) // valStruct: {name:Doraemon birth:2112} (main.Robot)
		}
	}

	fmt.Println(intValue) // 1234
}
