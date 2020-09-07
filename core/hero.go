package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/husk/validation"
	"time"
)

type Hero struct {
	EntityKey   hsk.Key
	Credits     int
	Experiences []Experience
	Level       Level
	TotalXP     int
	LastUpdated time.Time //update on save???
}

func (o Hero) Valid() error {
	return validation.Struct(o)
}

func GetHeroes(page, pagesize int) (records.Page, error) {
	return ctx.Heroes.Find(page, pagesize, op.Everything())
}

func GetHero(key hsk.Key) (hsk.Record, error) {
	return ctx.Heroes.FindByKey(key)
}

func (h *Hero) AddExperience(xpType ExperienceType) {
	// Add XP Record with (new) Level
	// Update Hero to reflect new TotalXP
	xpValue := XPValue(xpType)
	xp := Experience{
		Points: xpValue,
		Type:   xpType,
	}

	h.Experiences = append(h.Experiences, xp)
	h.TotalXP += xpValue
}

func (h *Hero) AddRequisition() {

}

func (h *Hero) AddCredit(amount int) {

}
