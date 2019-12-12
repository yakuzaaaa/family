package main

const (
	MALE   = 0
	FEMALE = 1
)

type Couple struct {
	Husband  *Node
	Wife     *Node
	Children []*Node
}

type Node struct {
	Name            string
	Gender          int
	Parent          *Couple
	MarriageDetails *Couple
}

func (husband *Node) marriage(wife *Node) *Couple {
	couple := Couple{husband, wife, nil}

	husband.MarriageDetails = &couple
	wife.MarriageDetails = &couple

	Couples[couple.Husband.Name] = &couple
	Couples[couple.Wife.Name] = &couple

	return &couple
}

func (couple *Couple) addChild(child *Node) {
	couple.Children = append(couple.Children, child)
}

var Persons = make(map[string]Node)
var Couples = make(map[string]*Couple)
