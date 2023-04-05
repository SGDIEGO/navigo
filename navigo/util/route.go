package util

func ValidateUrl(absolute, relative string) string {

	if len(relative) == 0 {
		return absolute
	}

	if relative[0] == '/' {
		return ValidateUrl(absolute, relative[1:])
	}

	return absolute + relative
}
