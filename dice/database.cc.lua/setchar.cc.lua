{{$user:= .User.ID}}

{{$force:= (dbGet $user "force").Value}}
{{$endurance := (dbGet $user "endurance").Value}}
{{$agi:=(dbGet $user "agi").Value}}
{{$preci:=(dbGet $user "preci").Value}}
{{$intel:=(dbGet $user "intelligence").Value}}
{{$karma:=(dbGet $user "karma").Value}}

{{$iforce:=(dbGet  $user "i_force").Value}}
{{$iendu:=(dbGet $user "i_endu").Value}}
{{$iagi:=(dbGet $user "i_agi").Value}}
{{$ipreci:=(dbGet $user "i_preci").Value}}
{{$iintel:=(dbGet $user "i_intel").Value}}
{{$ikarma:=(dbGet $user "i_karma").Value}}


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
			{{dbSet $user "i_karma" (toInt (index .CmdArgs 14))}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{$force}}
	:white_small_square: Endurance : {{$endurance}}
	:white_small_square: Agilité : {{$agi}}
	:white_small_square: Précision : {{$preci}}
	:white_small_square: Intelligence : {{$intel}}
	:white_small_square: Karma : {{$karma}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{$iforce}}
	:white_small_square: Endurance : {{$iendu}}
	:white_small_square: Agilité : {{$iagi}}
	:white_small_square: Précision : {{$ipreci}}
	:white_small_square: Intelligence : {{$iintel}}
	:white_small_square: Karma : {{$ikarma}}

			{{else}}
				**Bad User ID**
			{{end}}

	{{else if eq (index .CmdArgs 0) "-implant"}}
		{{dbSet $user "i_force" (toInt (index .CmdArgs 1)) }}
		{{dbSet $user "i_endu" (toInt (index .CmdArgs 2))}}
		{{dbSet $user "i_agi" (toInt (index .CmdArgs 3))}}
		{{dbSet $user "i_preci" (toInt (index .CmdArgs 4))}}
		{{dbSet $user "i_intel" (toInt (index .CmdArgs 5))}}
		{{dbSet $user "i_karma" (toInt (index .CmdArgs 6))}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{$iforce}}
	:white_small_square: Endurance : {{$iendu}}
	:white_small_square: Agilité : {{$iagi}}
	:white_small_square: Précision : {{$ipreci}}
	:white_small_square: Intelligence : {{$iintel}}
	:white_small_square: Karma : {{$ikarma}}

	{{else if and (eq (index .CmdArgs 0) "-stats") (ne (toFloat (index .CmdArgs 1)) (toFloat 0))}}
		{{dbSet .User.ID "force" (toInt (index .CmdArgs 1)) }}
		{{dbSet .User.ID "endurance" (toInt (index .CmdArgs 2))}}
		{{dbSet .User.ID "agi" (toInt (index .CmdArgs 3))}}
		{{dbSet .User.ID "preci" (toInt (index .CmdArgs 4))}}
		{{dbSet .User.ID "intelligence" (toInt (index .CmdArgs 5))}}
		{{dbSet .User.ID "karma" (toInt (index .CmdArgs 6))}}

	**Statistiques de <@{{$user}}>**
		:white_small_square: Force : {{$force}}
		:white_small_square: Endurance : {{$endurance}}
		:white_small_square: Agilité : {{$agi}}
		:white_small_square: Précision : {{$preci}}
		:white_small_square: Intelligence : {{$intel}}
		:white_small_square: Karma : {{$karma}}

	{{else if and (eq (index .CmdArgs 0) "-stats") (eq (index .CmdArgs 1) "-ID") }}
	{{if gt (toFloat (index .CmdArgs 2)) (toFloat 11)}}
		{{$user = (toInt (index .CmdArgs 2))}}
		{{dbSet $user "force" (toInt (index .CmdArgs 3)) }}
		{{dbSet $user "endurance" (toInt (index .CmdArgs 4))}}
		{{dbSet $user "agi" (toInt (index .CmdArgs 5))}}
		{{dbSet $user "preci" (toInt (index .CmdArgs 6))}}
		{{dbSet $user "intelligence" (toInt (index .CmdArgs 7))}}
		{{dbSet $user "karma" (toInt (index .CmdArgs 8))}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{$force}}
	:white_small_square: Endurance : {{$endurance}}
	:white_small_square: Agilité : {{$agi}}
	:white_small_square: Précision : {{$preci}}
	:white_small_square: Intelligence : {{$intel}}
	:white_small_square: Karma : {{$karma}}


	{{else}}
		**BAD USER ID**
	{{end}}

	{{else if and (eq (index .CmdArgs 0) "-all") (ne (toFloat (index .CmdArgs 1)) (toFloat 0))}}
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
		{{dbSet $user "i_karma" (toInt (index .CmdArgs 12))}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{$force}}
	:white_small_square: Endurance : {{$endurance}}
	:white_small_square: Agilité : {{$agi}}
	:white_small_square: Précision : {{$preci}}
	:white_small_square: Intelligence : {{$intel}}
	:white_small_square: Karma : {{$karma}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{$iforce}}
	:white_small_square: Endurance : {{$iendu}}
	:white_small_square: Agilité : {{$iagi}}
	:white_small_square: Précision : {{$ipreci}}
	:white_small_square: Intelligence : {{$iintel}}
	:white_small_square: Karma : {{$ikarma}}


	{{else if and (eq (index .CmdArgs 1) "-ID") (eq (index .CmdArgs 0) "-implant" )}}
	{{if gt (toFloat (index .CmdArgs 2)) (toFloat 11)}}
		{{$user = (toInt (index .CmdArgs 2))}}
		{{dbSet $user "i_force" (toInt (index .CmdArgs 3)) }}
		{{dbSet $user "i_endu" (toInt (index .CmdArgs 4))}}
		{{dbSet $user "i_agi" (toInt (index .CmdArgs 5))}}
		{{dbSet $user "i_preci" (toInt (index .CmdArgs 6))}}
		{{dbSet $user "i_intel" (toInt (index .CmdArgs 7))}}
		{{dbSet $user "i_karma" (toInt (index .CmdArgs 8))}}
	{{else}}
		**BAD USER ID**
	{{end}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{$iforce}}
	:white_small_square: Endurance : {{$iendu}}
	:white_small_square: Agilité : {{$iagi}}
	:white_small_square: Précision : {{$ipreci}}
	:white_small_square: Intelligence : {{$iintel}}
	:white_small_square: Karma : {{$ikarma}}

	{{else}}
	**Usage** : `$setchar -(all|stats|implant) (-ID <id>) Force Endurance Agilité Précision Intelligence Karma`
	{{end}}
{{else}}
**Usage** : `$set -(all|stats|implant) (-ID <id>) Force Endurance Agilité Précision Intelligence Karma`

{{end}}
