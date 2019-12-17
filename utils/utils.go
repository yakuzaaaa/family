package utils

import (
	"fmt"
	"strings"

	"github.com/family/models"
	"github.com/family/responder"
)

func GetSon(person *models.Node) {
	GetChildrenByGender(person, models.MALE)
}

func GetDaughter(person *models.Node) {
	GetChildrenByGender(person, models.FEMALE)
}

func GetSiblings(person *models.Node) {
	GetSiblingsByGender(person, -1)
}

func GetPaternalUncle(person *models.Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Husband, models.MALE)
	} else {
		responder.Response(responder.PersonNotFound)
	}
}

func GetPaternalAunt(person *models.Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Husband, models.FEMALE)
	} else {
		responder.Response(responder.PersonNotFound)
	}
}

func GetMaternalUncle(person *models.Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Wife, models.MALE)
	} else {
		responder.Response(responder.PersonNotFound)
	}
}

func GetMaternalAunt(person *models.Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Wife, models.FEMALE)
	} else {
		responder.Response(responder.PersonNotFound)
	}
}

func GetSiterInLaw(person *models.Node) {
	GetSiblingsOfInLawByGender(person, models.FEMALE)
}

func GetBrotherInLaw(person *models.Node) {
	GetSiblingsOfInLawByGender(person, models.MALE)
}

func GetSiblingsOfInLawByGender(person *models.Node, gender int) {
	spouse := GetSpouse(person)
	if spouse != nil {
		GetSiblingsByGender(spouse, gender)
	} else {
		responder.Response(responder.PersonNotFound)
	}
}

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
		responder.Response(responder.PersonNotFound)
	}
}

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
		responder.Response(responder.PersonNotFound)
	}
}

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

func GetPerson(name string) (*models.Node, bool) {
	person, found := models.Persons[name]
	return person, found
}

func GetCouple(name string) (*models.Couple, bool) {
	couple, found := models.Couples[name]
	return couple, found
}

func printFamilyTree(root *models.Couple) {
	fmt.Println(root.Husband.Name + "   " + root.Wife.Name)
	if len(root.Children) > 0 {
		for _, child := range root.Children {
			fmt.Println(child.Name)
			childCouple, found := models.Couples[child.Name]
			if found {
				printFamilyTree(childCouple)
			}
		}
	}
}
