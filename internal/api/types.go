package api

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	Me              APIKeyDetails          `json:"me"`
	Treasures       []Treasure             `json:"treasures"`
	Colors          []Color                `json:"colors"`
	GameInfo        GameInfo               `json:"game_info"`
	Nations         NationPaginator        `json:"nations"`
	Alliances       AlliancePaginator      `json:"alliances"`
	TradePrices     TradePricePaginator    `json:"tradeprices"`
	Trades          TradePaginator         `json:"trades"`
	Wars            WarPaginator           `json:"wars"`
	Bounties        BountyPaginator        `json:"bounties"`
	Warattacks      WarAttackPaginator     `json:"warattacks"`
	Treaties        TreatyPaginator        `json:"treaties"`
	Cities          CityPaginator          `json:"cities"`
	BankRecs        BankRecPaginator       `json:"bankrecs"`
	BaseballGames   BBGamePaginator        `json:"baseball_games"`
	BaseballTeams   BBTeamPaginator        `json:"baseball_teams"`
	BaseballPlayers BBPlayerPaginator      `json:"baseball_players"`
	TreasureTrades  TreasureTradePaginator `json:"treasure_trades"`
	Embargoes       EmbargoPaginator       `json:"embargoes"`

	Errors []Error `json:"errors"`
}

type APIKeyDetails struct {
	Nation      Nation `json:"nation"`
	Key         string `json:"key"`
	Requests    int    `json:"requests"`
	MaxRequests int    `json:"max_requests"`
}

