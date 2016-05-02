## Description

We are going to learn how to encode and decode JSON in Go. You've been given some code that reads data from a file and decodes the JSON into a struct. Your job is to complete the implementation.

## Comprehension Tasks

The json package implements encoding and decoding of JSON objects. The mapping between JSON objects and Go values is described in the documentation for the Marshal and Unmarshal functions.

Check out the code in [exhibit-a]() to see how the raw phone data is used by the application below.

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Phone ...
type Phone struct {
    Name string `json:"name"`
}

var allPhones []Phone

func setup() {
    data, err := ioutil.ReadFile("phones.json")
    if err != nil {
        fmt.Println("Error reading phones.json")
        os.Exit(1)
      }

    err = json.Unmarshal(data, &allPhones)
    if err != nil {
        fmt.Println("Error in unmarshalling phones")
    }
}

func phones(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(allPhones)
}

func main() {
    setup()
    http.HandleFunc("/phones", phones)
    http.ListenAndServe(":8080", nil)
}
```

Explain what each line of code does in `exhibit-a`.

## Engineering Tasks

1. Read the other attributes from the file and store them in the Phone struct
2. Create a database table for phones. When the server is started, if the table is empty, seed the table with the phones from the json file. The /phones endpoint should then return whatever records are found in the phones table.
