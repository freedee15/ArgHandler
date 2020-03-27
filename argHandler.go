package ArgHandler

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

var helpHasBeenDisplayed bool

type argHandler struct {
	Results     map[*parameter][]string
	HelpToggled bool
}

func NewArgHandler(parameters ParameterArray, args []string) (*argHandler, error) {

	helpHasBeenDisplayed = false
	var results map[*parameter][]string
	results = make(map[*parameter][]string)
	var currentParam *parameter
	var currentIndex int
	helpToggled := false

	for i, a := range args {

		sub := []rune(a)
		if string(sub[0]) == "-" {

			if string(sub[1]) == "-" {

				if a == "--help" {
					displayHelp(&parameters)
					helpToggled = true
				} else {

					//Cycle through long parameters (e.g. "--help") to see if parameter exists
					for _, p := range parameters {

						if p.Long == string(sub[2:]) {
							if currentParam == nil {
								currentParam = p
								results[currentParam] = append(results[currentParam], currentParam.Short)
								currentIndex = i
								break
							} else {
								displayHelp(&parameters)
								return nil, errors.New("not enough arguments given to parameter \"" + currentParam.Short + "\"")
							}
						}

					}

					//Crash if that didn't work
					if currentParam == nil {
						displayHelp(&parameters)
						return nil, errors.New("invalid parameter given: " + a)
					}
					if currentParam.Long != string(sub[2:]) {
						displayHelp(&parameters)
						return nil, errors.New("invalid parameter given: " + a)
					}

				}

			} else {

				if a == "-h" {
					displayHelp(&parameters)
					helpToggled = true
				} else {

					if len(a) != 2 {
						displayHelp(&parameters)
						return nil, errors.New("invalid parameter given: " + a)
					}

					//Cycle through short parameters (e.g. "-h") to see if parameter exists
					for _, p := range parameters {

						if p.Short == string(sub[1]) {
							if currentParam == nil {
								currentParam = p
								results[currentParam] = append(results[currentParam], currentParam.Short)
								currentIndex = i
								break
							} else {
								displayHelp(&parameters)
								return nil, errors.New("not enough arguments given to parameter \"" + currentParam.Short + "\"")
							}
						}

					}

					//Crash if that didn't work
					if currentParam == nil {
						displayHelp(&parameters)
						return nil, errors.New("invalid parameter given: " + a)
					}
					if currentParam.Short != string(sub[1]) {
						displayHelp(&parameters)
						return nil, errors.New("invalid parameter given: " + a)
					}

				}

			}

		} else {

			//Check if there's current parameter to send arguments to, if not then crash
			if currentParam == nil {
				displayHelp(&parameters)
				return nil, errors.New("no parameter to give option \"" + a + "\" to")
			} else {

				if currentParam.validArgs == nil {
					results[currentParam] = append(results[currentParam], a)
				} else {

					//Check if parameter supports arg
					contains := false
					duplicate := false
					for _, b := range currentParam.validArgs {

						if a == b {
							contains = true
							break
						}

					}

					//Check for duplicates
					if contains {
						for _, b := range results[currentParam] {

							if a == b {
								duplicate = true
								break
							}

						}
					}

					//If it doesn't support the argument fail. If it's a duplicate, ignore it
					if contains && !duplicate {
						results[currentParam] = append(results[currentParam], a)
					} else if !contains {
						displayHelp(&parameters)
						return nil, errors.New("parameter \"" + currentParam.Short + "\" does not take argument \"" + a + "\"")
					} else if duplicate {
						log.Printf("WARNING: Duplicate argument given to parameter \"%s\". I'm ignoring it. LALALALA I can't hear you duplicate argument!\n", currentParam.Short)
					}
				}
			}

		}

		//Check if currentParam has used up it's args, and if it's gone too far just crash because that should be impossible
		if currentParam != nil {
			if i-currentIndex == currentParam.args {
				currentParam = nil
			} else if i-currentIndex > currentParam.args {
				log.Fatalln("WOW, you really broke something. This should be impossible.")
			}
		}

	}

	if currentParam != nil {
		return nil, errors.New("not enough arguments passed to parameter \"" + currentParam.Short + "\"")
	}

	return &argHandler{Results: results, HelpToggled: helpToggled}, nil

}

func displayHelp(parameters *ParameterArray) {

	if !helpHasBeenDisplayed {

		//Format all the parameter definitions together
		lenCheck := []string{"h, --help     "}
		for _, i := range *parameters {

			lenStr := i.Short
			if i.Long != "" {
				lenStr += ", --" + i.Long
			}
			lenStr += "     "
			lenCheck = append(lenCheck, lenStr)

		}

		//Figure out which one is the longest
		var maxLen int
		for _, i := range lenCheck {

			if len(i) > maxLen {
				maxLen = len(i)
			}

		}

		//Format them all to be that length
		for a, i := range lenCheck {

			for len(lenCheck[a]) < maxLen {

				i += " "
				lenCheck[a] = i

			}

		}

		//Add descriptions onto the back
		for a, i := range lenCheck {

			if a == 0 {
				i += "Display this help text"
			} else {
				i += (*parameters)[a-1].description
			}
			lenCheck[a] = i

		}

		//Sort them all and then add a "-" to the beginning of each
		sort.Strings(lenCheck)
		for a, i := range lenCheck {

			i = "-" + i
			lenCheck[a] = i

		}

		//Print everything out
		for _, i := range lenCheck {
			fmt.Println(i)
		}

		//Toggle bool so help only gets displayed once
		helpHasBeenDisplayed = true

	}

}
