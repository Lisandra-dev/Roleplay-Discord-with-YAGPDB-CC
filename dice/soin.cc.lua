{{$user:= .Member.Nick}}
  {{if eq (len $user) 0}}
    {{$user = .User.Username}}
  {{end}}
{{$d:= (randInt 1 10)}}
{{$real:=$d}}
{{$arg1:=""}}
{{$arg2:= ""}}
{{$res := ""}}
{{$m := $d}}
{{$i := $d}}
{{$v := (joinStr "" "(" (toString $real) ")")}}
{{$comm:=""}}


{{$user:= .Member.Nick}}
  {{if eq (len $user) 0}}
    {{$user = .User.Username}}
  {{end}}

{{if .CmdArgs}}
	{{if lt (toFloat (len .CmdArgs)) (toFloat 2)}}
	  {{if ne (toInt (index .CmdArgs 0)) (toInt 0)}}
	    {{$i = (toInt  (index .CmdArgs 0))}}
			{{$d = add $d $i}}
			{{if gt $d (toInt 10)}}
				{{$d = (toInt 10)}}
			{{end}}
			{{if gt (toInt (index .CmdArgs 0)) (toInt 0)}}
				{{if eq (toInt (index .CmdArgs 0)) (toInt 1) }}
					{{$arg1 = "pénalité"}}
				{{else}}
					{{$arg1 = "pénalités"}}
				{{end}}
			{{else}}
				{{$i = mult $i (toInt -1)}}
				{{if ge $i (toInt 2)}}
					{{$arg1 = "implants"}}
				{{else}}
					{{$arg1 = "implant"}}
				{{end}}
			{{end}}
		{{$res = (joinStr "" "Dé : " (toString $d) " " $v " |  " (toString $i) " " $arg1)}}
  	{{else}}
		{{$res = (joinStr " " "Dé :" (toString $d))}}
	{{end}}

{{else}}
	{{if ne (toInt (index .CmdArgs 0)) (toInt 0) }}
			{{$i = (toInt  (index .CmdArgs 0))}}
			{{$d = add $d $i}}
			{{if gt (toInt (index .CmdArgs 0)) (toInt 0)}}
				{{if eq (toInt (index .CmdArgs 0)) (toInt 1)}}
					{{$arg1 = "pénalité"}}
				{{else}}
					{{$arg1 = "pénalités"}}
				{{end}}
			{{else}}
				{{$i = mult $i (toInt -1)}}
				{{if ge $i (toInt 2)}}
					{{$arg1 = "implants"}}
				{{else}}
					{{$arg1 = "implant"}}
				{{end}}
			{{end}}
		{{if ne (toInt (index .CmdArgs 1)) (toInt 0)}}
			{{$m := (toInt (index .CmdArgs 1))}}

			{{$d = add $d $m}}
			{{if gt $d (toInt 10)}}
				{{$d = (toInt 10)}}
			{{end}}
			{{if gt $m (toInt 0)}}
				{{if eq $m (toInt 1) }}
					{{$arg2 = "pénalité"}}
				{{else}}
					{{$arg2 = "pénalités"}}
				{{end}}
			{{else}}
				{{$m = mult $m (toInt -1)}}
				{{if ge $m (toInt 2)}}
					{{$arg2 = "implants"}}
				{{else}}
					{{$arg2 = "implant"}}
				{{end}}
			{{end}}

			{{$res = (joinStr "" "Dé : " (toString $d) " " $v  " |  " (toString $i) " " $arg1 " et " (toString $m) " " $arg2)}}
		{{else}}
		{{$res = (joinStr "" "Dé : " (toString $d) " " $v " |  " (toString $i) " " $arg1)}}
		{{end}}
		{{else}}
			{{$res = (joinStr " " "Dé :" (toString $d)) }}
		{{end}}
	{{end}}

	{{if lt (toFloat (len .CmdArgs)) (toFloat 2)}}
		{{if ne (toInt 0) (toInt (index .CmdArgs 0)) }}
			{{$comm = ""}}
		{{else}}
			{{$comm = (joinStr " " .CmdArgs ": ")}}
		{{end}}
	{{else}}
		{{if ne (toInt 0) (toInt (index .CmdArgs 0)) }}
			{{if ne (toFloat (index .CmdArgs 1)) (toFloat 0)}}
				{{if eq (toFloat 2) (toFloat (len .CmdArgs))}}
					{{$comm = ""}}
				{{else}}
					{{$comm = (joinStr " " (slice .CmdArgs 2) ": ") }}
				{{end}}
			{{else}}
				{{$comm = (joinStr " " (slice .CmdArgs 1) ": ") }}
			{{end}}
		{{else}}
			{{$comm = (joinStr " " .CmdArgs ": ") }}
		{{end}}

	{{end}}


{{else}}
	{{$res = (joinStr " " "Dé :" (toString $d)) }}
{{end}}

{{if le $d (toInt 0)}}
	{{$embed := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "**Ultra critique :** Votre cible regagne 8 PV *(+1 si module/PSI)* et obtient un bonus de votre choix.\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}
	{{sendMessage nil $embed}}
	{{else if eq $d (toInt 1)}}
		{{$embed := cembed
			"description" (joinStr "" "**" $user "** ▬ " $comm "**Réussite critique: ** Votre cible regagne 8 PV *(+1 si module/PSI).*\n"
			"<:next:723131844643651655> *[" $res "]*" )
			"color" 0xEFA3EA }}
		{{sendMessage nil $embed}}
	{{else if eq $d (toInt 2)}}
		{{$embed := cembed
			"description" (joinStr "" "**" $user "** ▬ " $comm "**Réussite** : Votre cible regagne 7 PV *(+1 si module/PSI).*\n"
			"<:next:723131844643651655> *[" $res "]*" )
			"color" 0xEFA3EA }}
		{{sendMessage nil $embed}}
	{{else if eq $d (toInt 3)}}
	{{$embed := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Si votre seuil le permet* : Votre cible regagne 6 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}
			{{sendMessage nil $embed}}
	{{else if eq $d (toInt 4)}}
	{{$embed := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Si votre seuil le permet* : Votre cible regagne 5 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}
			{{sendMessage nil $embed}}
	{{else if eq $d (toInt 5)}}
	{{$embed := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Si votre seuil le permet* : Votre cible regagne 4 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}
			{{sendMessage nil $embed}}
	{{else if eq $d (toInt 6)}}
	{{$embed := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Si votre seuil le permet* : Votre cible regagne 3 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}
			{{sendMessage nil $embed}}
	{{else if eq $d (toInt 7)}}
	{{$embed := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Si votre seuil le permet* : Votre cible regagne 2 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}
	{{sendMessage nil $embed}}
	{{else if eq $d (toInt 8)}}
	{{$embed := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Si votre seuil le permet* : Votre cible regagne 1 PV *(+1 si module/PSI)*\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}
			{{sendMessage nil $embed}}
	{{else if eq $d (toInt 9)}}
	{{$embed := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "**Echec du soin ...**\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}
			{{sendMessage nil $embed}}
	{{else if eq $d (toInt 10)}}
	{{$embed := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "**Echec critique du soin :** Votre cible gagne une altération.\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0xEFA3EA }}
			{{sendMessage nil $embed}}
	{{end}}

{{deleteTrigger 1}}
