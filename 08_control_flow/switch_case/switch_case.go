package switch_case

func RunSwitchCase(name string) string {
	switch name {
	case "Tim":
		return "!"
	case "David":
		return "@"
	default:
		return "#"
	}
}

func RunSwitchOnType(x interface{}) string {
	switch x.(type) {
	case string:
		return "This is a string"
	case int:
		return "This is an int"
	default:
		return "whatever"
	}
}
