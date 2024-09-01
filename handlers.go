package main

import (
	"fmt"
	"strconv"
)

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}

	return Value{typ: "string", str: args[0].bulk}
}

func set(args []Value) Value {
	if len(args) != 2{
		return Value{typ: "string", str: "SET expects 2 arguments"}
	}

	if Store[Index] == nil{
		Store[Index] = make(store)
	}

	Store[Index][args[0].bulk] = args[1].bulk

	return Value{typ: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1{
		return Value{typ: "string", str: "GET expects 1 argument"}
	}

	return Value{typ: "bulk", bulk: Store[Index][args[0].bulk]}
}

func keys(_ []Value) Value{
	values := Value{typ: "array", array: []Value{}}

	for key, _ :=range(Store[Index]){
		current := Value{typ: "string", str: key}
		values.array = append(values.array, current)
	}

	return values
}

func selectindex(args []Value) Value{
	if len(args) != 1 {
		return Value{typ: "string", str: "SELECT expects 1 argument"}
	}

	index, err := strconv.ParseInt(args[0].bulk, 10, 64)

	if err!=nil{
		return Value{typ: "string", str: err.Error()}
	}

	if index<0 || index > 9{
		return Value{typ: "error", str: "Invalid index!"}
	}

	Index = int(index)
		
	msg := fmt.Sprintf("selected index %d", index)

	return Value{typ: "string", str: msg}
}

var Handlers = map[string]func([]Value) Value{
	"PING": ping,
	"SET": set,
	"GET": get,
	"KEYS": keys,
	"SELECT": selectindex,
}