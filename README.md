# jexp
Process expressions defined in JSON format

## Installation
```
go get github.com/duhnnie/jexp
```

## Example
```go
package main

import (
	"fmt"

	"github.com/duhnnie/jexp"
	"github.com/duhnnie/valuebox"
)

var jsonData = []byte(`
{
  "type": "eq",
  "operands": [
    {
      "type": "subs",
      "operands": [
        {
          "type": "var",
          "dataType": "number",
          "value": "thompsons.0.age"
        },
        {
          "type": "var",
          "dataType": "number",
          "value": "thompsons.1.age"
        }
      ]
    },
    {
      "type": "subs",
      "operands": [
        {
          "type": "var",
          "dataType": "number",
          "value": "larsons.0.age"
        },
        {
          "type": "var",
          "dataType": "number",
          "value": "larsons.1.age"
        }
      ]
    }
  ]
}
`)

func main() {
	thompsons := []byte(`
		[
        {
            "name": "Jimmy",
            "age": 35
        },
        {
            "name": "Melissa",
            "age": 25
        }
    ]
	`)

	larsons := []byte(`
    [
      {
          "name": "Rebecca",
          "age": 15
      },
      {
          "name": "Brad",
          "age": 5
      }
    ]
	`)

	// Here we're using valuebox module,
	// but you can use any other struct that conforms
	// to the ExpressionContext interface.
	ctx := valuebox.New()
	ctx.Set("thompsons", thompsons)
	ctx.Set("larsons", larsons)

	exp, errPath, err := jexp.New[bool](jsonData)

	if err != nil {
		fmt.Println("Error: at creating expression", errPath, err)
		return
	}

	r, errPath2, err2 := exp.Resolve(ctx)

	if err != nil {
		fmt.Printf("Error at evaluating expression at %s: %s\n", errPath2, err2)
	}

	if r {
		fmt.Println("Thompson and Larson brothers have the same age gap")
	} else {
		fmt.Println("Thompson and Larson brothers don't have the same age gap")
	}
}

```

## Note
This project currently supports a few expressions. Contributions are welcome!