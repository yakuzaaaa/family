package utils

import (
	"strings"

	"github.com/family/models"
	"github.com/family/responder"
)

// GetSon returns son of the person
func GetSon(person *models.Node) {
	GetChildrenByGender(person, models.MALE)
}

// GetDaughter returns daughter of the person
func GetDaughter(person *models.Node) {
	GetChildrenByGender(person, models.FEMALE)
}

// GetSiblings returns siblings of the person
func GetSiblings(person *models.Node) {
	GetSiblingsByGender(person, -1)
}

// GetPaternalUncle returns paternal uncle of the person
func GetPaternalUncle(person *models.Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Husband, models.MALE)
	} else {
		responder.Response(responder.None)
	}
}

// GetPaternalAunt returns paternal aunt of the person
func GetPaternalAunt(person *models.Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Husband, models.FEMALE)
	} else {
		responder.Response(responder.None)
	}
}

// GetMaternalUncle returns maternal uncle of the person
func GetMaternalUncle(person *models.Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Wife, models.MALE)
	} else {
		responder.Response(responder.None)
	}
}

// GetMaternalAunt returns maternal aunt of the person
func GetMaternalAunt(person *models.Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Wife, models.FEMALE)
	} else {
		responder.Response(responder.None)
	}
}

// GetSisterInLaw returns sister-in-law of the person
func GetSisterInLaw(person *models.Node) {
	GetSiblingsOfInLawByGender(person, models.FEMALE)
}

// GetBrotherInLaw returns brother-in-law of the person
func GetBrotherInLaw(person *models.Node) {
	GetSiblingsOfInLawByGender(person, models.MALE)
}

// GetSiblingsOfInLawByGender returns siblings-in-law of the person by gender
func GetSiblingsOfInLawByGender(person *models.Node, gender int) {
	spouse := GetSpouse(person)
	if spouse != nil {
		GetSiblingsByGender(spouse, gender)
	} else {
		responder.Response(responder.None)
	}
}

// GetChildrenByGender returns children of the person by gender
func GetChildrenByGender(person *models.Node, gender int) {
	couple := person.MarriageDetails
	children := ""
	if couple != nil && len(couple.Children) > 0 {
		for _, child := range couple.Children {
			if child.Gender == gender {
				children += child.Name + " "
			}
		}

		if len(children) > 0 {
			responder.ResponseWithData(strings.TrimRight(children, " "))
		} else {
			responder.Response(responder.None)
		}

	} else {
		responder.Response(responder.None)
	}
}

// GetSiblingsByGender returns siblings of the person by gender
func GetSiblingsByGender(person *models.Node, gender int) {
	parent := person.Parent
	siblings := ""

	if parent != nil && len(parent.Children) > 0 {
		for _, child := range parent.Children {
			if child.Name != person.Name && (gender == -1 || child.Gender == gender) {
				siblings += child.Name + " "
			}
		}

		if len(siblings) > 0 {
			responder.ResponseWithData(strings.TrimRight(siblings, " "))
		} else {
			responder.Response(responder.None)
		}

	} else {
		responder.Response(responder.None)
	}
}

// GetSpouse returns spouse of the person
func GetSpouse(person *models.Node) *models.Node {
	if person.MarriageDetails != nil {
		couple := person.MarriageDetails
		if person.Gender == models.MALE {
			return couple.Wife
		}
		return couple.Husband
	}

	return nil
}

// GetPerson return person
func GetPerson(name string) (*models.Node, bool) {
	person, found := models.Persons[name]
	return person, found
}

// GetCouple returns married couple
func GetCouple(name string) (*models.Couple, bool) {
	couple, found := models.Couples[name]
	return couple, found
}
