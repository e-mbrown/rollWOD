// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Character interface {
	IsCharacter()
	GetID() *string
	GetName() string
	GetPlayer() *User
	GetChronicle() []*Campaign
	GetNature() *string
	GetDemeanor() *string
}

type Entry interface {
	IsEntry()
	GetName() string
	GetDescription() string
}

type Age struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (Age) IsEntry()                    {}
func (this Age) GetName() string        { return this.Name }
func (this Age) GetDescription() string { return this.Description }

type Campaign struct {
	ID           string      `json:"id"`
	User         []*User     `json:"user,omitempty"`
	Restrictions []*string   `json:"restrictions,omitempty"`
	Characters   []Character `json:"characters,omitempty"`
}

type Characteristic struct {
	Name        string   `json:"name"`
	Type        CharType `json:"type"`
	Description string   `json:"description"`
	DescbyVal   []string `json:"DescbyVal"`
}

func (Characteristic) IsEntry()                    {}
func (this Characteristic) GetName() string        { return this.Name }
func (this Characteristic) GetDescription() string { return this.Description }

type Clan struct {
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	AssociatedSect []*Sect           `json:"associatedSect"`
	Haven          string            `json:"haven"`
	Background     string            `json:"background"`
	Character      string            `json:"character"`
	Discipline     []*Characteristic `json:"discipline,omitempty"`
	Weakness       string            `json:"weakness"`
	Organizations  *string           `json:"organizations,omitempty"`
	SubClan        []*Clan           `json:"subClan,omitempty"`
	Strongholds    []string          `json:"strongholds,omitempty"`
	IsHighClan     *bool             `json:"isHighClan,omitempty"`
	IsSubclan      *bool             `json:"isSubclan,omitempty"`
}

func (Clan) IsEntry()                    {}
func (this Clan) GetName() string        { return this.Name }
func (this Clan) GetDescription() string { return this.Description }

type Generation struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (Generation) IsEntry()                    {}
func (this Generation) GetName() string        { return this.Name }
func (this Generation) GetDescription() string { return this.Description }

type NewCharacteristic struct {
	Name        string   `json:"name"`
	Type        CharType `json:"type"`
	Description string   `json:"description"`
	ValDesc     []string `json:"ValDesc"`
}

type Sect struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Titles      []Entry  `json:"titles,omitempty"`
	Practices   []string `json:"practices,omitempty"`
	Rituals     []string `json:"rituals,omitempty"`
	Strongholds []string `json:"strongholds,omitempty"`
}

func (Sect) IsEntry()                    {}
func (this Sect) GetName() string        { return this.Name }
func (this Sect) GetDescription() string { return this.Description }

type Title struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (Title) IsEntry()                    {}
func (this Title) GetName() string        { return this.Name }
func (this Title) GetDescription() string { return this.Description }

type Tradition struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Traditions  []Entry `json:"traditions,omitempty"`
}

func (Tradition) IsEntry()                    {}
func (this Tradition) GetName() string        { return this.Name }
func (this Tradition) GetDescription() string { return this.Description }

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Vampire struct {
	ID         *string           `json:"id,omitempty"`
	Name       string            `json:"name"`
	Player     *User             `json:"player"`
	Chronicle  []*Campaign       `json:"chronicle"`
	Nature     *string           `json:"nature,omitempty"`
	Demeanor   *string           `json:"demeanor,omitempty"`
	Concept    *string           `json:"concept,omitempty"`
	Clan       *Clan             `json:"clan"`
	Generation *Generation       `json:"generation"`
	Sire       *Vampire          `json:"Sire,omitempty"`
	Attributes []*Characteristic `json:"attributes"`
	Abilities  []*Characteristic `json:"abilities"`
	Advantages []*Characteristic `json:"advantages"`
}

func (Vampire) IsCharacter()          {}
func (this Vampire) GetID() *string   { return this.ID }
func (this Vampire) GetName() string  { return this.Name }
func (this Vampire) GetPlayer() *User { return this.Player }
func (this Vampire) GetChronicle() []*Campaign {
	if this.Chronicle == nil {
		return nil
	}
	interfaceSlice := make([]*Campaign, 0, len(this.Chronicle))
	for _, concrete := range this.Chronicle {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}
func (this Vampire) GetNature() *string   { return this.Nature }
func (this Vampire) GetDemeanor() *string { return this.Demeanor }

type CharType string

const (
	CharTypeAttribute  CharType = "attribute"
	CharTypeTalent     CharType = "talent"
	CharTypeDiscipline CharType = "discipline"
	CharTypeSkill      CharType = "skill"
	CharTypeBackground CharType = "background"
)

var AllCharType = []CharType{
	CharTypeAttribute,
	CharTypeTalent,
	CharTypeDiscipline,
	CharTypeSkill,
	CharTypeBackground,
}

func (e CharType) IsValid() bool {
	switch e {
	case CharTypeAttribute, CharTypeTalent, CharTypeDiscipline, CharTypeSkill, CharTypeBackground:
		return true
	}
	return false
}

func (e CharType) String() string {
	return string(e)
}

func (e *CharType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CharType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid charType", str)
	}
	return nil
}

func (e CharType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
