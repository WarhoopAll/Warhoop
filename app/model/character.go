package model

import "github.com/uptrace/bun"

type Characters struct {
	Guid                 int    `json:"guid,omitempty"`
	Name                 string `json:"name,omitempty"`
	Race                 int    `json:"race,omitempty"`
	Class                int    `json:"class,omitempty"`
	Gender               int    `json:"gender"`
	Level                int    `json:"level,omitempty"`
	Xp                   int    `json:"xp,omitempty"`
	Money                int    `json:"money,omitempty"`
	Skin                 int    `json:"skin,omitempty"`
	Face                 int    `json:"face,omitempty"`
	HairStyle            int    `json:"hairstyle,omitempty"`
	HairColor            int    `json:"haircolor,omitempty"`
	FacialStyle          int    `json:"facialstyle,omitempty"`
	Map                  *Map   `json:"map,omitempty"`
	Online               int    `json:"online,omitempty"`
	Totaltime            int    `json:"totaltime,omitempty"`
	Leveltime            int    `json:"leveltime,omitempty"`
	LogoutTime           int    `json:"logouttime,omitempty"`
	Zone                 *Zone  `json:"zone,omitempty"`
	ArenaPoints          int    `json:"arenapoints,omitempty"`
	TotalHonorPoints     int    `json:"totalhonorpoints,omitempty"`
	TodayHonorPoints     int    `json:"todayhonorpoints,omitempty"`
	YesterdayHonorPoints int    `json:"yesterdayhonorpoints,omitempty"`
	TotalKills           int    `json:"totalkills,omitempty"`
	TodayKills           int    `json:"todaykills,omitempty"`
	YesterdayKills       int    `json:"yesterdaykills,omitempty"`
	ChosenTitle          int    `json:"chosentitle,omitempty"`
	Health               int    `json:"health,omitempty"`
	Power1               int    `json:"power1,omitempty"`
	Power2               int    `json:"power2,omitempty"`
	Power3               int    `json:"power3,omitempty"`
	Power4               int    `json:"power4,omitempty"`
	Power5               int    `json:"power5,omitempty"`
	Power6               int    `json:"power6,omitempty"`
	Power7               int    `json:"power7,omitempty"`
	TalentGroupsCount    int    `json:"talentgroupscount,omitempty"`
	ActiveTalentGroup    int    `json:"activetalentgroup,omitempty"`
	EquipmentCache       string `json:"equipmentcache,omitempty"`
	KnownTitles          string `json:"knowntitles,omitempty"`
}

type CharactersSlice []Characters

type DBCharacters struct {
	bun.BaseModel        `bun:"table:characters,alias:characters"`
	Guid                 int     `bun:"guid,pk"`
	Name                 string  `bun:"name"`
	Race                 int     `bun:"race"`
	Class                int     `bun:"class"`
	Gender               int     `bun:"gender"`
	Level                int     `bun:"level"`
	Xp                   int     `bun:"xp"`
	Money                int     `bun:"money"`
	Skin                 int     `bun:"skin"`
	Face                 int     `bun:"face"`
	HairStyle            int     `bun:"hairStyle"`
	HairColor            int     `bun:"hairColor"`
	FacialStyle          int     `bun:"facialStyle"`
	Map                  int     `bun:"map,pk"`
	Maps                 *DBMap  `bun:"rel:belongs-to,join:map=ID"`
	Online               int     `bun:"online"`
	Totaltime            int     `bun:"totaltime"`
	Leveltime            int     `bun:"leveltime"`
	LogoutTime           int     `bun:"logout_time"`
	Zone                 int     `bun:"zone"`
	Zones                *DBZone `bun:"rel:belongs-to,join:zone=ID"`
	ArenaPoints          int     `bun:"arenaPoints"`
	TotalHonorPoints     int     `bun:"totalHonorPoints"`
	TodayHonorPoints     int     `bun:"todayHonorPoints"`
	YesterdayHonorPoints int     `bun:"yesterdayHonorPoints"`
	TotalKills           int     `bun:"totalKills"`
	TodayKills           int     `bun:"todayKills"`
	YesterdayKills       int     `bun:"yesterdayKills"`
	ChosenTitle          int     `bun:"chosenTitle"`
	Health               int     `bun:"health"`
	Power1               int     `bun:"power1"`
	Power2               int     `bun:"power2"`
	Power3               int     `bun:"power3"`
	Power4               int     `bun:"power4"`
	Power5               int     `bun:"power5"`
	Power6               int     `bun:"power6"`
	Power7               int     `bun:"power7"`
	TalentGroupsCount    int     `bun:"talentGroupsCount"`
	ActiveTalentGroup    int     `bun:"activeTalentGroup"`
	EquipmentCache       string  `bun:"equipmentCache"`
	KnownTitles          string  `bun:"knownTitles"`
}