type Nation struct {
	ID                               ID                   `json:"id"`
	AllianceID                       ID                   `json:"alliance_id"`
	AlliancePosition                 AlliancePositionEnum `json:"alliance_position"`
	AlliancePositionID               ID                   `json:"alliance_position_id"`
	AlliancePositionInfo             AlliancePosition     `json:"alliance_position_info"`
	Alliance                         Alliance             `json:"alliance"`
	NationName                       string               `json:"nation_name"`
	LeaderName                       string               `json:"leader_name"`
	Continent                        string               `json:"continent"`
	WarPolicy                        WarPolicy            `json:"war_policy"`
	DomesticPolicy                   DomesticPolicy       `json:"domestic_policy"`
	Color                            string               `json:"color"`
	NumCities                        int                  `json:"num_cities"`
	Cities                           []City               `json:"cities"`
	Score                            float64              `json:"score"`
	UpdateTz                         float64              `json:"update_tz"`
	Population                       int                  `json:"population"`
	Flag                             string               `json:"flag"`
	VacationModeTurns                int                  `json:"vacation_mode_turns"`
	BeigeTurns                       int                  `json:"beige_turns"`
	EspionageAvailable               bool                 `json:"espionage_available"`
	LastActive                       DateTimeAuto         `json:"last_active"`
	Date                             DateTimeAuto         `json:"date"`
	Soldiers                         int                  `json:"soldiers"`
	Tanks                            int                  `json:"tanks"`
	Aircraft                         int                  `json:"aircraft"`
	Ships                            int                  `json:"ships"`
	Missiles                         int                  `json:"missiles"`
	Nukes                            int                  `json:"nukes"`
	Spies                            int                  `json:"spies"`
	Discord                          string               `json:"discord"`
	Treasures                        []Treasure           `json:"treasures"`
	Wars                             []War                `json:"wars"`
	BankRecs                         []BankRec            `json:"bankrecs"`
	Taxrecs                          []BankRec            `json:"taxrecs"`
	Bounties                         []Bounty             `json:"bounties"`
	TurnsSinceLastCity               int                  `json:"turns_since_last_city"`
	TurnsSinceLastProject            int                  `json:"turns_since_last_project"`
	Money                            float64              `json:"money"`
	Coal                             float64              `json:"coal"`
	Oil                              float64              `json:"oil"`
	Uranium                          float64              `json:"uranium"`
	Iron                             float64              `json:"iron"`
	Bauxite                          float64              `json:"bauxite"`
	Lead                             float64              `json:"lead"`
	Gasoline                         float64              `json:"gasoline"`
	Munitions                        float64              `json:"munitions"`
	Steel                            float64              `json:"steel"`
	Aluminum                         float64              `json:"aluminum"`
	Food                             float64              `json:"food"`
	Projects                         int                  `json:"projects"`
	ProjectBits                      int                  `json:"project_bits"`
	IronWorks                        bool                 `json:"iron_works"`
	BauxiteWorks                     bool                 `json:"bauxite_works"`
	ArmsStockpile                    bool                 `json:"arms_stockpile"`
	EmergencyGasolineReserve         bool                 `json:"emergency_gasoline_reserve"`
	MassIrrigation                   bool                 `json:"mass_irrigation"`
	InternationalTradeCenter         bool                 `json:"international_trade_center"`
	MissileLaunchPad                 bool                 `json:"missile_launch_pad"`
	NuclearResearchFacility          bool                 `json:"nuclear_research_facility"`
	IronDome                         bool                 `json:"iron_dome"`
	VitalDefenseSystem               bool                 `json:"vital_defense_system"`
	CentralIntelligenceAgency        bool                 `json:"central_intelligence_agency"`
	CenterForCivilEngineering        bool                 `json:"center_for_civil_engineering"`
	PropagandaBureau                 bool                 `json:"propaganda_bureau"`
	UraniumEnrichmentProgram         bool                 `json:"uranium_enrichment_program"`
	UrbanPlanning                    bool                 `json:"urban_planning"`
	AdvancedUrbanPlanning            bool                 `json:"advanced_urban_planning"`
	SpaceProgram                     bool                 `json:"space_program"`
	SpySatellite                     bool                 `json:"spy_satellite"`
	MoonLanding                      bool                 `json:"moon_landing"`
	PirateEconomy                    bool                 `json:"pirate_economy"`
	RecyclingInitiative              bool                 `json:"recycling_initiative"`
	TelecommunicationsSatellite      bool                 `json:"telecommunications_satellite"`
	GreenTechnologies                bool                 `json:"green_technologies"`
	ArableLandAgency                 bool                 `json:"arable_land_agency"`
	ClinicalResearchCenter           bool                 `json:"clinical_research_center"`
	SpecializedPoliceTrainingProgram bool                 `json:"specialized_police_training_program"`
	AdvancedEngineeringCorps         bool                 `json:"advanced_engineering_corps"`
	GovernmentSupportAgency          bool                 `json:"government_support_agency"`
	ResearchAndDevelopmentCenter     bool                 `json:"research_and_development_center"`
	ResourceProductionCenter         bool                 `json:"resource_production_center"`
	WarsWon                          int                  `json:"wars_won"`
	WarsLost                         int                  `json:"wars_lost"`
	TaxID                            ID                   `json:"tax_id"`
	AllianceSeniority                int                  `json:"alliance_seniority"`
	BaseballTeam                     BBTeam               `json:"baseball_team"`
	GrossNationalIncome              float64              `json:"gross_national_income"`
	GrossDomesticProduct             float64              `json:"gross_domestic_product"`
	SoldierCasualties                int                  `json:"soldier_casualties"`
	SoldierKills                     int                  `json:"soldier_kills"`
	TankCasualties                   int                  `json:"tank_casualties"`
	TankKills                        int                  `json:"tank_kills"`
	AircraftCasualties               int                  `json:"aircraft_casualties"`
	AircraftKills                    int                  `json:"aircraft_kills"`
	ShipCasualties                   int                  `json:"ship_casualties"`
	ShipKills                        int                  `json:"ship_kills"`
	MissileCasualties                int                  `json:"missile_casualties"`
	MissileKills                     int                  `json:"missile_kills"`
	NukeCasualties                   int                  `json:"nuke_casualties"`
	NukeKills                        int                  `json:"nuke_kills"`
	SpyCasualties                    int                  `json:"spy_casualties"`
	SpyKills                         int                  `json:"spy_kills"`
	MoneyLooted                      float64              `json:"money_looted"`
}

