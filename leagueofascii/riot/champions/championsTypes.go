package champions

type ChampionSummaryStats struct {
	Hp                   float64
	Hpperlevel           float64
	Mp                   float64
	Mpperlevel           float64
	Movespeed            float64
	Armor                float64
	Amorperlevel         float64
	Spellblock           float64
	Spellblockperlevel   float64
	Attackrange          float64
	Hpregen              float64
	Hpregenperlevel      float64
	Mpregen              float64
	Mpregenperlevel      float64
	Crit                 float64
	Critperlevel         float64
	Attackdamage         float64
	Attackdamageperlevel float64
	Attackspeed          float64
	Attackspeedperlevel  float64
}

type ChampionSummaryImage struct {
	Full   string
	Sprite string
	X      float64
	Y      float64
	W      float64
	H      float64
}

type ChampionSummaryInfo struct {
	Attack     float64
	Defense    float64
	Magic      float64
	Difficulty float64
}

type ChampionSummary struct {
	Version  string
	Id       string
	Key      string
	Name     string
	Title    string
	Blurb    string
	Info     ChampionSummaryInfo
	Image    ChampionSummaryImage
	Tags     []string
	PartType string
	Stats    ChampionSummaryStats
}

type ChampionsSummary struct {
	Type    string
	Format  string
	Version string
	Data    map[string]ChampionSummary
}

type ChampionImage struct {
	Full   string
	Sprite string
	X      float64
	Y      float64
	W      float64
	H      float64
	Group  string
}

type ChampionResponse struct {
	Type    string
	Format  string
	Version string
	Data    map[string]Champion
}

type Champion struct {
	Id    string
	Key   string
	Name  string
	Title string
	Image ChampionImage
}
