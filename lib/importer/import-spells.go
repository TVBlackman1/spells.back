package importer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// structure from http://tentaculus.ru/spells/

func GetSpellsData(r io.Reader) *mainStructure {
	res := new(mainStructure)

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	dataInBytes := buf.Bytes()
	dataInBytes = bytes.TrimPrefix(dataInBytes, []byte("\xef\xbb\xbf"))
	err := json.Unmarshal(dataInBytes, res)
	if err != nil {
		fmt.Println(err)
	}
	//decoder := json.NewDecoder(r)
	//err := decoder.Decode(&res)
	//if err != nil {
	//	fmt.Println(err)
	//}
	return res
}

type mainStructure struct {
	SourceList map[string]source    `json:"sourceList"`
	SchoolList map[string]school    `json:"schoolList"`
	OLanguages map[string]languages `json:"oLanguages"`
	AllSpells  []struct {
		En spell `json:"en"`
		Ru spell `json:"ru"`
	} `json:"allSpells"`
	//LockedItems interface{} `json:"lockedItems"`
}

type source struct {
	Text     titledText `json:"text"`
	Official bool       `json:"official"`
	Checked  bool       `json:"checked"`
	Visible  bool       `json:"visible"`
}

type school struct {
	Text    titledText `json:"text"`
	Checked bool       `json:"checked"`
	Visible bool       `json:"visible"`
	I       int        `json:"i"`
}

type languages struct {
	Text titledText `json:"text"`
}

type spell struct {
	Name        string      `json:"name"`
	Level       json.Number `json:"level"`
	Text        string      `json:"text"`
	School      string      `json:"school"`
	CastingTime string      `json:"castingTime"`
	Range       string      `json:"range"`
	Materials   string      `json:"materials"`
	Components  string      `json:"components"`
	Duration    string      `json:"duration"`
	Source      string      `json:"source"`
	Ritual      string      `json:"ritual"`
}

type titledText struct {
	En struct {
		Title string
	} `json:"en"`
	Ru struct {
		Title string
	} `json:"ru"`
}