type ID string

type AlliancePosition struct {
	ID                  ID           `json:"id"`
	Date                DateTimeAuto `json:"date"`
	AllianceID          ID           `json:"alliance_id"`
	Name                string       `json:"name"`
	CreatorID           ID           `json:"creator_id"`
	LastEditorID        ID           `json:"last_editor_id"`
	DateModified        DateTimeAuto `json:"date_modified"`
	PositionLevel       int          `json:"position_level"`
	Leader              bool         `json:"leader"`
	Heir                bool         `json:"heir"`
	Officer             bool         `json:"officer"`
	Member              bool         `json:"member"`
	Permissions         int          `json:"permissions"`
	ViewBank            bool         `json:"view_bank"`
	WithdrawBank        bool         `json:"withdraw_bank"`
	ChangePermissions   bool         `json:"change_permissions"`
	SeeSpies            bool         `json:"see_spies"`
	SeeResetTimers      bool         `json:"see_reset_timers"`
	TaxBrackets         bool         `json:"tax_brackets"`
	PostAnnouncements   bool         `json:"post_announcements"`
	ManageAnnouncements bool         `json:"manage_announcements"`
	AcceptApplicants    bool         `json:"accept_applicants"`
	RemoveMembers       bool         `json:"remove_members"`
	EditAllianceInfo    bool         `json:"edit_alliance_info"`
	ManageTreaties      bool         `json:"manage_treaties"`
	ManageMarketShare   bool         `json:"manage_market_share"`
	ManageEmbargoes     bool         `json:"manage_embargoes"`
	PromoteSelfToLeader bool         `json:"promote_self_to_leader"`
}

type DateTimeAuto string

type Alliance struct {
	ID                ID                 `json:"id"`
	Name              string             `json:"name"`
	Acronym           string             `json:"acronym"`
	Score             float64            `json:"score"`
	Color             string             `json:"color"`
	Date              DateTimeAuto       `json:"date"`
	Nations           []Nation           `json:"nations"`
	Treaties          []Treaty           `json:"treaties"`
	AlliancePositions []AlliancePosition `json:"alliance_positions"`
	AcceptMembers     bool               `json:"accept_members"`
	Flag              string             `json:"flag"`
	ForumLink         string             `json:"forum_link"`
	DiscordLink       string             `json:"discord_link"`
	WikiLink          string             `json:"wiki_link"`
	BankRecs          []BankRec          `json:"bankrecs"`
	Taxrecs           []BankRec          `json:"taxrecs"`
	TaxBrackets       []TaxBracket       `json:"tax_brackets"`
	Money             float64            `json:"money"`
	Coal              float64            `json:"coal"`
	Oil               float64            `json:"oil"`
	Uranium           float64            `json:"uranium"`
	Iron              float64            `json:"iron"`
	Bauxite           float64            `json:"bauxite"`
	Lead              float64            `json:"lead"`
	Gasoline          float64            `json:"gasoline"`
	Munitions         float64            `json:"munitions"`
	Steel             float64            `json:"steel"`
	Aluminum          float64            `json:"aluminum"`
	Food              float64            `json:"food"`
}

type Treaty struct {
	ID           ID           `json:"id"`
	Date         DateTimeAuto `json:"date"`
	TreatyType   string       `json:"treaty_type"`
	TreatyUrl    string       `json:"treaty_url"`
	TurnsLeft    int          `json:"turns_left"`
	Alliance1_id ID           `json:"alliance1_id"`
	Alliance1    Alliance     `json:"alliance1"`
	Alliance2_id ID           `json:"alliance2_id"`
	Alliance2    Alliance     `json:"alliance2"`
}

