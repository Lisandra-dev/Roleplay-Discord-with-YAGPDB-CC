{{$user:= .Member.Nick}}
  {{if eq (len $user) 0}}
    {{$user = .User.Username}}
  {{end}}

{{$d:= (randInt 1 10)}}
{{$real:= $d}}
{{$arg1:=""}}
{{$arg2:= ""}}
{{$res := ""}}
{{$m := $d}}
{{$i := $d}}
{{$x := $i}}
{{$y := $m}}
{{$v := (joinStr "" "(" (toString $real) ")")}}
{{$comm:=""}}

{{$seuil := (toInt 8)}}
{{$idb := (toInt 0)}}

{{$arg := 0}}
{{$argm := 0}}
{{$c := ""}}

{{$manuel := reFind `s[1-9]` .Message.Content}}
{{$force:=reFindAllSubmatches `(force|Force)` .Message.Content}}
{{$agi := reFindAllSubmatches `(agilité|Agilité|agi|Agi)` .Message.Content}}
{{$endu := reFindAllSubmatches `(Endurance|endurance|endu|Endu)` .Message.Content}}
{{$preci := reFindAllSubmatches `(Précision|précision|préci|Préci)` .Message.Content}}
{{$intell := reFindAllSubmatches `(Intelligence|intelligence|intell|Intell|intel|Intel)` .Message.Content}}
{{$karma := reFindAllSubmatches `(Karma|karma)` .Message.Content}}

