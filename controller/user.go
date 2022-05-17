package controller

type ReturnValue struct {
	ValueOne string `json:valueOne`
	ValueTwo int    `json:valueTwo`
}

func UserController() ReturnValue {
	return ReturnValue{ValueOne: "hihi", ValueTwo: 10}
}
