{{$user:=.Member.Nick}}
{{if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$d:= (randInt 1 10)}}

{{$arg1:=""}}
{{$arg2:= ""}}
{{$v:=$d}}
{{$arg := 0}}
{{$argm := 0}}
{{$c:=""}}


{{$msg := ""}}

{{$seuil := (toInt 8)}}
{{$idb := (toInt 0)}}

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


{{ if .CmdArgs}}
	{{$c = joinStr " " .CmdArgs}}
	{{if $manuel}}
		{{$c = joinStr " " (split $c $manuel) }}
	{{end}}

	{{if ne (toFloat (index .CmdArgs 0)) (toFloat 0)}}
		{{$i:=(toInt (index .CmdArgs 0)) }}
		{{$x := $i}}
		{{if gt (toInt (index .CmdArgs 0)) (toInt 0)}}
			{{$arg = 1}}
			{{if eq (toInt (index .CmdArgs 0)) (toInt 1)}}
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
			{{$mimp = $i}}
		{{else}}
			{{$d = add $i $d}}
			{{if ge (toFloat (len .CmdArgs)) (toFloat 2 )}}
				{{if eq (toInt (index .CmdArgs 1)) 0}}
					{{$d = sub $d $idb}}
					{{$mimp = $idb}}
				{{end}}
			{{end}}
		{{end}}

		{{if gt $d (toInt 10)}}
			{{$d = toInt 10}}
		{{else if lt $d (toInt 0)}}
			{{$d = (toInt 0)}}
		{{end}}

		{{if eq $d (toInt 0)}}
			{{$msg ="**Ultra critique !** "}}
		{{else if eq $d (toInt 1)}}
			{{$msg = "**Réussite critique !** "}}
		{{else if le $d $seuil}}
				{{if ge $mimp (toInt 1)}}
					{{if eq $d $seuil}}
						{{$msg = "**Echec** "}}
					{{else}}
						{{$msg = "**Réussite** "}}
					{{end}}
				{{else}}
					{{$msg = "**Réussite** "}}
				{{end}}
		{{else if or (gt $d $seuil) (lt $d (toInt 9))}}
			{{$msg = "**Echec** "}}
		{{else if eq $d (toInt 10)}}
			{{$msg = "**Echec critique !** "}}
		{{else if eq $d (toInt 9)}}
			{{$msg = "**Echec...** "}}
		{{else}}
			{{$msg = " "}}
		{{end}}

			{{if eq (toFloat 1)  (toFloat (len .CmdArgs))}}
				{{$embed := cembed
				"description" (joinStr "" "**" $user "** ▬ " $msg "\n"
				"<:next:723131844643651655> [*Dé : " $d " (" $v ") | " $arg1 " : " $x " | " $imp " : " $idb " | Seuil : " $seuil "* ]")
				"color" 0x63AFE1}}
				{{sendMessage nil $embed}}
			{{else}}
				{{if $manuel}}
					{{$c =  joinStr "" (slice .CmdArgs 1)}}
					{{$c = joinStr " " (split $c $manuel)}}
					{{if ne $c " "}}
						{{$c = joinStr " " (split $c $manuel) ": "}}
					{{end}}
				{{else}}
					{{$c = joinStr " " (slice .CmdArgs 1) ": " }}
				{{end}}

				{{if eq (toFloat (index .CmdArgs 1)) (toFloat 0)}}
					{{$embed := cembed
					"description" (joinStr "" "**" $user "**  ▬ " $c $msg "\n"
					"<:next:723131844643651655> [*Dé : " $d " (" $v ") | " $arg1 " : " $x " | " $imp " : " $idb " | Seuil : " $seuil "* ]")
					"color" 0x63AFE1}}
					{{sendMessage nil $embed}}

				{{else if ne (toFloat (index .CmdArgs 1)) (toFloat 0)}}
					{{$m := (toInt (index .CmdArgs 1)) }}
					{{$y := $m}}
					{{if gt $m (toInt 0)}}
						{{$argm = 1}}
						{{if eq $m (toInt 1)}}
							{{$arg2 = "Pénalité "}}
						{{else}}
							{{$arg2 = "Pénalités "}}
						{{end}}
					{{else}}
						{{$argm = 0}}
						{{$m = mult $m (toInt -1)}}
						{{$y = mult $y (toInt -1)}}
						{{$arg2 = "Bonus "}}
					{{end}}


					{{if eq $arg 0}}
						{{if eq $argm 0}}
							{{$d = sub $d $m}}
						{{else}}
							{{$d = add $m $d}}
						{{end}}
					{{else}}
						{{if eq $argm 0}}
							{{$m = add $m $idb}}
							{{$mimp = $m}}
							{{$d = sub $d $m}}
						{{else}}
							{{$d = add $m $d}}
							{{$mimp = $idb}}
						{{end}}
					{{end}}


					{{if gt $d (toInt 10)}}
						{{$d = toInt 10}}
					{{else if lt $d (toInt 0)}}
						{{$d = (toInt 0)}}
					{{end}}

					{{if eq $d (toInt 0)}}
						{{$msg ="**Ultra critique !** "}}
					{{else if eq $d (toInt 1)}}
						{{$msg = "**Réussite critique !** "}}
					{{else if le $d $seuil}}
							{{if ge $mimp (toInt 1)}}
								{{if eq $d $seuil}}
									{{$msg = "**Echec** "}}
								{{else}}
									{{$msg = "**Réussite** "}}
								{{end}}
							{{else}}
								{{$msg = "**Réussite** "}}
							{{end}}
					{{else if or (gt $d $seuil) (lt $d (toInt 9))}}
						{{$msg = "**Echec** "}}
					{{else if eq $d (toInt 10)}}
						{{$msg = "**Echec critique !** "}}
					{{else if eq $d (toInt 9)}}
						{{$msg = "**Echec...** "}}
					{{else}}
						{{$msg = " "}}
					{{end}}

				{{if eq (toFloat 2) (toFloat (len .CmdArgs))}}
					{{$embed := cembed
						"description" (joinStr "" "**" $user "** ▬ " $msg "\n"
						"<:next:723131844643651655>[*Dé : " $d " (" $v ") | " $x " : " $arg1 " | " $y " " $arg2 " | " $imp " : " $idb " | Seuil : " $seuil "*]")
						"color" 0x63AFE1}}
					{{sendMessage nil $embed}}

					{{else}}
						{{if $manuel}}
							{{$c =  joinStr "" (slice .CmdArgs 2)}}
							{{$c = joinStr " " (split $c $manuel)}}
							{{if ne $c " "}}
								{{$c = joinStr " " (split $c $manuel) ": " }}
							{{end}}
						{{else}}
							{{$c = joinStr " " (slice .CmdArgs 2) ": "  }}
						{{end}}
						{{$embed := cembed
						"description" (joinStr "" "**" $user "**  ▬ " $c " " $msg "\n"
						"<:next:723131844643651655>[*Dé : " $d " (" $v ") | " $arg1 " : " $x " | " $arg2 " : " $y " | " $imp " : " $idb " | Seuil : " $seuil "*]")
						"color" 0x63AFE1}}
						{{sendMessage nil $embed}}


				{{end}}
			{{end}}
		{{end}}

	{{else}}
		{{$d = sub $d $idb}}

		{{if gt $d (toInt 10)}}
			{{$d = toInt 10}}
		{{else if lt $d (toInt 0)}}
			{{$d = (toInt 0)}}
		{{end}}

		{{if eq $d (toInt 0)}}
			{{$msg ="**Ultra critique !** "}}
		{{else if eq $d (toInt 1)}}
			{{$msg = "**Réussite critique !** "}}
		{{else if le $d $seuil}}
				{{if ge $idb (toInt 1)}}
					{{if eq $d $seuil}}
						{{$msg = "**Echec** "}}
					{{else}}
						{{$msg = "**Réussite** "}}
					{{end}}
				{{else}}
					{{$msg = "**Réussite** "}}
				{{end}}
		{{else if or (gt $d $seuil) (lt $d (toInt 9))}}
			{{$msg = "**Echec** "}}
		{{else if eq $d (toInt 10)}}
			{{$msg = "**Echec critique !** "}}
		{{else if eq $d (toInt 9)}}
			{{$msg = "**Echec...** "}}
		{{else}}
			{{$msg = " "}}
		{{end}}

		{{$embed := cembed
		"description" (joinStr "" "**" $user "**  ▬  " (joinStr " " $c) " : " $msg "\n"
	"<:next:723131844643651655>[*Dé : " $d  " (" $v ") "  " | " $imp " : " $idb " | Seuil : " $seuil "* ]")
		"color" 0x63AFE1}}
		{{sendMessage nil $embed}}
	{{end}}

