{{$user := .User.ID}}

{{if .CmdArgs}}
	{{if (userArg (index .CmdArgs 0))}}
		{{with (userArg (index .CmdArgs 0))}}
			{{$user = .}}
			{{$user = $user.ID}}
		{{end}}
	{{else}}
		{{$user = .User.ID}}
	{{end}}
{{end}}

{{$stats := sdict}}
{{with (dbGet $user "stats")}}
	{{$stats = sdict .Value}}
{{end}}

{{$force:= $stats.Get "force"}}
{{$endurance := $stats.Get "endurance"}}
{{$agi:=$stats.Get "agi"}}
{{$preci:=$stats.Get "preci"}}
{{$intel:=$stats.Get "intelligence"}}
{{$karma:=$stats.Get "karma"}}

{{$iforce:=$stats.Get "i_force"}}
{{$iendu:=$stats.Get "i_endu"}}
{{$iagi:=$stats.Get "i_agi"}}
{{$ipreci:=$stats.Get "i_preci"}}
{{$iintel:=$stats.Get "i_intel"}}

{{if .CmdArgs}}
{{if eq (len .CmdArgs) 2}}
	{{if eq (index .CmdArgs 1) "-stats"}}
**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{$force}}
	:white_small_square: Endurance : {{$endurance}}
	:white_small_square: Agilité : {{$agi}}
	:white_small_square: Précision : {{$preci}}
	:white_small_square: Intelligence : {{$intel}}
	:white_small_square: Karma : {{$karma}}

		{{else if eq (index .CmdArgs 1) "-implant"}}
**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{$iforce}}
	:white_small_square: Endurance : {{$iendu}}
	:white_small_square: Agilité : {{$iagi}}
	:white_small_square: Précision : {{$ipreci}}
	:white_small_square: Intelligence : {{$iintel}}
	{{end}}

{{else if eq (len .CmdArgs) 1}}
	{{if eq (index .CmdArgs 0) "-stats"}}
	**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{$force}}
	:white_small_square: Endurance : {{$endurance}}
	:white_small_square: Agilité : {{$agi}}
	:white_small_square: Précision : {{$preci}}
	:white_small_square: Intelligence : {{$intel}}
	:white_small_square: Karma : {{$karma}}

		{{else if eq (index .CmdArgs 0) "-implant"}}
	**Implant de <@{{$user}}>** :
	:white_small_square: Force : {{$iforce}}
	:white_small_square: Endurance : {{$iendu}}
	:white_small_square: Agilité : {{$iagi}}
	:white_small_square: Précision : {{$ipreci}}
	:white_small_square: Intelligence : {{$iintel}}

	{{end}}

{{else}}
	**Statistiques de <@{{$user}}>**
		:white_small_square: Force : {{$force}}
		:white_small_square: Endurance : {{$endurance}}
		:white_small_square: Agilité : {{$agi}}
		:white_small_square: Précision : {{$preci}}
		:white_small_square: Intelligence : {{$intel}}
		:white_small_square: Karma : {{$karma}}

	**Implants de <@{{$user}}>** :
		:white_small_square: Force : {{$iforce}}
		:white_small_square: Endurance : {{$iendu}}
		:white_small_square: Agilité : {{$iagi}}
		:white_small_square: Précision : {{$ipreci}}
		:white_small_square: Intelligence : {{$iintel}}
	{{end}}

{{else}}
**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{$force}}
	:white_small_square: Endurance : {{$endurance}}
	:white_small_square: Agilité : {{$agi}}
	:white_small_square: Précision : {{$preci}}
	:white_small_square: Intelligence : {{$intel}}
	:white_small_square: Karma : {{$karma}}

**Implants de <@{{$user}}>** :
	:white_small_square: Force : {{$iforce}}
	:white_small_square: Endurance : {{$iendu}}
	:white_small_square: Agilité : {{$iagi}}
	:white_small_square: Précision : {{$ipreci}}
	:white_small_square: Intelligence : {{$iintel}}
{{end}}
{{deleteTrigger 1}}
