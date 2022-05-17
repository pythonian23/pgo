package pgo

const whoAllianceQuery = "{alliances(%s){data{name id acronym score color nations{alliance_position score wars(active:true){war_type att_alliance_id att_resistance def_resistance}}treaties{alliance1{name}alliance2{name}treaty_type}}}}"
const whoNationQuery = "{nations(%s){data{nation_name}}}"
