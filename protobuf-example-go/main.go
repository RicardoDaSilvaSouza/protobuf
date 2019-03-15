package main

// import "/Users/ricasis/study-workspace/protobuf/protobuf-example-go/src/simple"
// import "github.com/RicardoDaSilvaSouza/protobuf-example-go/src/simple"

func main() {

	sm := example_simple.SimpleMessage{
		Id:       1,
		IsSimple: true,
		Name:     "Example",
		SimpleList: []string{
			"test",
			"test2",
		},
	}
}