type BankRec struct {
	ID           ID           `json:"id"`
	Date         DateTimeAuto `json:"date"`
	SenderID     ID           `json:"sender_id"`
	SenderType   int          `json:"sender_type"`
	ReceiverID   ID           `json:"receiver_id"`
	ReceiverType int          `json:"receiver_type"`
	BankerID     ID           `json:"banker_id"`
	Note         string       `json:"note"`
	Money        float64      `json:"money"`
	Coal         float64      `json:"coal"`
	Oil          float64      `json:"oil"`
	Uranium      float64      `json:"uranium"`
	Iron         float64      `json:"iron"`
	Bauxite      float64      `json:"bauxite"`
	Lead         float64      `json:"lead"`
	Gasoline     float64      `json:"gasoline"`
	Munitions    float64      `json:"munitions"`
	Steel        float64      `json:"steel"`
	Aluminum     float64      `json:"aluminum"`
	Food         float64      `json:"food"`
	TaxID        ID           `json:"tax_id"`
}

type TaxBracket struct {
	ID              ID           `json:"id"`
	AllianceID      ID           `json:"alliance_id"`
	Alliance        Alliance     `json:"alliance"`
	Date            DateTimeAuto `json:"date"`
	DateModified    DateTimeAuto `json:"date_modified"`
	LastModifierID  ID           `json:"last_modifier_id"`
	LastModifier    Nation       `json:"last_modifier"`
	TaxRate         int          `json:"tax_rate"`
	ResourceTaxRate int          `json:"resource_tax_rate"`
	BracketName     string       `json:"bracket_name"`
}

type City struct {
	ID               ID      `json:"id"`
	NationID         ID      `json:"nation_id"`
	Nation           Nation  `json:"nation"`
	Name             string  `json:"name"`
	Date             Date    `json:"date"`
	Infrastructure   float64 `json:"infrastructure"`
	Land             float64 `json:"land"`
	Powered          bool    `json:"powered"`
	OilPower         int     `json:"oil_power"`
	WindPower        int     `json:"wind_power"`
	CoalPower        int     `json:"coal_power"`
	NuclearPower     int     `json:"nuclear_power"`
	CoalMine         int     `json:"coal_mine"`
	OilWell          int     `json:"oil_well"`
	UraniumMine      int     `json:"uranium_mine"`
	Barracks         int     `json:"barracks"`
	Farm             int     `json:"farm"`
	PoliceStation    int     `json:"police_station"`
	Hospital         int     `json:"hospital"`
	RecyclingCenter  int     `json:"recycling_center"`
	Subway           int     `json:"subway"`
	Supermarket      int     `json:"supermarket"`
	Bank             int     `json:"bank"`
	ShoppingMall     int     `json:"shopping_mall"`
	Stadium          int     `json:"stadium"`
	LeadMine         int     `json:"lead_mine"`
	IronMine         int     `json:"iron_mine"`
	BauxiteMine      int     `json:"bauxite_mine"`
	OilRefinery      int     `json:"oil_refinery"`
	AluminumRefinery int     `json:"aluminum_refinery"`
	SteelMill        int     `json:"steel_mill"`
	MunitionsFactory int     `json:"munitions_factory"`
	Factory          int     `json:"factory"`
	Hangar           int     `json:"hangar"`
	Drydock          int     `json:"drydock"`
	NukeDate         Date    `json:"nuke_date"`
}

type Date string

type Treasure struct {
	Name      string `json:"name"`
	Color     string `json:"color"`
	Continent string `json:"continent"`
	Bonus     int    `json:"bonus"`
	SpawnDate Date   `json:"spawn_date"`
	NationID  ID     `json:"nation_id"`
	Nation    Nation `json:"nation"`
}

