package main

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

	Store[args[0].bulk] = args[1].bulk

	return Value{typ: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1{
		return Value{typ: "string", str: "GET expects 1 argument"}
	}

	return Value{typ: "bulk", bulk: Store[args[0].bulk]}
}

func keys(_ []Value) Value{
	values := Value{typ: "array", array: []Value{}}

	for key, _ :=range(Store){
		current := Value{typ: "string", str: key}
		values.array = append(values.array, current)
	}

	return values
}

var Handlers = map[string]func([]Value) Value{
	"PING": ping,
	"SET": set,
	"GET": get,
	"KEYS": keys,
}