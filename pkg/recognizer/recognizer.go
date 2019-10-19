package recognizer

import (
	"encoding/json"
	"fmt"
	"github.com/RadikChernyshov/klingon/pkg/recognizer/rest"
	"strings"
)

type Recognize struct {
	In  string
	Out string
}

type Characters struct {
	Characters []Character
}

type Character struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

type CharacterSpecies struct {
	Character struct {
		UID              string
		Name             string
		CharacterSpecies []struct {
			UID  string
			Name string
		}
		CharacterRelations []interface{}
	}
}

func getCharacter(name string) (c Character, err error) {
	var r Characters
	p := strings.NewReader(fmt.Sprintf("name=%s", name))
	res, err := rest.PostReq("/character/search", p)
	if err != nil {
		return c, err
	}
	_ = json.Unmarshal(res, &r)
	if len(r.Characters) == 0 || !strings.Contains(strings.ToLower(r.Characters[0].Name), strings.ToLower(name)) {
		err = fmt.Errorf("character `%s` not found", name)
	} else {
		c = r.Characters[0]
	}
	return c, err
}

func getCharacterSpecie(c Character) (s string, err error) {
	res, err := rest.GetReq(fmt.Sprintf("/character?uid=%s", c.UID))
	if err != nil {
		return s, err
	}
	var r CharacterSpecies
	_ = json.Unmarshal(res, &r)
	if len(r.Character.CharacterSpecies) == 0 {
		err = fmt.Errorf("characters `%s` species not found", c.Name)
	} else {
		s = r.Character.CharacterSpecies[0].Name
	}
	return s, err
}

func (r *Recognize) Recognize() (err error) {
	c, err := getCharacter(r.In)
	if err != nil {
		return err
	}
	specie, err := getCharacterSpecie(c)
	if err != nil {
		return err
	}
	r.Out = strings.Title(strings.ToLower(specie))
	return err
}

func New() *Recognize {
	return &Recognize{
		In:  "",
		Out: "",
	}
}
