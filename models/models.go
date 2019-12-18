package models

// Represents app level constants
const (
	MALE          = 0
	FEMALE        = 1
	SON           = "Son"
	DAUGHTER      = "Daughter"
	SIBLINGS      = "Siblings"
	PATERNALUNCLE = "Paternal-Uncle"
	MATERNALUNCLE = "Maternal-Uncle"
	PATERNALAUNT  = "Paternal-Aunt"
	MATERNALAUNT  = "Maternal-Aunt"
	SISTERINLAW   = "Sister-In-Law"
	BROTHERINLAW  = "Borther-In-Law"
)

// Couple represents married structure 
// of husband and wife
type Couple struct {
	Husband  *Node
	Wife     *Node
	Children []*Node
}

// Node represents a person
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

// AddChild adds relation of child parent to a couple
func (couple *Couple) AddChild(child *Node) {
	couple.Children = append(couple.Children, child)
}

// Persons is an app level index of persons
var Persons = make(map[string]*Node)
// Couples is an app level index of couples
var Couples = make(map[string]*Couple)

// CreateNode creates a person
func CreateNode(name string, gender int, parent *Couple) *Node {
	node := Node{name, gender, parent, nil}

	// TODO: Move indexing logic from here to remove the side effect
	indexNode(&node)

	return &node
}

func indexNode(node *Node) {
	Persons[node.Name] = node
}

// StaticTreeGenerator generates the hard-coded
// family tree given in the problem
func StaticTreeGenerator() *Couple {
	// King Shan
	KingShan := CreateNode("King Shan", MALE, nil)
	QueenAnga := CreateNode("Queen Anga", FEMALE, nil)

	ShanAnga := KingShan.marriage(QueenAnga)

	Chit := CreateNode("Chit", MALE, ShanAnga)
	Ish := CreateNode("Ish", MALE, ShanAnga)
	Vich := CreateNode("Vich", MALE, ShanAnga)
	Aras := CreateNode("Aras", MALE, ShanAnga)
	Satya := CreateNode("Satya", FEMALE, ShanAnga)

	ShanAnga.AddChild(Chit)
	ShanAnga.AddChild(Ish)
	ShanAnga.AddChild(Vich)
	ShanAnga.AddChild(Aras)
	ShanAnga.AddChild(Satya)

	// Chit Amba
	Amba := CreateNode("Amba", FEMALE, nil)
	ChitAmba := Chit.marriage(Amba)

	Dritha := CreateNode("Dritha", FEMALE, ChitAmba)
	Tritha := CreateNode("Tritha", FEMALE, ChitAmba)
	Vritha := CreateNode("Vritha", FEMALE, ChitAmba)

	ChitAmba.AddChild(Dritha)
	ChitAmba.AddChild(Tritha)
	ChitAmba.AddChild(Vritha)

	// Jaya Dritha
	Jaya := CreateNode("Jaya", MALE, nil)
	JayaDritha := Jaya.marriage(Dritha)

	Yodhan := CreateNode("Yodhan", MALE, JayaDritha)
	JayaDritha.AddChild(Yodhan)

	// Vich Lika
	Lika := CreateNode("Lika", FEMALE, nil)
	VichLika := Vich.marriage(Lika)

	Vila := CreateNode("Vila", FEMALE, VichLika)
	Chika := CreateNode("Chika", FEMALE, VichLika)
	VichLika.AddChild(Vila)
	VichLika.AddChild(Chika)

	// Aras Chitra
	Chitra := CreateNode("Chitra", FEMALE, nil)
	ArasChitra := Aras.marriage(Chitra)
	Jinki := CreateNode("Jinki", FEMALE, ArasChitra)
	Ahit := CreateNode("Ahit", MALE, ArasChitra)
	ArasChitra.AddChild(Jinki)
	ArasChitra.AddChild(Ahit)

	// Arit Jinki
	Arit := CreateNode("Arit", MALE, nil)
	AritJinki := Arit.marriage(Jinki)
	Laki := CreateNode("Laki", MALE, AritJinki)
	Lavnya := CreateNode("Lavnya", FEMALE, AritJinki)
	AritJinki.AddChild(Laki)
	AritJinki.AddChild(Lavnya)

	// Satya Viyan
	Vyan := CreateNode("Vyan", MALE, nil)
	SatyaViyan := Vyan.marriage(Satya)
	Asva := CreateNode("Asva", MALE, SatyaViyan)
	Vyas := CreateNode("Vyas", MALE, SatyaViyan)
	Atya := CreateNode("Atya", FEMALE, SatyaViyan)
	SatyaViyan.AddChild(Asva)
	SatyaViyan.AddChild(Vyas)
	SatyaViyan.AddChild(Atya)

	// Satvy Asva
	Satvy := CreateNode("Satvy", FEMALE, nil)
	AsvaSatvy := Asva.marriage(Satvy)
	Vasa := CreateNode("Vasa", MALE, AsvaSatvy)
	AsvaSatvy.AddChild(Vasa)

	// Vyas Krpi
	Krpi := CreateNode("Krpi", FEMALE, nil)
	VyasKrpi := Vyas.marriage(Krpi)
	Kriya := CreateNode("Kriya", MALE, VyasKrpi)
	Krithi := CreateNode("Krithi", FEMALE, VyasKrpi)
	VyasKrpi.AddChild(Kriya)
	VyasKrpi.AddChild(Krithi)

	return ShanAnga
}