type War struct {
	ID                     ID           `json:"id"`
	Date                   DateTimeAuto `json:"date"`
	Reason                 string       `json:"reason"`
	WarType                WarType      `json:"war_type"`
	GroundControl          ID           `json:"ground_control"`
	AirSuperiority         ID           `json:"air_superiority"`
	NavalBlockade          ID           `json:"naval_blockade"`
	WinnerID               ID           `json:"winner_id"`
	Attacks                []WarAttack  `json:"attacks"`
	TurnsLeft              int          `json:"turns_left"`
	AttID                  ID           `json:"att_id"`
	AttAllianceID          ID           `json:"att_alliance_id"`
	ATtacker               Nation       `json:"attacker"`
	DefID                  ID           `json:"def_id"`
	DefAllianceID          ID           `json:"def_alliance_id"`
	Defender               Nation       `json:"defender"`
	AttPoints              int          `json:"att_points"`
	DefPoints              int          `json:"def_points"`
	AttPeace               bool         `json:"att_peace"`
	DefPeace               bool         `json:"def_peace"`
	AttResistance          int          `json:"att_resistance"`
	DefResistance          int          `json:"def_resistance"`
	AttFortify             bool         `json:"att_fortify"`
	DefFortify             bool         `json:"def_fortify"`
	AttGasUsed             float64      `json:"att_gas_used"`
	DefGasUsed             float64      `json:"def_gas_used"`
	AttMunUsed             float64      `json:"att_mun_used"`
	DefMunUsed             float64      `json:"def_mun_used"`
	AttAlumUsed            int          `json:"att_alum_used"`
	DefAlumUsed            int          `json:"def_alum_used"`
	AttSteelUsed           int          `json:"att_steel_used"`
	DefSteelUsed           int          `json:"def_steel_used"`
	AttInfraDestroyed      float64      `json:"att_infra_destroyed"`
	DefInfraDestroyed      float64      `json:"def_infra_destroyed"`
	AttMoneyLooted         float64      `json:"att_money_looted"`
	DefMoneyLooted         float64      `json:"def_money_looted"`
	AttSoldiersKilled      int          `json:"att_soldiers_killed"`
	DefSoldiersKilled      int          `json:"def_soldiers_killed"`
	AttTanksKilled         int          `json:"att_tanks_killed"`
	DefTanksKilled         int          `json:"def_tanks_killed"`
	AttAircraftKilled      int          `json:"att_aircraft_killed"`
	DefAircraftKilled      int          `json:"def_aircraft_killed"`
	AttShipsKilled         int          `json:"att_ships_killed"`
	DefShipsKilled         int          `json:"def_ships_killed"`
	AttMissilesUsed        int          `json:"att_missiles_used"`
	DefMissilesUsed        int          `json:"def_missiles_used"`
	AttNukesUsed           int          `json:"att_nukes_used"`
	DefNukesUsed           int          `json:"def_nukes_used"`
	AttInfraDestroyedValue float64      `json:"att_infra_destroyed_value"`
	DefInfraDestroyedValue float64      `json:"def_infra_destroyed_value"`
}

type WarAttack struct {
	ID                    ID           `json:"id"`
	Date                  DateTimeAuto `json:"date"`
	AttID                 ID           `json:"att_id"`
	Attacker              Nation       `json:"attacker"`
	DefID                 ID           `json:"def_id"`
	Defender              Nation       `json:"defender"`
	Type                  AttackType   `json:"type"`
	WarID                 ID           `json:"war_id"`
	War                   War          `json:"war"`
	Victor                ID           `json:"victor"`
	Success               int          `json:"success"`
	Attcas1               int          `json:"attcas1"`
	Defcas1               int          `json:"defcas1"`
	Attcas2               int          `json:"attcas2"`
	Defcas2               int          `json:"defcas2"`
	CityID                ID           `json:"city_id"`
	InfraDestroyed        float64      `json:"infra_destroyed"`
	ImprovementsLost      int          `json:"improvements_lost"`
	MoneyStolen           float64      `json:"money_stolen"`
	LootInfo              string       `json:"loot_info"`
	ResistanceEliminated  int          `json:"resistance_eliminated"`
	CityInfraBefore       float64      `json:"city_infra_before"`
	InfraDestroyedValue   float64      `json:"infra_destroyed_value"`
	AttMunUsed            float64      `json:"att_mun_used"`
	DefMunUsed            float64      `json:"def_mun_used"`
	AttGasUsed            float64      `json:"att_gas_used"`
	DefGasUsed            float64      `json:"def_gas_used"`
	AircraftKilledByTanks int          `json:"aircraft_killed_by_tanks"`
}

