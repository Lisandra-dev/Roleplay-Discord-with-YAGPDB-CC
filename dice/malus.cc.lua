{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id := .User.ID }}

{{if $name}}
	{{$user = $name}}
	{{$idperso := (toRune (lower $name))}}
	{{$id = ""}}
	{{range $idperso}}
		{{- $id = (print $id .)}}
	{{- end}}
	{{$id = (toInt $id)}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$user = title $user}}

{{$stats := sdict}}
{{with (dbGet $id "stats")}}
	{{$stats = sdict .Value}}
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
{{$agi := reFindAllSubmatches `(agilit(é|e)|Agilit(é|e)|\dagi|\dAgi)` .Message.Content}}
{{$endu := reFindAllSubmatches `(Endurance|endurance|\dendu|\dEndu)` .Message.Content}}
{{$preci := reFindAllSubmatches `(Pr(é|e)cision|pr(é|e)cision|pr(é|e)ci|Pr(é|e)ci)` .Message.Content}}
{{$intell := reFindAllSubmatches `(Intelligence|intelligence|intell|Intell|intel|Intel)` .Message.Content}}
{{$karma := reFindAllSubmatches `(Karma|karma)` .Message.Content}}

{{if $manuel}}
	{{$seuil = (toInt (joinStr "" (split $manuel "s")))}}
	{{$idb = (toInt 0)}}
{{else if $force}}
	{{$seuil = (toInt ($stats.Get "force"))}}
	{{$idb = (toInt ($stats.Get "i_force"))}}
{{else if $agi}}
	{{$seuil = (toInt ($stats.Get "agi"))}}
	{{$idb = (toInt ($stats.Get "i_agi"))}}
{{else if $endu}}
	{{$seuil = (toInt ($stats.Get "endurance"))}}
	{{$idb = (toInt ($stats.Get "i_endu"))}}
{{else if $preci}}
	{{$seuil = (toInt ($stats.Get "preci"))}}
	{{$idb = (toInt ($stats.Get "i_preci"))}}
{{else if $intell}}
	{{$seuil = (toInt ($stats.Get "intelligence"))}}
	{{$idb = (toInt ($stats.Get "i_intel"))}}
{{else if $karma}}
	{{$seuil = (toInt ($stats.Get "karma"))}}
	{{$idb = (toInt 0)}}
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
	{{$c = joinStr " " .CmdArgs}}
	{{if $name}}
		{{$c = joinStr " " (split $c $name) }}
		{{$c = joinStr " " (split $c "#")}}
	{{end}}

	{{if $manuel}}
		{{$c = joinStr " " (split $c $manuel) }}
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
			{{$comm = (joinStr " " " ▬" $c ": ")}}
		{{end}}
	{{else}}
		{{if ne (toInt 0) (toInt (index .CmdArgs 0)) }}
			{{if ne (toFloat (index .CmdArgs 1)) (toFloat 0)}}
				{{if eq (toFloat 2) (toFloat (len .CmdArgs))}}
					{{$comm = ""}}
				{{else}}
					{{$comm =  joinStr " " (slice .CmdArgs 2) }}
					{{if $manuel}}
							{{$comm = joinStr " " (split $comm $manuel)}}
					{{end}}
					{{if $name}}
						{{$comm = joinStr " " (split $comm $name)}}
						{{$comm = joinStr " " (split $comm "#")}}
					{{end}}
					{{if ne $comm " "}}
						{{$comm = joinStr " " " ▬" $comm ": "}}
					{{else}}
						{{$comm = joinStr " " $comm}}
					{{end}}
				{{end}}
			{{else}}
				{{$comm =  joinStr " " (slice .CmdArgs 1)}}
				{{if $manuel}}
					{{$comm = joinStr " " (split $comm $manuel) }}
				{{end}}
				{{if $name}}
					{{$comm = joinStr "" (split $comm $name)}}
					{{$comm = joinStr " " (split $comm "#")}}
				{{end}}
				{{if ne $comm " "}}
					{{$comm = joinStr " " " ▬" $comm ": "}}
				{{else}}
					{{$comm = joinStr " " $comm}}
				{{end}}
			{{end}}
		{{else}}
			{{$comm = (joinStr " " " ▬" $c ": ") }}
		{{end}}
	{{end}}
{{else}}
	{{$res = (joinStr "" "Dé : " (toString $d) " " $v " | " $imp " : " $idb " | Seuil : " $seuil )}}
{{end}}

{{$img := "https://i.imgur.com/fuHvIUn.png"}}

{{$urc := cembed
	"author" (sdict "name" $user "icon_url" $img)
	"description" (joinStr "" "**Ultra critique**" $comm "\n Votre cible a une pénalité de +4 à son dé.\n"
	"<:next:723131844643651655> *[" $res "]*" )
	"color" 0x7E2257 }}

{{$rc := cembed
	"author" (sdict "name" $user "icon_url" $img)
	"description" (joinStr "" "**Réussite critique**" $comm "\n Votre cible a une pénalité de +3 à son dé.\n"
	"<:next:723131844643651655> *[" $res "]*" )
	"color" 0x7E2257 }}

{{$r := ""}}

{{if eq $d (toInt 2)}}
	{{$r = cembed
		"author" (sdict "name" $user "icon_url" $img)
		"description" (joinStr "" "**Réussite**" $comm "\n Votre cible a une pénalité de +3 à son dé.\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0x7E2257 }}
{{else if and (le $d (toInt 3)) (ge $d (toInt 5))}}
	{{$r = cembed
		"author" (sdict "name" $user "icon_url" $img)
		"description" (joinStr "" "**Réussite**" $comm "\n Votre cible a une pénalité de +2 à son dé.\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0x7E2257 }}
{{else if and (le $d (toInt 8)) (ge $d (toInt 6))}}
	{{$r = cembed
		"author" (sdict "name" $user "icon_url" $img)
		"description" (joinStr "" "**Réussite**" $comm "\n Votre cible a une pénalité de +1 à son dé.\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0x7E2257 }}
{{end}}

{{$echec := cembed
	"author" (sdict "name" $user "icon_url" $img)
	"description" (joinStr "" "**Echec de l'altération**" $comm "\n"
	"<:next:723131844643651655> *[" $res "]*" )
	"color" 0x7E2257 }}

{{$ec := cembed
	"author" (sdict "name" $user "icon_url" $img)
	"description" (joinStr "" $comm "**Echec critique de l'altération** " $comm " Votre cible a un bonus de -2 à son dé.\n"
	"<:next:723131844643651655> *[" $res "]*" )
	"color" 0x7E2257}}

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
