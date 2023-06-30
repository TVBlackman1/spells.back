package importer_v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

func GetSpellsData(r io.Reader) *MainStructure {
	res := new(MainStructure)

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	dataInBytes := buf.Bytes()
	dataInBytes = bytes.TrimPrefix(dataInBytes, []byte("\xef\xbb\xbf"))
	err := json.Unmarshal(dataInBytes, res)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

type MainStructure struct {
	Name       string  `json:"name"`
	ShortName  string  `json:"shortName"`
	IsOfficial bool    `json:"isOfficial"`
	Spells     []spell `json:"spells"`
}

type spell struct {
	Name                 string `json:"name"`
	Level                int    `json:"level"`
	Description          string `json:"description"`
	School               string `json:"magicalSchool"`
	CastingTime          string `json:"castingTime"`
	Distance             string `json:"distance"`
	Materials            string `json:"materialContent"`
	IsVerbal             bool   `json:"isVerbal"`
	IsSomatic            bool   `json:"isSomatic"`
	HasMaterialComponent bool   `json:"hasMaterialComponent"`
	Duration             string `json:"duration"`
	IsRitual             bool   `json:"isRitual"`
}