type DBCharactersSlice []DBCharacters

func (entry *Characters) ToDB() *DBCharacters {
	if entry == nil {
		return nil
	}
	return &DBCharacters{
		Guid:                 entry.Guid,
		Name:                 entry.Name,
		Race:                 entry.Race,
		Class:                entry.Class,
		Gender:               entry.Gender,
		Level:                entry.Level,
		Xp:                   entry.Xp,
		Money:                entry.Money,
		Skin:                 entry.Skin,
		Face:                 entry.Face,
		HairStyle:            entry.HairStyle,
		HairColor:            entry.HairColor,
		FacialStyle:          entry.FacialStyle,
		Online:               entry.Online,
		Totaltime:            entry.Totaltime,
		Leveltime:            entry.Leveltime,
		LogoutTime:           entry.LogoutTime,
		ArenaPoints:          entry.ArenaPoints,
		TotalHonorPoints:     entry.TotalHonorPoints,
		TodayHonorPoints:     entry.TodayHonorPoints,
		YesterdayHonorPoints: entry.YesterdayHonorPoints,
		TotalKills:           entry.TotalKills,
		TodayKills:           entry.TodayKills,
		YesterdayKills:       entry.YesterdayKills,
		ChosenTitle:          entry.ChosenTitle,
		Health:               entry.Health,
		Power1:               entry.Power1,
		Power2:               entry.Power2,
		Power3:               entry.Power3,
		Power4:               entry.Power4,
		Power5:               entry.Power5,
		Power6:               entry.Power6,
		Power7:               entry.Power7,
		TalentGroupsCount:    entry.TalentGroupsCount,
		ActiveTalentGroup:    entry.ActiveTalentGroup,
		EquipmentCache:       entry.EquipmentCache,
		KnownTitles:          entry.KnownTitles,
	}
}

func (entry *DBCharacters) ToWeb() *Characters {
	if entry == nil {
		return nil
	}
	return &Characters{
		Guid:                 entry.Guid,
		Name:                 entry.Name,
		Race:                 entry.Race,
		Class:                entry.Class,
		Gender:               entry.Gender,
		Level:                entry.Level,
		Xp:                   entry.Xp,
		Money:                entry.Money,
		Skin:                 entry.Skin,
		Face:                 entry.Face,
		HairStyle:            entry.HairStyle,
		HairColor:            entry.HairColor,
		FacialStyle:          entry.FacialStyle,
		Online:               entry.Online,
		Totaltime:            entry.Totaltime,
		Leveltime:            entry.Leveltime,
		LogoutTime:           entry.LogoutTime,
		ArenaPoints:          entry.ArenaPoints,
		TotalHonorPoints:     entry.TotalHonorPoints,
		TodayHonorPoints:     entry.TodayHonorPoints,
		YesterdayHonorPoints: entry.YesterdayHonorPoints,
		TotalKills:           entry.TotalKills,
		TodayKills:           entry.TodayKills,
		YesterdayKills:       entry.YesterdayKills,
		ChosenTitle:          entry.ChosenTitle,
		Health:               entry.Health,
		Power1:               entry.Power1,
		Power2:               entry.Power2,
		Power3:               entry.Power3,
		Power4:               entry.Power4,
		Power5:               entry.Power5,
		Power6:               entry.Power6,
		Power7:               entry.Power7,
		TalentGroupsCount:    entry.TalentGroupsCount,
		ActiveTalentGroup:    entry.ActiveTalentGroup,
		EquipmentCache:       entry.EquipmentCache,
		KnownTitles:          entry.KnownTitles,
		Map:                  entry.Maps.ToWeb(),
		Zone:                 entry.Zones.ToWeb(),
	}
}

func (data CharactersSlice) ToDB() *DBCharactersSlice {
	result := make(DBCharactersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return &result
}

func (data DBCharactersSlice) ToWeb() *CharactersSlice {
	result := make(CharactersSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return &result
}
