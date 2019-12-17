package responder

import (
	"fmt"

	"github.com/family/translator"
)

type _response int

// Enum-like const declaration for supported response types
const (
	ChildAdditionSuccessful _response = iota
	ChildAdditionFailed     _response = iota
	PersonNotFound          _response = iota
	None                    _response = iota
)

// Response emits the params in form of a reponse
func Response(reponseType _response) {
	switch reponseType {
	case ChildAdditionSuccessful:
		translator.Get("child_addition_succeeded", "en")
		break
	case ChildAdditionFailed:
		translator.Get("child_addition_failed", "en")
		break
	case PersonNotFound:
		translator.Get("person_not_found", "en")
		break
	case None:
		translator.Get("none", "en")
		break
	}
}

// ResponseWithData emits the params in form of a reponse
func ResponseWithData(data string) {
	fmt.Println(data)
}
