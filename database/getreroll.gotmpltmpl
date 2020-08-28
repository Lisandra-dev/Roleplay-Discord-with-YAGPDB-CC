{{if .CmdArgs}}
	{{$name := toRune (lower (index .CmdArgs 0))}}
	{{$user :=.User.ID }}
	{{range $name}}
		{{- $user = (add $user .)}}
	{{- end}}

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

	{{if eq (len .CmdArgs ) 2}}
		{{if eq (index .CmdArgs 1) "-stats"}}

**Statistiques de {{index .CmdArgs 0}} : **
	:white_small_square: Force : {{$force}}
	:white_small_square: Endurance : {{$endurance}}
	:white_small_square: Agilité : {{$agi}}
	:white_small_square: Précision : {{$preci}}
	:white_small_square: Intelligence : {{$intel}}
	:white_small_square: Karma : {{$karma}}

		{{else if eq (index .CmdArgs 1) "-implant"}}

**Implant de {{index .CmdArgs 0}}** :
	:white_small_square: Force : {{$iforce}}
	:white_small_square: Endurance : {{$iendu}}
	:white_small_square: Agilité : {{$iagi}}
	:white_small_square: Précision : {{$ipreci}}
	:white_small_square: Intelligence : {{$iintel}}

		{{else}}
**Statistiques de {{index .CmdArgs 0}} : **
	:white_small_square: Force : {{$force}}
	:white_small_square: Endurance : {{$endurance}}
	:white_small_square: Agilité : {{$agi}}
	:white_small_square: Précision : {{$preci}}
	:white_small_square: Intelligence : {{$intel}}
	:white_small_square: Karma : {{$karma}}

**Implants de {{index .CmdArgs 0}}** :
	:white_small_square: Force : {{$iforce}}
	:white_small_square: Endurance : {{$iendu}}
	:white_small_square: Agilité : {{$iagi}}
	:white_small_square: Précision : {{$ipreci}}
	:white_small_square: Intelligence : {{$iintel}}
		{{end}}
	{{else}}
	**Statistiques de {{index .CmdArgs 0}} : **
		:white_small_square: Force : {{$force}}
		:white_small_square: Endurance : {{$endurance}}
		:white_small_square: Agilité : {{$agi}}
		:white_small_square: Précision : {{$preci}}
		:white_small_square: Intelligence : {{$intel}}
		:white_small_square: Karma : {{$karma}}

	**Implants de {{index .CmdArgs 0}}** :
		:white_small_square: Force : {{$iforce}}
		:white_small_square: Endurance : {{$iendu}}
		:white_small_square: Agilité : {{$iagi}}
		:white_small_square: Précision : {{$ipreci}}
		:white_small_square: Intelligence : {{$iintel}}

	{{end}}
{{else}}
	**ERREUR** : Vous devez indiquer votre personnage !
{{end}}
{{deleteTrigger 1}}
