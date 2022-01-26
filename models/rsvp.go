package models

import (
	"github.com/kamva/mgm/v3"
)


type RSVP struct {
	mgm.DefaultModel `bson:",inline"`
        Name                string `json:"name" bson:"name"`
	DietaryRestrictions string `json:"diet" bson:"diet"`
	Attending           bool `json:"attending" bson:"attending"`
}

func NewRSVP(name string, diet string, attending bool) *RSVP {
	return &RSVP{
		Name: name,
		DietaryRestrictions: diet,
		Attending: attending,
	}
}
