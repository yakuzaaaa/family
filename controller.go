package main

func AddChild(motherName string, childName string, gender int) {
	mother, foundMother := GetPerson(motherName)
	if foundMother {
		parent := mother.MarriageDetails
		if parent != nil {
			child := createNode(childName, gender, parent)
			parent.addChild(child)
			Response("CHILD_ADDITION_SUCCEEDED")
		} else {
			Response("CHILD_ADDITION_FAILED")
		}
	} else {
		Response("PERSON_NOT_FOUND")
	}
}

func ResolveRelation(personName string, relationName string) {
	person, found := GetPerson(personName)
	if found {
		switch relationName {
		case SON:
			GetSon(person)
			break
		case DAUGHTER:
			GetDaughter(person)
			break
		case SIBLINGS:
			GetSiblings(person)
			break
		case PATERNAL_UNCLE:
			GetPaternalUncle(person)
			break
		case PATERNAL_AUNT:
			GetPaternalAunt(person)
			break
		case MATERNAL_UNCLE:
			GetMaternalUncle(person)
			break
		case MATERNAL_AUNT:
			GetMaternalAunt(person)
			break
		case SISTER_IN_LAW:
			GetSiterInLaw(person)
			break
		case BROTHER_IN_LAW:
			GetBrotherInLaw(person)
			break
		default:
			Response("None")
		}
	} else {
		Response("PERSON_NOT_FOUND")
	}
}
