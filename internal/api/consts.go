package api

type AlliancePositionEnum string

const (
	PositionNoalliance AlliancePositionEnum = "NOALLIANCE"
	PositionApplicant  AlliancePositionEnum = "APPLICANT"
	PositionMember     AlliancePositionEnum = "MEMBER"
	PositionOfficer    AlliancePositionEnum = "OFFICER"
	PositionHeir       AlliancePositionEnum = "HEIR"
	PositionLeader     AlliancePositionEnum = "LEADER"
)

type WarPolicy string

const (
	WarPolicyAttrition  WarPolicy = "ATTRITION"
	WarPolicyTurtle     WarPolicy = "TURTLE"
	WarPolicyBlitzkrieg WarPolicy = "BLITZKRIEG"
	WarPolicyFortress   WarPolicy = "FORTRESS"
	WarPolicyMoneybags  WarPolicy = "MONEYBAGS"
	WarPolicyPirate     WarPolicy = "PIRATE"
	WarPolicyTactician  WarPolicy = "TACTICIAN"
	WarPolicyGuardian   WarPolicy = "GUARDIAN"
	WarPolicyCovert     WarPolicy = "COVERT"
	WarPolicyArcane     WarPolicy = "ARCANE"
)

type DomesticPolicy string

const (
	DomesticPolicyManifestDestiny          DomesticPolicy = "MANIFEST_DESTINY"
	DomesticPolicyOpenMarkets              DomesticPolicy = "OPEN_MARKETS"
	DomesticPolicyTechnologicalAdvancement DomesticPolicy = "TECHNOLOGICAL_ADVANCEMENT"
	DomesticPolicyImperialism              DomesticPolicy = "IMPERIALISM"
	DomesticPolicyUrbanization             DomesticPolicy = "URBANIZATION"
	DomesticPolicyRapid_expansion          DomesticPolicy = "RAPID_EXPANSION"
)

type WarType string

const (
	WarOrdinary  WarType = "ORDINARY"
	WarAttrition WarType = "ATTRITION"
	WarRaid      WarType = "RAID"
)

type AttackType string

const (
	AttackAirvinfra    AttackType = "AIRVINFRA"
	AttackAirvsoldiers AttackType = "AIRVSOLDIERS"
	AttackAirvtanks    AttackType = "AIRVTANKS"
	AttackAirvmoney    AttackType = "AIRVMONEY"
	AttackAirvships    AttackType = "AIRVSHIPS"
	AttackAirvair      AttackType = "AIRVAIR"
	AttackGround       AttackType = "GROUND"
	AttackMissile      AttackType = "MISSILE"
	AttackMissilefail  AttackType = "MISSILEFAIL"
	AttackNuke         AttackType = "NUKE"
	AttackNukefail     AttackType = "NUKEFAIL"
	AttackNaval        AttackType = "NAVAL"
	AttackFortify      AttackType = "FORTIFY"
	AttackPeace        AttackType = "PEACE"
	AttackVictory      AttackType = "VICTORY"
	AttackAllianceloot AttackType = "ALLIANCELOOT"
)

type BountyType string

const (
	BountyOrdinary  BountyType = "ORDINARY"
	BountyAttrition BountyType = "ATTRITION"
	BountyRaid      BountyType = "RAID"
	BountyNuclear   BountyType = "NUCLEAR"
)

type TradeType string

const (
	TradeGlobal   TradeType = "GLOBAL"
	TradePersonal TradeType = "PERSONAL"
	TradeAlliance TradeType = "ALLIANCE"
)