{{if $manuel}}
	{{$seuil = (toInt (joinStr "" (split $manuel "s")))}}
	{{$idb = (toInt 0)}}
{{else if $force}}
	{{$seuil = (toInt (dbGet .User.ID "force").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_force").Value)}}
{{else if $agi}}
	{{$seuil = (toInt (dbGet .User.ID "agi").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_agi").Value)}}
{{else if $endu}}
	{{$seuil = (toInt (dbGet .User.ID "endurance").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_endu").Value)}}
{{else if $preci}}
	{{$seuil = (toInt (dbGet .User.ID "preci").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_preci").Value)}}
{{else if $intell}}
	{{$seuil = (toInt (dbGet .User.ID "intelligence").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_intel").Value)}}
{{else if $karma}}
	{{$seuil = (toInt (dbGet .User.ID "karma").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_karma").Value)}}
{{else}}
	{{$seuil = (toInt 8)}}
	{{$idb = (toInt 0)}}
{{end}}

{{$imp := ""}}

{{if le $idb 1}}
	{{$imp = "Implant"}}
{{else}}
	{{$imp = "Implants"}}
{{end}}

{{$mimp := $idb}}


{{if .CmdArgs}}
	{{if $manuel}}
		{{$c = split (joinStr " " .CmdArgs) $manuel}}
	{{else}}
		{{$c = (joinStr " " .CmdArgs)}}
	{{end}}

	{{if lt (toFloat (len .CmdArgs)) (toFloat 2)}}
	  {{if ne (toInt (index .CmdArgs 0)) (toInt 0)}}
	    {{$i := (toInt  (index .CmdArgs 0))}}
			{{$x := $i}}
			{{if gt (toInt (index .CmdArgs 0)) (toInt 0)}}
				{{$arg = 1}}
				{{if eq (toInt (index .CmdArgs 0)) (toInt 1) }}
					{{$arg1 = "Pénalité"}}
				{{else}}
					{{$arg1 = "Pénalités"}}
				{{end}}
			{{else}}
				{{$i = mult $i (toInt -1)}}
				{{$x = mult $x (toInt -1)}}
				{{$arg = 0}}
				{{$arg1 = "Bonus"}}
			{{end}}

			{{if eq $arg 0}}
				{{$i = add $i $idb}}
				{{$d = sub $d $i}}
			{{else}}
				{{$d = add $i $d}}
				{{$d = sub $d $idb}}
			{{end}}

			{{if gt $d (toInt 10)}}
				{{$d = toInt 10}}
			{{else if lt $d (toInt 0)}}
				{{$d = (toInt 0)}}
			{{end}}

			{{$res = (joinStr "" "Dé : " (toString $d) " " $v " |  " (toString $x) " : " $arg1 " | " $imp " : " $idb " | Seuil : " $seuil )}}
  	{{else}}
			{{$d = sub $d $idb}}
			{{$res = (joinStr "" "Dé : " (toString $d) " " $v " | " $imp " : " $idb " | Seuil : " $seuil )}}
		{{end}}
	{{else}}
		{{if ne (toInt (index .CmdArgs 0)) (toInt 0) }}
			{{$i = (toInt  (index .CmdArgs 0))}}
			{{$x := $i}}
			{{if gt (toInt (index .CmdArgs 0)) (toInt 0)}}
				{{$arg = 1}}
				{{if eq (toInt (index .CmdArgs 0)) (toInt 1)}}
					{{$arg1 = "pénalité"}}
				{{else}}
					{{$arg1 = "pénalités"}}
				{{end}}
			{{else}}
				{{$arg = 0}}
				{{$i = mult $i (toInt -1)}}
				{{$x = mult $x (toInt -1)}}
				{{$arg1 = "Bonus"}}
			{{end}}

			{{if ne (toInt (index .CmdArgs 1)) (toInt 0)}}
				{{$m := (toInt (index .CmdArgs 1))}}
				{{$y := $m}}
				{{if gt $m (toInt 0)}}
					{{$argm = 1}}
					{{if eq $m (toInt 1) }}
						{{$arg2 = "Pénalité"}}
					{{else}}
						{{$arg2 = "Pénalités"}}
					{{end}}
				{{else}}
					{{$argm = 0}}
					{{$m = mult $m (toInt -1)}}
					{{$y = mult $y (toInt -1)}}
					{{$arg2 = "Bonus"}}
				{{end}}

				{{if eq $arg 0}}
					{{$i = add $i $idb}}
					{{$d = sub $d $i}}
					{{$mimp = $i}}
					{{if eq $argm 0}}
						{{$d = sub $d $m}}
					{{else}}
						{{$d = add $d $m}}
					{{end}}
				{{else}}
					{{$d = add $d $i}}
					{{if eq $argm 0}}
						{{$m = add $m $idb}}
						{{$d = sub $d $m }}
						{{$mimp = $m}}
					{{else}}
						{{$d = add $d $m}}
						{{$d = sub $d $idb}}
						{{$mimp = $d}}
					{{end}}
				{{end}}


				{{if gt $d (toInt 10)}}
					{{$d = toInt 10}}
				{{else if lt $d (toInt 0)}}
					{{$d = (toInt 0)}}
				{{end}}


				{{$res = (joinStr "" "Dé : " $d " " $v " | " $arg1 " : " $x " | " $arg2 " : " $y " | " $imp " : " $idb " | Seuil : " $seuil )}}
			{{else}}

						{{if eq $arg 0}}
							{{$i = add $i $idb}}
							{{$d = sub $d $i}}
							{{$mimp = $i}}
						{{else}}
							{{$d = add $i $d}}
							{{$d = sub $d $idb}}
							{{$mimp = $idb}}
						{{end}}

						{{if gt $d (toInt 10)}}
							{{$d = toInt 10}}
						{{else if lt $d (toInt 0)}}
							{{$d = (toInt 0)}}
						{{end}}

				{{$res = (joinStr "" "Dé : " (toString $d) " " $v " |  " $arg1 " : " (toString $x) " | " $imp " : " $idb " | Seuil : " $seuil)}}
			{{end}}
		{{else}}
			{{$d = sub $d $idb}}
			{{$res = (joinStr "" "Dé : " (toString $d) " " $v " | " $imp " : " $idb " | Seuil : " $seuil )}}
		{{end}}
	{{end}}

	{{if lt (toFloat (len .CmdArgs)) (toFloat 2)}}
		{{if ne (toInt 0) (toInt (index .CmdArgs 0)) }}
			{{$comm = ""}}
		{{else}}
			{{$comm = (joinStr " " $c ": ")}}
		{{end}}
	{{else}}
		{{if ne (toInt 0) (toInt (index .CmdArgs 0)) }}
			{{if ne (toFloat (index .CmdArgs 1)) (toFloat 0)}}
				{{if eq (toFloat 2) (toFloat (len .CmdArgs))}}
					{{$comm = ""}}
				{{else}}
					{{if $manuel}}
							{{$comm =  joinStr "" (slice .CmdArgs 2)}}
							{{$comm = joinStr " " (split $comm $manuel)}}
							{{if ne $comm " "}}
								{{$comm = joinStr " " (split $comm $manuel) ": " }}
							{{end}}
					{{else}}
						{{$comm = joinStr " " (slice .CmdArgs 2) ": "}}
					{{end}}
				{{end}}
			{{else}}
				{{if $manuel}}
						{{$comm =  joinStr "" (slice .CmdArgs 1)}}
						{{$comm = joinStr " " (split $comm $manuel)}}
						{{if ne $comm " "}}
							{{$comm = joinStr " " (split $comm $manuel) ": "}}
						{{end}}
					{{else}}
						{{$comm = joinStr " " (slice .CmdArgs 1) ": " }}
					{{end}}
			{{end}}
		{{else}}
			{{$comm = (joinStr " " $c ": ") }}
		{{end}}
	{{end}}
{{else}}
	{{$res = (joinStr "" "Dé : " (toString $d) " " $v " | " $imp " : " $idb " | Seuil : " $seuil )}}
{{end}}

{{$r := ""}}

{{$urc := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "**Ultra critique :** Votre cible regagne 8 PV *(+1 si module/PSI)* et obtient un bonus de votre choix.\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}

	{{$rc := cembed
			"description" (joinStr "" "**" $user "** ▬ " $comm "**Réussite critique: ** Votre cible regagne 8 PV *(+1 si module/PSI).*\n"
			"<:next:723131844643651655> *[" $res "]*" )
			"color" 0xEFA3EA }}

{{$r := ""}}

{{if eq $seuil (toInt 2)}}
	{{$r = cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "**Réussite** : Votre cible regagne 7 PV *(+1 si module/PSI).*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}

{{else if eq $seuil (toInt 3)}}
	{{$r = cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Réussite* : Votre cible regagne 6 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}

{{else if eq $seuil (toInt 4)}}
	{{$r = cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Réussite* : Votre cible regagne 5 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}

{{else if eq $d (toInt 5)}}
	{{$r = cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Réussite* : Votre cible regagne 4 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}

{{else if eq $seuil (toInt 6)}}
	{{$r = cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Réussite* : Votre cible regagne 3 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}

{{else if eq $seuil (toInt 7)}}
	{{$r = cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Réussite* : Votre cible regagne 2 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}

{{else if eq $seuil (toInt 8)}}
	{{$r = cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Réussite* : Votre cible regagne 1 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}
{{end}}

{{$echec := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "**Echec du soin ...**\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}

{{$ec := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "**Echec critique du soin :** Votre cible gagne une altération.\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}

{{if eq $d (toInt 0)}}
	{{sendMessage nil $urc}}
{{else if eq $d (toInt 1)}}
	{{sendMessage nil $rc}}
{{else if le $d $seuil}}
	{{if ge $mimp (toInt 1)}}
		{{if eq $d $seuil}}
			{{sendMessage nil $echec}}
		{{else}}
			{{sendMessage nil $r}}
		{{end}}
	{{else}}
		{{sendMessage nil $r}}
	{{end}}
{{else if or (gt $d $seuil) (lt $d (toInt 9)) }}
	{{sendMessage nil $echec}}
{{else if eq $d (toInt 10)}}
	{{sendMessage nil $ec}}
{{else if eq $d (toInt 9)}}
	{{sendMessage nil $echec}}
{{else}}
{{end}}

{{deleteTrigger 1}}