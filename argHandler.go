package ArgHandler

import (
	"errors"
)

type argHandler struct {
	Results map[*parameter]string
}

func NewArgHandler(parameters ParameterArray, args []string) (*argHandler, error) {

	var results map[*parameter]string
	results = make(map[*parameter]string)
	var currentParam *parameter

	for _, a := range args {

		sub := []rune(a)
		if string(sub[0]) == "-" {

			if string(sub[1]) == "-" {

				//Cycle through long parameters (e.g. "--help") to see if parameter exists
				for _, p := range parameters {

					if p.long == string(sub[2:]) {
						currentParam = p
						break
					}

				}

				//Crash if that didn't work
				if currentParam == nil {
					return nil, errors.New("no parameter to give option \"" + a + "\" to")
				}
				if currentParam.long != string(sub[2:]) {
					return nil, errors.New("invalid parameter given: " + a)
				}

			} else {

				if len(a) != 2 {
					return nil, errors.New("invalid parameter given: " + a)
				}

				//Cycle through short parameters (e.g. "-h") to see if parameter exists
				for _, p := range parameters {

					if p.short == string(sub[1]) {
						currentParam = p
						break
					}

				}

				//Crash if that didn't work
				if currentParam == nil {
					return nil, errors.New("no parameter to give option \"" + a + "\" to")
				}
				if currentParam.short != string(sub[1]) {
					return nil, errors.New("invalid parameter given: " + a)
				}

			}
		} else {

			//Check if there's current parameter to send arguments to, if not then crash
			if currentParam == nil {
				return nil, errors.New("no parameter to give option \"" + a + "\" to")
			} else {
				//If the parameter supports the argument, keep going. Otherwise crash
				contains := false
				for _, b := range currentParam.validArgs {

					if a == b {
						contains = true
						break
					}

				}
				if contains {
					results[currentParam] = a
					currentParam = nil
				} else {
					return nil, errors.New("parameter \"" + currentParam.short + "\" does not take argument \"" + a + "\"")
				}
			}

		}
	}

	return &argHandler{Results: results}, nil

}
