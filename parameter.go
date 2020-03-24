package ArgHandler

import (
	"errors"
	"fmt"
	"strings"
)

//TODO: Add option to define how many arguments a parameter can take

type ParameterArray []*parameter

type parameter struct {
	short     string
	long      string
	validArgs []string
}

func NewParameter(short string, long string, validArgs []string) (*parameter, error) {

	//Check if "short" argument meets requirements
	if len(short) != 1 {
		return nil, errors.New("short argument required to create parameter, must be one character long")
	}

	//Remove empty argument parameters
	warn := true
	for i := 0; i < len(validArgs); i++ {

		//Remove whitespace in arguments
		newStr := strings.Replace(validArgs[i], " ", "", -1)
		if newStr != validArgs[i] {
			fmt.Printf("WARNING: parameter \"%s\" given for argument \"-%s\" contained whitespace. It has been utterly smashed!\n", validArgs[i], short)
		}
		validArgs[i] = newStr

		//Remove empty arguments
		if validArgs[i] == "" {
			if warn {
				fmt.Printf("WARNING: One or more parameters given for the argument \"-%s\" are empty. I will now OBLITERATE them.\n", short)
				warn = false
			}
			validArgs = append(validArgs[:i], validArgs[i+1:]...)
			i--
		}

	}

	return &parameter{short: short, long: long, validArgs: validArgs}, nil
}