type Bounty struct {
	ID       ID           `json:"id"`
	Date     DateTimeAuto `json:"date"`
	NationID ID           `json:"nation_id"`
	Nation   Nation       `json:"nation"`
	Amount   int          `json:"amount"`
	Type     BountyType   `json:"type"`
}

type BBTeam struct {
	ID       ID           `json:"id"`
	Date     DateTimeAuto `json:"date"`
	NationID ID           `json:"nation_id"`
	// nation       Nation       `json:"nation"`
	// ^ removes cycle
	Name        string     `json:"name"`
	Logo        string     `json:"logo"`
	HomeJersey  string     `json:"home_jersey"`
	AwayJersey  string     `json:"away_jersey"`
	Stadium     string     `json:"stadium"`
	Quality     int        `json:"quality"`
	Seating     int        `json:"seating"`
	Rating      float64    `json:"rating"`
	Wins        int        `json:"wins"`
	Glosses     int        `json:"glosses"`
	Runs        int        `json:"runs"`
	Homers      int        `json:"homers"`
	Strikeouts  int        `json:"strikeouts"`
	GamesPlayed int        `json:"games_played"`
	Games       []BBGame   `json:"games"`
	Players     []BBPlayer `json:"players"`
}

type BBGame struct {
	ID           ID           `json:"id"`
	Date         DateTimeAuto `json:"date"`
	HomeID       ID           `json:"home_id"`
	AwayID       ID           `json:"away_id"`
	HomeTeam     BBTeam       `json:"home_team"`
	AwayTeam     BBTeam       `json:"away_team"`
	HomeNationID ID           `json:"home_nation_id"`
	AwayNationID ID           `json:"away_nation_id"`
	HomeNation   Nation       `json:"home_nation"`
	AwayNation   Nation       `json:"away_nation"`
	StadiumName  string       `json:"stadium_name"`
	HomeScore    int          `json:"home_score"`
	AwayScore    int          `json:"away_score"`
	SimText      string       `json:"sim_text"`
	Highlights   string       `json:"highlights"`
	HomeRevenue  float64      `json:"home_revenue"`
	Spoils       float64      `json:"spoils"`
	Open         int          `json:"open"`
	Wager        float64      `json:"wager"`
}

type BBPlayer struct {
	ID        ID           `json:"id"`
	Date      DateTimeAuto `json:"date"`
	NationID  ID           `json:"nation_id"`
	Nation    Nation       `json:"nation"`
	TeamID    ID           `json:"team_id"`
	Team      BBTeam       `json:"team"`
	Name      string       `json:"name"`
	Age       int          `json:"age"`
	Position  string       `json:"position"`
	Pitching  float64      `json:"pitching"`
	Batting   float64      `json:"batting"`
	Speed     float64      `json:"speed"`
	Awareness float64      `json:"awareness"`
	Overall   float64      `json:"overall"`
	Birthday  int          `json:"birthday"`
}

type Color struct {
	Color     string `json:"color"`
	BlocName  string `json:"bloc_name"`
	TurnBonus int    `json:"turn_bonus"`
}

type GameInfo struct {
	GameDate  DateTimeAuto `json:"game_date"`
	Radiation Radiation    `json:"radiation"`
}

type Radiation struct {
	Global       float64 `json:"global"`
	NorthAmerica float64 `json:"north_america"`
	SouthAmerica float64 `json:"south_america"`
	Europe       float64 `json:"europe"`
	Africa       float64 `json:"africa"`
	Asia         float64 `json:"asia"`
	Australia    float64 `json:"australia"`
	Antarctica   float64 `json:"antarctica"`
}

