# ArgHandler
A command-line argument handler for Golang

`go get github.com/freedee15/ArgHandler`

## Version 0.2 CHANGELOG
- Added built in `-h`/`--help` recognition
- Added support for multi-argument parameters
- Added option to allow undefined arguments to pass to parameters


## Usage:

### Required Imports:

```go
import (
    "github.com/freedee15/ArgHandler"
    "errors"
    "os"
)
```

### Step by Step:

- Create a `ParameterArray`: 

  `yourParameterArray := ArgHandler.ParameterArray{}`

- Create `*parameter` object(s), and for each check for an error. If there is no error, append it to the array: 

```go
yourParameter, err := ArgHandler.NewParameter("p", "param", []string{"arg1", "arg2"}, "Description", 0) //The 0 means take 0 args
if err != nil {
    fmt.Println(err.Error())
} else {
    yourParameterArray = append(yourParameterArray, yourParameter)
}
```

- Create `ArgHandler` object with the parameters `(yourParameterArray, os.Args[1:])` and check for errors:
    
    **NOTE:** *If you do not use the* `[1:]` *argument on* `os.Args` *it will include the program you are running in the arguments and the package will not work.*

 ```go
yourArgHandler, err := ArgHandler.NewArgHandler(yourParameterArray, os.Args[1:])
if err != nil {
    fmt.Println(err.Error())
}
 ```

- To access the argument data, use `ArgHandler.Results[yourParameter]`.

  - If the parameter was not used at all, it will return a nil string array.
  - If the parameter was used with no arguments, it will return a string array with one element, the element being the `short` value of the parameter.
  - If the parameter was used with arguments, it will return a string array with multiple elements, the first being the `short` value of the parameter and the rest being the arguments called on it.
  
### Example program:

main.go:

```go
package main

import (
    "github.com/freedee15/ArgHandler"
    "errors"
    "fmt"
    "os"
)

func main() {
	
    yourParameterArray, err := ArgHandler.ParameterArray{}

    yourParameter, err := ArgHandler.NewParameter("p", "param", []string{"arg1", "arg2"}, "Description", 1)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        yourParameterArray = append(yourParameterArray, yourParameter)
    }

    yourArgHandler, err := ArgHandler.NewArgHandler(yourParameterArray, os.Args[1:])
    if err != nil {
        fmt.Println(err.Error())
    } else {
        if yourArgHandler.Results[yourParameter] != nil {
            for _, recievedArgument := range yourArgHandler.Results[yourParameter] {

                fmt.Printf("Recieved argument: %s\n", recievedArgument)

            }
        }   
    }

}
```
  
 
## Known issues to fix:
None at the moment, but I'm sure there will be some soon!

Report any bugs or comments to [noahfriedman2@gmail.com](mailto:noahfriedman2@gmail.com)
