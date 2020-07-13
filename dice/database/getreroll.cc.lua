{{if .CmdArgs}}
	{{$name := toRune (lower (index .CmdArgs 0))}}
	{{$user :=""}}
	{{range $name}}
	{{- $user = (print $user .)}}
	{{- end}}
	{{$user = (toInt $user)}}

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
