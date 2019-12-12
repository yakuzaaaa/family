package main

import "fmt"

func init() {
	// King Shan
	KingShan := createNode("King Shan", MALE, nil)
	QueenAnga := createNode("Queen Anga", FEMALE, nil)

	ShanAnga := KingShan.marriage(&QueenAnga)

	Chit := createNode("Chit", MALE, ShanAnga)
	Ish := createNode("Ish", MALE, ShanAnga)
	Vich := createNode("Vich", MALE, ShanAnga)
	Aras := createNode("Aras", MALE, ShanAnga)
	Satya := createNode("Satya", FEMALE, ShanAnga)

	ShanAnga.addChild(&Chit)
	ShanAnga.addChild(&Ish)
	ShanAnga.addChild(&Vich)
	ShanAnga.addChild(&Aras)
	ShanAnga.addChild(&Satya)

	// Chit Amba
	Amba := createNode("Amba", FEMALE, nil)
	ChitAmba := Chit.marriage(&Amba)

	Dritha := createNode("Dritha", FEMALE, ChitAmba)
	Tritha := createNode("Tritha", FEMALE, ChitAmba)
	Vritha := createNode("Vritha", FEMALE, ChitAmba)

	ChitAmba.addChild(&Dritha)
	ChitAmba.addChild(&Tritha)
	ChitAmba.addChild(&Vritha)

	// Jaya Dritha
	Jaya := createNode("Jaya", MALE, nil)
	JayaDritha := Jaya.marriage(&Dritha)

	Yodhan := createNode("Yodhan", MALE, JayaDritha)
	JayaDritha.addChild(&Yodhan)

	// Vich Lika
	Lika := createNode("Lika", FEMALE, nil)
	VichLika := Vich.marriage(&Lika)

	Vila := createNode("Vila", FEMALE, VichLika)
	Chika := createNode("Chika", FEMALE, VichLika)
	VichLika.addChild(&Vila)
	VichLika.addChild(&Chika)

	// Aras Chitra
	Chitra := createNode("Chitra", FEMALE, nil)
	ArasChitra := Aras.marriage(&Chitra)
	Jinki := createNode("Jinki", FEMALE, ArasChitra)
	Ahit := createNode("Ahit", MALE, ArasChitra)
	ArasChitra.addChild(&Jinki)
	ArasChitra.addChild(&Ahit)

	// Arit Jinki
	Arit := createNode("Arit", MALE, nil)
	AritJinki := Arit.marriage(&Jinki)
	Laki := createNode("Laki", MALE, AritJinki)
	Lavnya := createNode("Lavnya", FEMALE, AritJinki)
	AritJinki.addChild(&Laki)
	AritJinki.addChild(&Lavnya)

	// Satya Viyan
	Vyan := createNode("Vyan", MALE, nil)
	SatyaViyan := Vyan.marriage(&Satya)
	Asva := createNode("Asva", MALE, SatyaViyan)
	Vyas := createNode("Vyas", MALE, SatyaViyan)
	Atya := createNode("Atya", FEMALE, SatyaViyan)
	SatyaViyan.addChild(&Asva)
	SatyaViyan.addChild(&Vyas)
	SatyaViyan.addChild(&Atya)

	// Satvy Asva
	Satvy := createNode("Satvy", FEMALE, nil)
	AsvaSatvy := Asva.marriage(&Satvy)
	Vasa := createNode("Vasa", MALE, AsvaSatvy)
	AsvaSatvy.addChild(&Vasa)

	// Vyas Krpi
	Krpi := createNode("Krpi", FEMALE, nil)
	VyasKrpi := Vyas.marriage(&Krpi)
	Kriya := createNode("Kriya", MALE, VyasKrpi)
	Krithi := createNode("Krithi", FEMALE, VyasKrpi)
	VyasKrpi.addChild(&Kriya)
	VyasKrpi.addChild(&Krithi)

	printFamilyTree(ShanAnga)

}

func createNode(name string, gender int, parent *Couple) Node {
	node := Node{name, gender, parent, nil}

	// TODO: Move indexing logic from here to remove the side effect
	indexNode(node)

	return node
}

func indexNode(node Node) {
	Persons[node.Name] = node
}

func getPerson() {}

func resolveRelation() {}

func resolveRelationMaternal() {}

func resolveRelationPaternal() {}

func printFamilyTree(root *Couple) {
	fmt.Println(root.Husband.Name + "   " + root.Wife.Name)
	fmt.Printf("  DEBUG %d \n", len(root.Children))
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

func main() {
}
