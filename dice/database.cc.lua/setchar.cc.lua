{{$user:= .User.ID}}

{{if .CmdArgs}}
	{{if and (eq (index .CmdArgs 0) "-all") (eq (index .CmdArgs 1) "-ID") }}
		{{if gt (toFloat (index .CmdArgs 2)) (toFloat 11)}}
			{{$user = (toInt (index .CmdArgs 2))}}
			{{dbSet $user "force" (toInt (index .CmdArgs 3)) }}
			{{dbSet $user "i_force" (toInt (index .CmdArgs 4)) }}
			{{dbSet $user "endurance" (toInt (index .CmdArgs 5))}}
			{{dbSet $user "i_endu" (toInt (index .CmdArgs 6))}}
			{{dbSet $user "agi" (toInt (index .CmdArgs 7))}}
			{{dbSet $user "i_agi" (toInt (index .CmdArgs 8))}}
			{{dbSet $user "preci" (toInt (index .CmdArgs 9))}}
			{{dbSet $user "i_preci" (toInt (index .CmdArgs 10))}}
			{{dbSet $user "intelligence" (toInt (index .CmdArgs 11))}}
			{{dbSet $user "i_intel" (toInt (index .CmdArgs 12))}}
			{{dbSet $user "karma" (toInt (index .CmdArgs 13))}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{(dbGet $user "force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "endurance").Value}}
	:white_small_square: Agilité : {{(dbGet $user "agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "intelligence").Value}}
	:white_small_square: Karma : {{(dbGet $user "karma").Value}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{(dbGet $user "i_force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "i_endu").Value}}
	:white_small_square: Agilité : {{(dbGet $user "i_agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "i_preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "i_intel").Value}}

			{{else}}
				**Bad User ID**
			{{end}}


	{{else if and (eq (index .CmdArgs 0) "-stats") (eq (index .CmdArgs 1) "-ID") }}
	{{if gt (toFloat (index .CmdArgs 2)) (toFloat 11)}}
		{{$user = (toInt (index .CmdArgs 2)) }}
		{{dbSet $user "force" (toInt (index .CmdArgs 3)) }}
		{{dbSet $user "endurance" (toInt (index .CmdArgs 4))}}
		{{dbSet $user "agi" (toInt (index .CmdArgs 5))}}
		{{dbSet $user "preci" (toInt (index .CmdArgs 6))}}
		{{dbSet $user "intelligence" (toInt (index .CmdArgs 7))}}
		{{dbSet $user "karma" (toInt (index .CmdArgs 8))}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{(dbGet $user "force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "endurance").Value}}
	:white_small_square: Agilité : {{(dbGet $user "agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "intelligence").Value}}
	:white_small_square: Karma : {{(dbGet $user "karma").Value}}


	{{else}}
		**BAD USER ID**
	{{end}}


	{{else if and (eq (index .CmdArgs 0) "-implant" ) (eq (index .CmdArgs 1) "-ID") }}
		{{if gt (toFloat (index .CmdArgs 2)) (toFloat 11)}}
			{{$user = (toInt (index .CmdArgs 2))}}
			{{dbSet $user "i_force" (toInt (index .CmdArgs 3)) }}
			{{dbSet $user "i_endu" (toInt (index .CmdArgs 4))}}
			{{dbSet $user "i_agi" (toInt (index .CmdArgs 5))}}
			{{dbSet $user "i_preci" (toInt (index .CmdArgs 6))}}
			{{dbSet $user "i_intel" (toInt (index .CmdArgs 7))}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{(dbGet $user "i_force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "i_endu").Value}}
	:white_small_square: Agilité : {{(dbGet $user "i_agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "i_preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "i_intel").Value}}

		{{else}}
			**BAD USER ID**
		{{end}}

		{{else if and (eq (index .CmdArgs 0) "-all") (lt (toFloat (index .CmdArgs 1)) (toFloat 11))}}
			{{dbSet $user "force" (toInt (index .CmdArgs 1)) }}
			{{dbSet $user "i_force" (toInt (index .CmdArgs 2)) }}
			{{dbSet $user "endurance" (toInt (index .CmdArgs 3))}}
			{{dbSet $user "i_endu" (toInt (index .CmdArgs 4))}}
			{{dbSet $user "agi" (toInt (index .CmdArgs 5))}}
			{{dbSet $user "i_agi" (toInt (index .CmdArgs 6))}}
			{{dbSet $user "preci" (toInt (index .CmdArgs 7))}}
			{{dbSet $user "i_preci" (toInt (index .CmdArgs 8))}}
			{{dbSet $user "intelligence" (toInt (index .CmdArgs 9))}}
			{{dbSet $user "i_intel" (toInt (index .CmdArgs 10))}}
			{{dbSet $user "karma" (toInt (index .CmdArgs 11))}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{(dbGet $user "force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "endurance").Value}}
	:white_small_square: Agilité : {{(dbGet $user "agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "intelligence").Value}}
	:white_small_square: Karma : {{(dbGet $user "karma").Value}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{(dbGet $user "i_force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "i_endu").Value}}
	:white_small_square: Agilité : {{(dbGet $user "i_agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "i_preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "i_intel").Value}}

	{{else if and (eq (index .CmdArgs 0) "-stats") (lt (toFloat (index .CmdArgs 1)) (toFloat 11))}}
		{{dbSet .User.ID "force" (toInt (index .CmdArgs 1)) }}
		{{dbSet .User.ID "endurance" (toInt (index .CmdArgs 2))}}
		{{dbSet .User.ID "agi" (toInt (index .CmdArgs 3))}}
		{{dbSet .User.ID "preci" (toInt (index .CmdArgs 4))}}
		{{dbSet .User.ID "intelligence" (toInt (index .CmdArgs 5))}}
		{{dbSet .User.ID "karma" (toInt (index .CmdArgs 6))}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{(dbGet $user "force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "endurance").Value}}
	:white_small_square: Agilité : {{(dbGet $user "agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "intelligence").Value}}
	:white_small_square: Karma : {{(dbGet $user "karma").Value}}

	{{else if and (eq (index .CmdArgs 0) "-implant") (lt (toFloat (index .CmdArgs 1)) (toFloat 11))}}
		{{dbSet $user "i_force" (toInt (index .CmdArgs 1)) }}
		{{dbSet $user "i_endu" (toInt (index .CmdArgs 2))}}
		{{dbSet $user "i_agi" (toInt (index .CmdArgs 3))}}
		{{dbSet $user "i_preci" (toInt (index .CmdArgs 4))}}
		{{dbSet $user "i_intel" (toInt (index .CmdArgs 5))}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{(dbGet $user "i_force").Value}}
	:white_small_square: Endurance : {{(dbGet $user "i_endu").Value}}
	:white_small_square: Agilité : {{(dbGet $user "i_agi").Value}}
	:white_small_square: Précision : {{(dbGet $user "i_preci").Value}}
	:white_small_square: Intelligence : {{(dbGet $user "i_intel").Value}}

	{{else}}
	**Usage** : `$setchar -(all|stats|implant) (-ID <id>) Force Endurance Agilité Précision Intelligence Karma`
	{{end}}
{{else}}
**Usage** : `$setchar -(all|stats|implant) (-ID <id>) Force Endurance Agilité Précision Intelligence Karma`

{{end}}
