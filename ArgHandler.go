package ArgHandler

import "errors"

type option struct {
	short string
	long  string
	valid []string
}

func NewOption(short string, long string, valid []string) (*option, error) {
	if len(short) != 1 {
		return nil, errors.New("short argument required to create option, must be one character long")
	}
	if len(valid) == 0 {
		return nil, errors.New("at least one valid argument to the option is required")
	}
	return &option{short: short, long: long, valid: valid}, nil
}

type argHandler struct {
}
