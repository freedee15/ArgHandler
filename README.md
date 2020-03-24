# ArgHandler
A command-line argument handler for Golang

`go get github.com/freedee15/ArgHandler`

Version 0.1

## Usage:

### Step by Step:

- Create `*parameter` object(s): 

  `yourParameter, _ := NewParameter("h", "help", []string{"arg1", "arg2"})`

- Create `ParameterArray` and add your parameter object(s) to it: 

  `yourParameterArray := ArgHandler.ParameterArray{yourParameter}`

- Create `ArgHandler` object with the parameters: 

  `(yourParameterArray, os.Args[1:])`

  **NOTE:** *If you do not use the* `[1:]` *argument on* `os.Args` *it will include the program you are running in the arguments and the package will not work.*

- To access the argument data, use `ArgHandler.Results[yourParameter]`.

  - If the parameter was not used at all, it will have a value of `""`.
  - If the parameter was used with no arguments, it will have a value of `ArgHandler.NOARGS`, which has a value of `" "`.
  - If the parameter was used with arguments, it will have a value of `arguments`.
  
### Example program:

main.go:

```go
package main

import (
  "fmt"
  "github.com/freedee15/ArgHandler
)

func main() {

  var err error
  
  yourParameter, err := ArgHandler.NewParameter("h", "help", []string{"arg1", "arg2"})
  if err != nil {
    log.Errorln(err.Error())
  }
  
  yourParameterArray := ArgHandler.ParameterArray{yourParameter}
  
  argHandler, err := ArgHandler.NewArgHandler(yourParameterArray, os.Args[1:])
  if err != nil {
    log.Errorln(err.Error())
  }
  
  if argHandler.Results[yourParameter] == ArgHandler.NOARGS {
    fmt.Println("No arguments given")
  } else if argHandler.Results[yourParameter] == "arg1" {
  
    fmt.Println("Arg 1 given")
  
  } else if argHandler.Results[yourParameter] == "arg2" {
  
    fmt.Println("Arg 2 given")
  
  }

}
```
  
 
## Known issues to fix:
- No support for multi-argument scanning.

Report any bugs to noahfriedman2@gmail.com