{{else}}
{{if gt $d (toInt 10)}}
	{{$d = toInt 10}}
{{else if lt $d (toInt 0)}}
	{{$d = (toInt 0)}}
{{end}}

{{if eq $d (toInt 0)}}
	{{$msg ="**Ultra critique !** "}}
{{else if eq $d (toInt 1)}}
	{{$msg = "**Réussite critique !** "}}
{{else if le $d $seuil}}
		{{if ge $idb (toInt 1)}}
			{{if eq $d $seuil}}
				{{$msg = "**Echec** "}}
			{{else}}
				{{$msg = "**Réussite** "}}
			{{end}}
		{{else}}
			{{$msg = "**Réussite** "}}
		{{end}}
{{else if or (gt $d $seuil) (lt $d (toInt 9))}}
	{{$msg = "**Echec** "}}
{{else if eq $d (toInt 10)}}
	{{$msg = "**Echec critique !** "}}
{{else if eq $d (toInt 9)}}
	{{$msg = "**Echec...** "}}
{{else}}
	{{$msg = " "}}
{{end}}
	{{$embed := cembed
	"description" (joinStr "" "**" $user "** ▬ " $msg "\n"
	"<:next:723131844643651655>[*Dé : " $d "* ]")
	"color" 0x63AFE1}}
	{{sendMessage nil $embed}}
{{end}}

{{deleteTrigger 1}}
