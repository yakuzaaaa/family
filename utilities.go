package main

import (
	"fmt"
	"strings"
)

func Response(message string) {
	fmt.Println(message)
}

func GetSon(person *Node) {
	GetChildrenByGender(person, MALE)
}

func GetDaughter(person *Node) {
	GetChildrenByGender(person, FEMALE)
}

func GetSiblings(person *Node) {
	GetSiblingsByGender(person, -1)
}

func GetPaternalUncle(person *Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Husband, MALE)
	} else {
		Response("PERSON_NOT_FOUND")
	}
}

func GetPaternalAunt(person *Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Husband, FEMALE)
	} else {
		Response("PERSON_NOT_FOUND")
	}
}

func GetMaternalUncle(person *Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Wife, MALE)
	} else {
		Response("PERSON_NOT_FOUND")
	}
}

func GetMaternalAunt(person *Node) {
	parent := person.Parent
	if parent != nil {
		GetSiblingsByGender(parent.Wife, FEMALE)
	} else {
		Response("PERSON_NOT_FOUND")
	}
}

func GetSiterInLaw(person *Node) {
	GetSiblingsOfInLawByGender(person, FEMALE)
}

func GetBrotherInLaw(person *Node) {
	GetSiblingsOfInLawByGender(person, MALE)
}

func GetSiblingsOfInLawByGender(person *Node, gender int) {
	spouse := GetSpouse(person)
	if spouse != nil {
		GetSiblingsByGender(spouse, gender)
	} else {
		Response("PERSON_NOT_FOUND")
	}
}

func GetChildrenByGender(person *Node, gender int) {
	couple := person.MarriageDetails
	children := ""
	if couple != nil && len(couple.Children) > 0 {
		for _, child := range couple.Children {
			if child.Gender == gender {
				children += child.Name + " "
			}
		}

		if len(children) > 0 {
			Response(strings.TrimRight(children, " "))
		} else {
			Response("None")
		}

	} else {
		Response("PERSON_NOT_FOUND")
	}
}

func GetSiblingsByGender(person *Node, gender int) {
	parent := person.Parent
	siblings := ""

	if parent != nil && len(parent.Children) > 0 {
		for _, child := range parent.Children {
			if child.Name != person.Name && (gender == -1 || child.Gender == gender) {
				siblings += child.Name + " "
			}
		}

		if len(siblings) > 0 {
			Response(strings.TrimRight(siblings, " "))
		} else {
			Response("None")
		}

	} else {
		Response("PERSON_NOT_FOUND")
	}
}

func GetSpouse(person *Node) *Node {
	if person.MarriageDetails != nil {
		couple := person.MarriageDetails
		if person.Gender == MALE {
			return couple.Wife
		}
		return couple.Husband
	}

	return nil
}

func GetPerson(name string) (*Node, bool) {
	person, found := Persons[name]
	return person, found
}

func GetCouple(name string) (*Couple, bool) {
	couple, found := Couples[name]
	return couple, found
}

func printFamilyTree(root *Couple) {
	fmt.Println(root.Husband.Name + "   " + root.Wife.Name)
	if len(root.Children) > 0 {
		for _, child := range root.Children {
			fmt.Println(child.Name)
			childCouple, found := Couples[child.Name]
			if found {
				printFamilyTree(childCouple)
			}
		}
	}
}
