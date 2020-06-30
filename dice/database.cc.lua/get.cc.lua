{{$seuil:=(toInt 0)}}
{{$implant := (toInt 0)}}

{{$force:=reFindAllSubmatches `(force|Force)` .Message.Content}}
{{$agi := reFindAllSubmatches `(agilité|Agilité|agi|Agi)` .Message.Content}}
{{$endu := reFindAllSubmatches `(Endurance|endurance|endu|Endu)` .Message.Content}}
{{$preci := reFindAllSubmatches `(Précision|précision|préci|Préci)` .Message.Content}}
{{$intell := reFindAllSubmatches `(Intelligence|intelligence|intell|Intell|intel|Intel)` .Message.Content}}
{{$karma := reFindAllSubmatches `(Karma|karma)` .Message.Content}}

{{if $force}}
	{{$seuil = (toInt (dbGet .User.ID "force").Value)}}
	{{$implant = (toInt (dbGet .User.ID "i_force").Value)}}
	Vous avez :
	:white_small_square: Force : {{$seuil}}
	:white_small_square: Implant de force : {{$implant}}

{{else if $agi}}
	{{$seuil = (toInt (dbGet .User.ID "agi").Value)}}
	{{$implant = (toInt (dbGet .User.ID "i_agi").Value)}}
	Vous avez :
	:white_small_square: Agilité : {{$seuil}}
	:white_small_square: Implant d'agilité : {{$implant}}

{{else if $endu}}
	{{$seuil = (toInt (dbGet .User.ID "endurance").Value)}}
	{{$implant = (toInt (dbGet .User.ID "i_endu").Value)}}
	Vous avez :
	:white_small_square: Endurance : {{$seuil}}
	:white_small_square: Implant d'endurance : {{$implant}}

{{else if $preci}}
	{{$seuil = (toInt (dbGet .User.ID "preci").Value)}}
	{{$implant = (toInt (dbGet .User.ID "i_preci").Value)}}
	Vous avez :
	:white_small_square: Précision : {{$seuil}}
	:white_small_square: Implant de précision : {{$implant}}

{{else if $intell}}
	{{$seuil = (toInt (dbGet .User.ID "intelligence").Value)}}
	{{$implant = (toInt (dbGet .User.ID "i_intel").Value)}}
	Vous avez :
	:white_small_square: Intelligence : {{$seuil}}
	:white_small_square: Implant d'intelligence : {{$implant}}

{{else if $karma}}
	{{$seuil = (toInt (dbGet .User.ID "karma").Value)}}
	{{ $implant = (toInt (dbGet .User.ID "i_karma").Value)}}
	Vous avez :
	:white_small_square: Karma : {{$seuil}}
	:white_small_square: Implant de Karma : {{$implant}}

{{else}}
Vous avez :
	:white_small_square: Force : {{(dbGet .User.ID "force").Value}}
	:white_small_square: Endurance : {{(dbGet .User.ID "endurance").Value}}
	:white_small_square: Agilité : {{(dbGet .User.ID "agi").Value}}
	:white_small_square: Précision : {{(dbGet .User.ID "preci").Value}}
	:white_small_square: Intelligence : {{(dbGet .User.ID "intelligence").Value}}
	:white_small_square: Karma : {{(dbGet .User.ID "karma").Value}}

Avec, comme implants :
	:white_small_square: Force : {{(dbGet  .User.ID "i_force").Value}}
	:white_small_square: Endurance : {{(dbGet .User.ID "i_endu").Value}}
	:white_small_square: Agilité : {{(dbGet .User.ID "i_agi").Value}}
	:white_small_square: Précision : {{(dbGet .User.ID "i_preci").Value}}
	:white_small_square: Intelligence : {{(dbGet .User.ID "i_intel").Value}}
	:white_small_square: Karma : {{(dbGet .User.ID "i_karma").Value}}
{{end}}
