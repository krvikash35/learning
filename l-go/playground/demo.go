package main

import "log"

func main() {
	// c := make(chan int, 1)
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// }()
	// c <- 1
	// c <- 1

	obj := map[string]interface{}{
		"exp1": true,
		"exp2": false,
		"exp3": "true",
	}

	value, keyExist := obj["exp1"]
	log.Printf("value=%+v keyExist=%+v valueType=%T", value, keyExist, value)

	value, keyExist = obj["exp2"]
	log.Printf("value=%+v keyExist=%+v valueType=%T", value, keyExist, value)

	value, keyExist = obj["exp3"]
	log.Printf("value=%+v keyExist=%+v valueType=%T", value, keyExist, value)

	value, keyExist = obj["exp1"].(bool)
	log.Printf("value=%+v keyExist=%+v valueType=%T", value, keyExist, value)

	value, keyExist = obj["exp2"].(bool)
	log.Printf("value=%+v keyExist=%+v valueType=%T", value, keyExist, value)

	value, keyExist = obj["exp3"].(bool)
	log.Printf("value=%+v keyExist=%+v valueType=%T", value, keyExist, value)

}
