package char

import "github.com/uptrace/bun"

type CharacterStats struct {
	Guid              int     `json:"-"`
	Maxhealth         int     `json:"maxhealth,omitempty"`
	Maxpower1         int     `json:"maxpower1,omitempty"`
	Maxpower2         int     `json:"maxpower2,omitempty"`
	Maxpower3         int     `json:"maxpower3,omitempty"`
	Maxpower4         int     `json:"maxpower4,omitempty"`
	Maxpower5         int     `json:"maxpower5,omitempty"`
	Maxpower6         int     `json:"maxpower6,omitempty"`
	Maxpower7         int     `json:"maxpower7,omitempty"`
	Strength          int     `json:"strength,omitempty"`
	Agility           int     `json:"agility,omitempty"`
	Stamina           int     `json:"stamina,omitempty"`
	Intellect         int     `json:"intellect,omitempty"`
	Spirit            int     `json:"spirit,omitempty"`
	Armor             int     `json:"armor,omitempty"`
	ResHoly           int     `json:"resholy,omitempty"`
	ResFire           int     `json:"resfire,omitempty"`
	ResNature         int     `json:"resnature,omitempty"`
	ResFrost          int     `json:"resfrost,omitempty"`
	ResShadow         int     `json:"resshadow,omitempty"`
	ResArcane         int     `json:"resarcane,omitempty"`
	BlockPct          float64 `json:"blockpct,omitempty"`
	DodgePct          float64 `json:"dodgepct,omitempty"`
	ParryPct          float64 `json:"parrypct,omitempty"`
	CritPct           float64 `json:"critpct,omitempty"`
	RangedCritPct     float64 `json:"rangedcritpct,omitempty"`
	SpellCritPct      float64 `json:"spellcritpct,omitempty"`
	AttackPower       int     `json:"attackpower,omitempty"`
	RangedAttackPower int     `json:"rangedattackpower,omitempty"`
	SpellPower        int     `json:"spellpower,omitempty"`
	Resilience        int     `json:"resilience,omitempty"`
}

type CharacterStatsSlice []CharacterStats

type DBCharacterStats struct {
	bun.BaseModel     `bun:"table:character_stats,alias:character_stats"`
	Guid              int     `bun:"guid"`
	Maxhealth         int     `bun:"maxhealth"`
	Maxpower1         int     `bun:"maxpower1"`
	Maxpower2         int     `bun:"maxpower2"`
	Maxpower3         int     `bun:"maxpower3"`
	Maxpower4         int     `bun:"maxpower4"`
	Maxpower5         int     `bun:"maxpower5"`
	Maxpower6         int     `bun:"maxpower6"`
	Maxpower7         int     `bun:"maxpower7"`
	Strength          int     `bun:"strength"`
	Agility           int     `bun:"agility"`
	Stamina           int     `bun:"stamina"`
	Intellect         int     `bun:"intellect"`
	Spirit            int     `bun:"spirit"`
	Armor             int     `bun:"armor"`
	ResHoly           int     `bun:"resHoly"`
	ResFire           int     `bun:"resFire"`
	ResNature         int     `bun:"resNature"`
	ResFrost          int     `bun:"resFrost"`
	ResShadow         int     `bun:"resShadow"`
	ResArcane         int     `bun:"resArcane"`
	BlockPct          float64 `bun:"blockPct"`
	DodgePct          float64 `bun:"dodgePct"`
	ParryPct          float64 `bun:"parryPct"`
	CritPct           float64 `bun:"critPct"`
	RangedCritPct     float64 `bun:"rangedCritPct"`
	SpellCritPct      float64 `bun:"spellCritPct"`
	AttackPower       int     `bun:"attackPower"`
	RangedAttackPower int     `bun:"rangedAttackPower"`
	SpellPower        int     `bun:"spellPower"`
	Resilience        int     `bun:"resilience"`
}

type DBCharacterStatsSlice []DBCharacterStats

func (entry *CharacterStats) ToDB() *DBCharacterStats {
	if entry == nil {
		return nil
	}
	return &DBCharacterStats{
		Guid:              entry.Guid,
		Maxhealth:         entry.Maxhealth,
		Maxpower1:         entry.Maxpower1,
		Maxpower2:         entry.Maxpower2,
		Maxpower3:         entry.Maxpower3,
		Maxpower4:         entry.Maxpower4,
		Maxpower5:         entry.Maxpower5,
		Maxpower6:         entry.Maxpower6,
		Maxpower7:         entry.Maxpower7,
		Strength:          entry.Strength,
		Agility:           entry.Agility,
		Stamina:           entry.Stamina,
		Intellect:         entry.Intellect,
		Spirit:            entry.Spirit,
		Armor:             entry.Armor,
		ResHoly:           entry.ResHoly,
		ResFire:           entry.ResFire,
		ResNature:         entry.ResNature,
		ResFrost:          entry.ResFrost,
		ResShadow:         entry.ResShadow,
		ResArcane:         entry.ResArcane,
		BlockPct:          entry.BlockPct,
		DodgePct:          entry.DodgePct,
		ParryPct:          entry.ParryPct,
		CritPct:           entry.CritPct,
		RangedCritPct:     entry.RangedCritPct,
		SpellCritPct:      entry.SpellCritPct,
		AttackPower:       entry.AttackPower,
		RangedAttackPower: entry.RangedAttackPower,
		SpellPower:        entry.SpellPower,
		Resilience:        entry.Resilience,
	}
}

func (entry *DBCharacterStats) ToWeb() *CharacterStats {
	if entry == nil {
		return nil
	}
	return &CharacterStats{
		Guid:              entry.Guid,
		Maxhealth:         entry.Maxhealth,
		Maxpower1:         entry.Maxpower1,
		Maxpower2:         entry.Maxpower2,
		Maxpower3:         entry.Maxpower3,
		Maxpower4:         entry.Maxpower4,
		Maxpower5:         entry.Maxpower5,
		Maxpower6:         entry.Maxpower6,
		Maxpower7:         entry.Maxpower7,
		Strength:          entry.Strength,
		Agility:           entry.Agility,
		Stamina:           entry.Stamina,
		Intellect:         entry.Intellect,
		Spirit:            entry.Spirit,
		Armor:             entry.Armor,
		ResHoly:           entry.ResHoly,
		ResFire:           entry.ResFire,
		ResNature:         entry.ResNature,
		ResFrost:          entry.ResFrost,
		ResShadow:         entry.ResShadow,
		ResArcane:         entry.ResArcane,
		BlockPct:          entry.BlockPct,
		DodgePct:          entry.DodgePct,
		ParryPct:          entry.ParryPct,
		CritPct:           entry.CritPct,
		RangedCritPct:     entry.RangedCritPct,
		SpellCritPct:      entry.SpellCritPct,
		AttackPower:       entry.AttackPower,
		RangedAttackPower: entry.RangedAttackPower,
		SpellPower:        entry.SpellPower,
		Resilience:        entry.Resilience,
	}
}

func (data CharacterStatsSlice) ToDB() DBCharacterStatsSlice {
	result := make(DBCharacterStatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterStatsSlice) ToWeb() CharacterStatsSlice {
	result := make(CharacterStatsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
