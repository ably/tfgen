package tfgen_test

import (
	"bytes"
	"fmt"

	"github.com/ably/tfgen"
)

// Example builds a minimal stack using only the core API and writes it as
// Terraform JSON.
func Example() {
	stack := tfgen.NewStack("example")

	file := &tfgen.Resource{
		Type: "local_file",
		Name: "hello",
		Config: tfgen.Config{
			"content":  "hello, world",
			"filename": "hello.txt",
		},
	}
	stack.AddResource(file)

	// Reference an attribute of the resource in an output.
	stack.AddOutput(&tfgen.Output{
		Name:  "path",
		Value: file.Ref("filename"),
	})

	var buf bytes.Buffer
	if err := stack.Write(&buf); err != nil {
		panic(err)
	}

	fmt.Print(buf.String())
	// Output:
	// {
	//   "output": {
	//     "path": {
	//       "value": "${local_file.hello.filename}"
	//     }
	//   },
	//   "resource": {
	//     "local_file": {
	//       "hello": {
	//         "content": "hello, world",
	//         "filename": "hello.txt"
	//       }
	//     }
	//   }
	// }
}