type DateTime string

type NationPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []Nation      `json:"data"`
}

type PaginatorInfo struct {
	Count        int  `json:"count"`
	CurrentPage  int  `json:"currentPage"`
	FirstItem    int  `json:"firstItem"`
	HasMorePages bool `json:"hasMorePages"`
	LastItem     int  `json:"lastItem"`
	LastPage     int  `json:"lastPage"`
	PerPage      int  `json:"perPage"`
	Total        int  `json:"total"`
}

type AlliancePaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []Alliance    `json:"data"`
}

type TradePricePaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`

	Data []TradePrice
}

type TradePrice struct {
	ID        ID      `json:"id"`
	Date      Date    `json:"date"`
	Coal      float64 `json:"coal"`
	Oil       float64 `json:"oil"`
	Uranium   float64 `json:"uranium"`
	Iron      float64 `json:"iron"`
	Bauxite   float64 `json:"bauxite"`
	Lead      float64 `json:"lead"`
	Gasoline  float64 `json:"gasoline"`
	Munitions float64 `json:"munitions"`
	Steel     float64 `json:"steel"`
	Aluminum  float64 `json:"aluminum"`
	Food      float64 `json:"food"`
	Credits   float64 `json:"credits"`
}

type TradePaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []Trade       `json:"data"`
}

type Trade struct {
	ID              ID           `json:"id"`
	Type            TradeType    `json:"type"`
	Date            DateTimeAuto `json:"date"`
	SenderID        ID           `json:"sender_id"`
	ReceiverID      ID           `json:"receiver_id"`
	Sender          Nation       `json:"sender"`
	Receiver        Nation       `json:"receiver"`
	OfferResource   string       `json:"offer_resource"`
	OfferAmount     int          `json:"offer_amount"`
	BuyOrSell       string       `json:"buy_or_sell"`
	Price           int          `json:"price"`
	Accepted        bool         `json:"accepted"`
	DateAccepted    DateTimeAuto `json:"date_accepted"`
	OriginalTradeID ID           `json:"original_trade_id"`
}

type WarPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []War         `json:"data"`
}

type BountyPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []Bounty      `json:"data"`
}

type WarAttackPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []WarAttack   `json:"data"`
}

type TreatyPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []Treaty      `json:"data"`
}

type CityPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []City        `json:"data"`
}

type BankRecPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []BankRec     `json:"data"`
}

type BBGamePaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []BBGame      `json:"data"`
}

type BBTeamPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []BBTeam      `json:"data"`
}

type BBPlayerPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []BBPlayer    `json:"data"`
}

type TreasureTradePaginator struct {
	PaginatorInfo PaginatorInfo   `json:"paginatorInfo"`
	Data          []TreasureTrade `json:"data"`
}

type TreasureTrade struct {
	ID              ID           `json:"id"`
	OfferDate       DateTimeAuto `json:"offer_date"`
	AcceptDate      DateTimeAuto `json:"accept_date"`
	SenderID        ID           `json:"sender_id"`
	Sender          Nation       `json:"sender"`
	ReceiverID      ID           `json:"receiver_id"`
	Receiver        Nation       `json:"receiver"`
	Buying          bool         `json:"buying"`
	Selling         bool         `json:"selling"`
	Treasure        string       `json:"treasure"`
	Money           int          `json:"money"`
	Accepted        bool         `json:"accepted"`
	Rejected        bool         `json:"rejected"`
	SellerCancelled bool         `json:"seller_cancelled"`
}

type EmbargoPaginator struct {
	PaginatorInfo PaginatorInfo `json:"paginatorInfo"`
	Data          []Embargo     `json:"data"`
}

type Embargo struct {
	ID         ID     `json:"id"`
	Date       Date   `json:"date"`
	SenderID   ID     `json:"sender_id"`
	Sender     Nation `json:"sender"`
	ReceiverID ID     `json:"receiver_id"`
	Receiver   Nation `json:"receiver"`
	Reason     string `json:"reason"`
}

type Mutation struct {
	BankDeposit  BankRec `json:"bankDeposit"`
	BankWithdraw BankRec `json:"bankWithdraw"`
}

type Subscription struct {
	AllianceCreate         Alliance         `json:"allianceCreate"`
	AllianceDelete         Alliance         `json:"allianceDelete"`
	AllianceUpdate         Alliance         `json:"allianceUpdate"`
	AlliancePositionCreate AlliancePosition `json:"alliancePositionCreate"`
	AlliancePositionDelete AlliancePosition `json:"alliancePositionDelete"`
	AlliancePositionUpdate AlliancePosition `json:"alliancePositionUpdate"`
	BankRecCreate          BankRec          `json:"bankrecCreate"`
	BBGameCreate           BBGame           `json:"bbgameCreate"`
	BBGameDelete           BBGame           `json:"bbgameDelete"`
	BBGameUpdate           BBGame           `json:"bbgameUpdate"`
	BBPlayerCreate         BBPlayer         `json:"bbplayerCreate"`
	BBPlayerDelete         BBPlayer         `json:"bbplayerDelete"`
	BBPlayerUpdate         BBPlayer         `json:"bbplayerUpdate"`
	BBTeamCreate           BBTeam           `json:"bbteamCreate"`
	BBTeamDelete           BBTeam           `json:"bbteamDelete"`
	BBTeamUpdate           BBTeam           `json:"bbteamUpdate"`
	BountyCreate           Bounty           `json:"bountyCreate"`
	BountyDelete           Bounty           `json:"bountyDelete"`
	BountyUpdate           Bounty           `json:"bountyUpdate"`
	CityCreate             City             `json:"cityCreate"`
	CityDelete             City             `json:"cityDelete"`
	CityUpdate             City             `json:"cityUpdate"`
	EmbargoCreate          Embargo          `json:"embargoCreate"`
	EmbargoDelete          Embargo          `json:"embargoDelete"`
	NationCreate           Nation           `json:"nationCreate"`
	NationDelete           Nation           `json:"nationDelete"`
	NationUpdate           Nation           `json:"nationUpdate"`
	TaxBracketCreate       TaxBracket       `json:"taxBracketCreate"`
	TaxBracketDelete       TaxBracket       `json:"taxBracketDelete"`
	TaxBracketUpdate       TaxBracket       `json:"taxBracketUpdate"`
	TradeCreate            Trade            `json:"tradeCreate"`
	TradeDelete            Trade            `json:"tradeDelete"`
	TradeUpdate            Trade            `json:"tradeUpdate"`
	TreasureTradeUpdate    TreasureTrade    `json:"treasureTradeUpdate"`
	TreatyCreate           Treaty           `json:"treatyCreate"`
	TreatyUpdate           Treaty           `json:"treatyUpdate"`
	WarCreate              War              `json:"warCreate"`
	WarDelete              War              `json:"warDelete"`
	WarUpdate              War              `json:"warUpdate"`
	WarAttackCreate        WarAttack        `json:"warAttackCreate"`
	WarAttackDelete        WarAttack        `json:"warAttackDelete"`
}

type DateTimeTz string

type SimplePaginatorInfo struct {
	Count        int  `json:"count"`
	CurrentPage  int  `json:"currentPage"`
	FirstItem    int  `json:"firstItem"`
	LastItem     int  `json:"lastItem"`
	PerPage      int  `json:"perPage"`
	HasMorePages bool `json:"hasMorePages"`
}

type PageInfo struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
	Total           int    `json:"total"`
	Count           int    `json:"count"`
	CurrentPage     int    `json:"currentPage"`
	LastPage        int    `json:"lastPage"`
}
