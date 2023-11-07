package main

func main() {
	myMap := map[string]string{
		"first key":  "first value",
		"second key": "second value",
		"third key":  "third value",
		"fourth key": "fourth value",
		"fifth key":  "fifth value",
	}
}

// With an empty line it's slightly better, but still not great.
func bar(s string,
	i int,
) {
	println("bar")
}

//go:noinline

// Foo is awesome.
func Foo() {}

var _ = [][]int{{1}}

const perm = 0o755
