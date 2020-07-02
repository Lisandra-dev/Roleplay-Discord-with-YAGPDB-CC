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
	{{if eq (len .CmdArgs ) 1}}
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
	:white_small_square: Karma : {{$ikarma}}

		{{else}}
			**Usage** : `$getchar -(stats|implant)`
		{{end}}

	{{else if eq (len .CmdArgs) 3}}
		{{if and (eq (index .CmdArgs 0) "-all") (eq (index .CmdArgs 1) "-ID")}}
			{{$user = (toInt (index .CmdArgs 2))}}

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
	:white_small_square: Karma : {{$ikarma}}

		{{else if and (eq (index .CmdArgs 0) "-stats") (eq (index .CmdArgs 1) "-ID")}}
			{{$user = (toInt (index .CmdArgs 2))}}

**Statistiques de <@{{$user}}>**
	:white_small_square: Force : {{$force}}
	:white_small_square: Endurance : {{$endurance}}
	:white_small_square: Agilité : {{$agi}}
	:white_small_square: Précision : {{$preci}}
	:white_small_square: Intelligence : {{$intel}}
	:white_small_square: Karma : {{$karma}}

		{{else if and (eq (index .CmdArgs 0) "-implant") (eq (index .CmdArgs 1) "-ID")}}
			{{$user = (toInt (index .CmdArgs 2))}}

**Implants de <@{{$user}}>** :
	:white_small_square: Force : {{$iforce}}
	:white_small_square: Endurance : {{$iendu}}
	:white_small_square: Agilité : {{$iagi}}
	:white_small_square: Précision : {{$ipreci}}
	:white_small_square: Intelligence : {{$iintel}}
	:white_small_square: Karma : {{$ikarma}}

		{{else}}
**Arguments non reconnus**
	**Usage** : `$getchar -(all|implant|stats) (-ID <ID>)`
		{{end}}
	{{else if or (gt (len .CmdArgs) 3) (le (len .CmdArgs) 2)}}
		{{if gt (len .CmdArgs) 3}}
**Trop d'argument indiqué**
	**Usage** : `$getchar -(all|implant|stats) (-ID <ID>)`
		{{else}}
**Pas assez d'argument** : Vous devez indiqué L'ID de la personne ciblée.
	**Usage** : `$getchar -(all|implant|stats) (-ID <ID>)`
		{{end}}
	{{else}}
**Pas assez d'argument**
	**Usage** : `$getchar -(all|implant|stats) (-ID <ID>)`
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
	:white_small_square: Karma : {{$ikarma}}
{{end}}
