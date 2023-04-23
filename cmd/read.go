package main

import (
	"fmt"
	"strings"
)

func main() {
	//file, err := os.Open("./init/seeds/dnd-spells.json")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//data := importer.GetSpellsData(file)
	//fmt.Println(len(data.AllSpells))
	//for _, spell := range data.AllSpells[:20] {
	//	println(spell.Ru.Name)
	//	ruSpell := spell.Ru
	//	isVerbal := slices.Contains(strings.Split(ruSpell.Components, ", "), "В")
	//	isSomatic := slices.Contains(strings.Split(ruSpell.Components, ", "), "С")
	//	hasMaterial := slices.Contains(strings.Split(ruSpell.Components, ", "), "М")
	//	println(isVerbal, isSomatic, hasMaterial)
	//}
	fmt.Println(strings.Replace("kasjdkas' as' asd", "'", "", -1))
}
