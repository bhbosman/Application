package DFA

import (
	"fmt"
	"strings"
)

func PlainNodeMultiLink(a, b byte, start, end *PlainNode) error {
	errorList := make([]string, 0)
	for i := a; i <= b; i++ {
		_, ok := start.nextNode[i]
		if ok {
			errorList = append(
				errorList,
				fmt.Sprintf("there is already a link from %v using %v", start.Name, i))
		}
	}
	if len(errorList) > 0{
		return fmt.Errorf(strings.Join(errorList, "\n"))
	}
	for i := a; i <= b; i++ {
		start.nextNode[i] = end
	}
	return nil
}
