{{$user:= .User.ID}}

{{/* Stats dictionnaire */}}

{{$stats := sdict}}
{{with (dbGet $user "stats")}}
	{{$stats = sdict .Value}}
{{end}}

{{if .CmdArgs}}
	{{if and (eq (index .CmdArgs 0) "-all") (eq (index .CmdArgs 1) "-ID") }}
		{{if gt (toFloat (index .CmdArgs 2)) (toFloat 11)}}
			{{$user = (toInt (index .CmdArgs 2))}}
			{{$stats.Set "force" (toInt (index .CmdArgs 3)) }}
			{{$stats.Set "i_force" (toInt (index .CmdArgs 4)) }}
			{{$stats.Set "endurance" (toInt (index .CmdArgs 5))}}
			{{$stats.Set "i_endu" (toInt (index .CmdArgs 6))}}
			{{$stats.Set "agi" (toInt (index .CmdArgs 7))}}
			{{$stats.Set "i_agi" (toInt (index .CmdArgs 8))}}
			{{$stats.Set "preci" (toInt (index .CmdArgs 9))}}
			{{$stats.Set "i_preci" (toInt (index .CmdArgs 10))}}
			{{$stats.Set "intelligence" (toInt (index .CmdArgs 11))}}
			{{$stats.Set "i_intel" (toInt (index .CmdArgs 12))}}
			{{$stats.Set "karma" (toInt (index .CmdArgs 13))}}
			{{dbSet $user "stats" $stats}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{$stats.Get "force"}}
	:white_small_square: Endurance : {{$stats.Get "endurance"}}
	:white_small_square: Agilité : {{$stats.Get "agi"}}
	:white_small_square: Précision : {{$stats.Get "preci"}}
	:white_small_square: Intelligence : {{$stats.Get "intelligence"}}
	:white_small_square: Karma : {{$stats.Get "karma"}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{$stats.Get "i_force"}}
	:white_small_square: Endurance : {{$stats.Get "i_endu"}}
	:white_small_square: Agilité : {{$stats.Get "i_agi"}}
	:white_small_square: Précision : {{$stats.Get "i_preci"}}
	:white_small_square: Intelligence : {{$stats.Get "i_intel"}}

			{{else}}
				**Bad User ID**
			{{end}}


	{{else if and (eq (index .CmdArgs 0) "-stats") (eq (index .CmdArgs 1) "-ID") }}
	{{if gt (toFloat (index .CmdArgs 2)) (toFloat 11)}}
		{{$user = (toInt (index .CmdArgs 2)) }}
		{{$stats.Set "force" (toInt (index .CmdArgs 3)) }}
		{{$stats.Set "endurance" (toInt (index .CmdArgs 4))}}
		{{$stats.Set "agi" (toInt (index .CmdArgs 5))}}
		{{$stats.Set "preci" (toInt (index .CmdArgs 6))}}
		{{$stats.Set "intelligence" (toInt (index .CmdArgs 7))}}
		{{$stats.Set "karma" (toInt (index .CmdArgs 8))}}
		{{dbSet $user "stats" $stats}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{$stats.Get "force"}}
	:white_small_square: Endurance : {{$stats.Get "endurance"}}
	:white_small_square: Agilité : {{$stats.Get "agi"}}
	:white_small_square: Précision : {{$stats.Get "preci"}}
	:white_small_square: Intelligence : {{$stats.Get "intelligence"}}
	:white_small_square: Karma : {{$stats.Get "karma"}}


	{{else}}
		**BAD USER ID**
	{{end}}


	{{else if and (eq (index .CmdArgs 0) "-implant" ) (eq (index .CmdArgs 1) "-ID") }}
		{{if gt (toFloat (index .CmdArgs 2)) (toFloat 11)}}
			{{$user = (toInt (index .CmdArgs 2))}}
			{{$stats.Set "i_force" (toInt (index .CmdArgs 3)) }}
			{{$stats.Set "i_endu" (toInt (index .CmdArgs 4))}}
			{{$stats.Set "i_agi" (toInt (index .CmdArgs 5))}}
			{{$stats.Set "i_preci" (toInt (index .CmdArgs 6))}}
			{{$stats.Set "i_intel" (toInt (index .CmdArgs 7))}}
			{{dbSet $user "stats" $stats}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{$stats.Get "i_force"}}
	:white_small_square: Endurance : {{$stats.Get "i_endu"}}
	:white_small_square: Agilité : {{$stats.Get "i_agi"}}
	:white_small_square: Précision : {{$stats.Get "i_preci"}}
	:white_small_square: Intelligence : {{$stats.Get "i_intel"}}

		{{else}}
			**BAD USER ID**
		{{end}}

		{{else if and (eq (index .CmdArgs 0) "-all") (lt (toFloat (index .CmdArgs 1)) (toFloat 11))}}
			{{$stats.Set "force" (toInt (index .CmdArgs 1)) }}
			{{$stats.Set "i_force" (toInt (index .CmdArgs 2)) }}
			{{$stats.Set "endurance" (toInt (index .CmdArgs 3))}}
			{{$stats.Set "i_endu" (toInt (index .CmdArgs 4))}}
			{{$stats.Set "agi" (toInt (index .CmdArgs 5))}}
			{{$stats.Set "i_agi" (toInt (index .CmdArgs 6))}}
			{{$stats.Set "preci" (toInt (index .CmdArgs 7))}}
			{{$stats.Set "i_preci" (toInt (index .CmdArgs 8))}}
			{{$stats.Set "intelligence" (toInt (index .CmdArgs 9))}}
			{{$stats.Set "i_intel" (toInt (index .CmdArgs 10))}}
			{{$stats.Set "karma" (toInt (index .CmdArgs 11))}}
			{{dbSet $user "stats" $stats}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{$stats.Get "force"}}
	:white_small_square: Endurance : {{$stats.Get "endurance"}}
	:white_small_square: Agilité : {{$stats.Get "agi"}}
	:white_small_square: Précision : {{$stats.Get "preci"}}
	:white_small_square: Intelligence : {{$stats.Get "intelligence"}}
	:white_small_square: Karma : {{$stats.Get "karma"}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{$stats.Get "i_force"}}
	:white_small_square: Endurance : {{$stats.Get "i_endu"}}
	:white_small_square: Agilité : {{$stats.Get "i_agi"}}
	:white_small_square: Précision : {{$stats.Get "i_preci"}}
	:white_small_square: Intelligence : {{$stats.Get "i_intel"}}

	{{else if and (eq (index .CmdArgs 0) "-stats") (lt (toFloat (index .CmdArgs 1)) (toFloat 11))}}
		{{$stats.Set "force" (toInt (index .CmdArgs 1)) }}
		{{$stats.Set "endurance" (toInt (index .CmdArgs 2))}}
		{{$stats.Set "agi" (toInt (index .CmdArgs 3))}}
		{{$stats.Set "preci" (toInt (index .CmdArgs 4))}}
		{{$stats.Set "intelligence" (toInt (index .CmdArgs 5))}}
		{{$stats.Set "karma" (toInt (index .CmdArgs 6))}}
		{{dbSet $user "stats" $stats}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{$stats.Get "force"}}
	:white_small_square: Endurance : {{$stats.Get "endurance"}}
	:white_small_square: Agilité : {{$stats.Get "agi"}}
	:white_small_square: Précision : {{$stats.Get "preci"}}
	:white_small_square: Intelligence : {{$stats.Get "intelligence"}}
	:white_small_square: Karma : {{$stats.Get "karma"}}

	{{else if and (eq (index .CmdArgs 0) "-implant") (lt (toFloat (index .CmdArgs 1)) (toFloat 11))}}
		{{$stats.Set "i_force" (toInt (index .CmdArgs 1)) }}
		{{$stats.Set "i_endu" (toInt (index .CmdArgs 2))}}
		{{$stats.Set "i_agi" (toInt (index .CmdArgs 3))}}
		{{$stats.Set "i_preci" (toInt (index .CmdArgs 4))}}
		{{$stats.Set "i_intel" (toInt (index .CmdArgs 5))}}
		{{dbSet $user "stats" $stats}}

**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{$stats.Get "i_force"}}
	:white_small_square: Endurance : {{$stats.Get "i_endu"}}
	:white_small_square: Agilité : {{$stats.Get "i_agi"}}
	:white_small_square: Précision : {{$stats.Get "i_preci"}}
	:white_small_square: Intelligence : {{$stats.Get "i_intel"}}
	{{else}}
	**Usage** : `$setchar -(all|stats|implant) (-ID <id>) Force Endurance Agilité Précision Intelligence Karma`
	{{end}}
{{else}}
**Usage** : `$setchar -(all|stats|implant) (-ID <id>) Force Endurance Agilité Précision Intelligence Karma`
{{end}}
{{deleteTrigger 1}}
