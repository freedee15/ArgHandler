# ArgHandler
A command-line argument handler for Golang

`go get github.com/freedee15/ArgHandler`

Version 0.1

## Usage (in order):

- Create `*parameter` object(s): 

  `yourParameter := NewParameter("h", "help", []string{"arg1", "arg2"})`

- Create `ParameterArray` and add your parameter object(s) to it: 

  `yourParameterArray := ArgHandler.ParameterArray{yourParameter}`

- Create `ArgHandler` object with the parameters: 

  `(yourParameterArray ParameterArray, os.Args[1:])`

  **NOTE:** *If you do not use the* `[1:]` *argument on* `os.Args` *it will include the program you are running in the arguments and the package will not work.*

- To access the argument data, use `ArgHandler.Results[yourParameter]`.

  - If the parameter was not used at all, it will have a value of `""`.
  - If the parameter was used with no arguments, it will have a value of `ArgHandler.NOARGS`, which has a value of `" "`.
  - If the parameter was used with arguments, it will have a value of `arguments`.
 
## Known issues to fix:
- If an unrecognized parameter and/or argument is used the program will crash.
- No support for multi-argument scanning.

Report any bugs to noahfriedman2@gmail.com
