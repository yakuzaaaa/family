package controller

import (
	"github.com/family/models"
	"github.com/family/responder"
	"github.com/family/utils"
)

func Handle(query string, input []string) {
	switch query {
	case "ADD_CHILD":
		_addChild(string(input[0]), string(input[1]), 1)
		break
	case "GET_RELATIONSHIP":
		_resolveRelation(string(input[0]), string(input[1]))
		break
	}
}

func _addChild(motherName string, childName string, gender int) {
	mother, foundMother := utils.GetPerson(motherName)
	if foundMother {
		parent := mother.MarriageDetails
		if parent != nil {
			child := models.CreateNode(childName, gender, parent)
			parent.AddChild(child)
			responder.Response(responder.ChildAdditionSuccessful)
		} else {
			responder.Response(responder.ChildAdditionFailed)
		}
	} else {
		responder.Response(responder.PersonNotFound)
	}
}

func _resolveRelation(personName string, relationName string) {
	person, found := utils.GetPerson(personName)
	if found {
		switch relationName {
		case models.SON:
			utils.GetSon(person)
			break
		case models.DAUGHTER:
			utils.GetDaughter(person)
			break
		case models.SIBLINGS:
			utils.GetSiblings(person)
			break
		case models.PATERNAL_UNCLE:
			utils.GetPaternalUncle(person)
			break
		case models.PATERNAL_AUNT:
			utils.GetPaternalAunt(person)
			break
		case models.MATERNAL_UNCLE:
			utils.GetMaternalUncle(person)
			break
		case models.MATERNAL_AUNT:
			utils.GetMaternalAunt(person)
			break
		case models.SISTER_IN_LAW:
			utils.GetSiterInLaw(person)
			break
		case models.BROTHER_IN_LAW:
			utils.GetBrotherInLaw(person)
			break
		default:
			responder.Response(responder.None)
		}
	} else {
		responder.Response(responder.PersonNotFound)
	}
}
