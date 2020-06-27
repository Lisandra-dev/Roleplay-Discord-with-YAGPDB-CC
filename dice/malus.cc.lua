


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
		{{$res = (joinStr "" "Dé : " (toString $d) " " $v " | En ayant " (toString $i) " " $arg1 "." )}}
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

			{{$res = (joinStr "" "Dé : " (toString $d) " " $v  " | En ayant " (toString $i) " " $arg1 " et " (toString $m) " " $arg2 "." )}}
		{{else}}
		{{$res = (joinStr "" "Dé : " (toString $d) " " $v " | En ayant " (toString $i) " " $arg1 "." )}}
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
**{{$user}}** ▬ {{$comm}}**Ultra critique : ** Votre cible a une pénalité de +4 à son dé.
		*[{{$res}}]*
	{{else if eq $d (toInt 1) }}
**{{$user}}** ▬ {{$comm}}**Réussite critique :** Votre cible a une pénalité de +3 à son dé.
		*[{{$res}}]*
    {{else if eq $d (toInt 2)}}
**{{$user}}** ▬ {{$comm}}**Réussite critique :** Votre cible a une pénalité de +3 à son dé.
	*[{{$res}}]*
	{{else if and (le $d (toInt 5)) (ge $d (toInt 3))}}
**{{$user}}** ▬ {{$comm}}*Si votre seuil le permet* : Votre cible a une pénalité de +2 à son dé.
	*[{{$res}}]*
	{{else if  and (le $d (toInt 8)) (ge $d (toInt 6))}}
**{{$user}}** ▬ {{$comm}}*Si votre seuil le permet* : Votre cible a une pénalité de +1 à son dé.
	*[{{$res}}]*
	{{else if eq $d (toInt 9)}}
**{{$user}}** ▬ {{$comm}}**Echec de l'altération**...
	*[{{$res}}]*
	{{else if eq $d (toInt 10)}}
**{{$user}}** ▬ {{$comm}}**Echec critique de l'altération : ** Votre cible a un bonus de -2 à son dé.
		*[{{$res}}]*
	{{end}}

{{deleteTrigger 1}}
