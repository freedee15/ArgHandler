package ArgHandler

import (
	"errors"
	"log"
	"strings"
)

type ParameterArray []*parameter

type parameter struct {
	Short       string
	Long        string
	validArgs   []string
	description string
	args        int
}

func NewParameter(short, long string, validArgs []string, description string, args int) (*parameter, error) {

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
			log.Printf("WARNING: parameter \"%s\" given for argument \"-%s\" contained whitespace. It has been utterly smashed!\n", validArgs[i], short)
		}
		validArgs[i] = newStr

		//Remove empty arguments
		if validArgs[i] == "" {
			if warn {
				log.Printf("WARNING: One or more parameters given for the argument \"-%s\" are empty. I will now OBLITERATE them.\n", short)
				warn = false
			}
			validArgs = append(validArgs[:i], validArgs[i+1:]...)
			i--
		}

		//Check that args amount makes sense
		if args < 0 {
			log.Printf("WARNING: You can't have a negative amount of arguments, you FOOL! Setting \"args\" for parameter \"%s\" to 0.\n", short)
			args = 0
		} else if args >= 5 {
			log.Printf("WARNING: You have 5 or more arguments for the parameter \"%s\". Are you sure you want to do that? That's like, a whole bunch.\n", short)
		}

		if args == 0 && len(validArgs) > 0 {
			log.Printf("WARNING: You told me not to accept any arguments for parameter \"%s\", but then you gave me things to check for. I'm going to get rid of them now.\n", short)
		}

		//If there aren't enough validArgs to args, set args to amount of validArgs
		if args > len(validArgs) {
			log.Printf("WARNING: You didn't give me enough valid arguments to work with for parameter \"%s\"! That's a little unsafe bud!\n", short)
			args = len(validArgs)
		}

		//If there's no description, warn
		if description == "" {
			log.Printf("WARNING: You haven't given a description for parameter \"%s\"! How is everyone supposed to know what your beautiful parameter does now??\n", short)
		}

	}

	return &parameter{Short: short, Long: long, validArgs: validArgs, description: description, args: args}, nil
}
