{{if .CmdArgs}}
	{{$name := toRune (lower (index .CmdArgs 0))}}
	{{$user :=""}}
	{{range $name}}
	{{- $user = (print $user .)}}
	{{- end}}
	{{$user = (toInt $user)}}

	{{if eq (index .CmdArgs 1) "-all"}}
		{{dbSet $user "force" (toInt (index .CmdArgs 2)) }}
		{{dbSet $user "i_force" (toInt (index .CmdArgs 3)) }}
		{{dbSet $user "endurance" (toInt (index .CmdArgs 4))}}
		{{dbSet $user "i_endu" (toInt (index .CmdArgs 5))}}
		{{dbSet $user "agi" (toInt (index .CmdArgs 6))}}
		{{dbSet $user "i_agi" (toInt (index .CmdArgs 7))}}
		{{dbSet $user "preci" (toInt (index .CmdArgs 8))}}
		{{dbSet $user "i_preci" (toInt (index .CmdArgs 9))}}
		{{dbSet $user "intelligence" (toInt (index .CmdArgs 10))}}
		{{dbSet $user "i_intel" (toInt (index .CmdArgs 11))}}
		{{dbSet $user "karma" (toInt (index .CmdArgs 12))}}

**Statistiques de {{index .CmdArgs 0}}**
	:white_small_square: Force : {{(dbGet $user "force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "endurance").Value}}
	:white_small_square: Agilité : {{(dbGet $user "agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "intelligence").Value}}
	:white_small_square: Karma : {{(dbGet $user "karma").Value}}

**Implant de {{index .CmdArgs 0}}** :
	:white_small_square: Force : {{(dbGet $user "i_force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "i_endu").Value}}
	:white_small_square: Agilité : {{(dbGet $user "i_agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "i_preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "i_intel").Value}}

	{{else if eq (index .CmdArgs 1) "-stats"}}
		{{dbSet $user "force" (toInt (index .CmdArgs 2)) }}
		{{dbSet $user "endurance" (toInt (index .CmdArgs 3))}}
		{{dbSet $user "agi" (toInt (index .CmdArgs 4))}}
		{{dbSet $user "preci" (toInt (index .CmdArgs 5))}}
		{{dbSet $user "intelligence" (toInt (index .CmdArgs 6))}}
		{{dbSet $user "karma" (toInt (index .CmdArgs 7))}}

**Statistiques de {{index .CmdArgs 0}}**
	:white_small_square: Force : {{(dbGet $user "force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "endurance").Value}}
	:white_small_square: Agilité : {{(dbGet $user "agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "intelligence").Value}}
	:white_small_square: Karma : {{(dbGet $user "karma").Value}}

	{{else if eq (index .CmdArgs 1) "-implant"}}
		{{dbSet $user "i_force" (toInt (index .CmdArgs 2)) }}
		{{dbSet $user "i_endu" (toInt (index .CmdArgs 3))}}
		{{dbSet $user "i_agi" (toInt (index .CmdArgs 4))}}
		{{dbSet $user "i_preci" (toInt (index .CmdArgs 5))}}
		{{dbSet $user "i_intel" (toInt (index .CmdArgs 6))}}

**Implant de {{index .CmdArgs 0}}** :
	:white_small_square: Force : {{(dbGet $user "i_force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "i_endu").Value}}
	:white_small_square: Agilité : {{(dbGet $user "i_agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "i_preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "i_intel").Value}}
	{{else}}
		**Usage** : `$setreroll -(all|stats|implant) (nomperso) Force Endurance Agilité Précision Intelligence Karma`
	{{end}}
{{else}}
**Usage** : `$setreroll -(all|stats|implant) (nomperso) Force Endurance Agilité Précision Intelligence Karma`
{{end}}
