{{$user:=.Member.Nick}}
{{if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$d:= (randInt 1 10)}}

{{$arg1:=""}}
{{$arg2:= ""}}
{{$v:=$d}}

{{$msg := ""}}

{{ if .CmdArgs}}
	{{if ne (toFloat (index .CmdArgs 0)) (toFloat 0)}}
		{{$i:=(toInt (index .CmdArgs 0)) }}
		{{$d =add $d $i}}
		{{if gt $d (toInt 10)}}
			{{$d = toInt 10}}
		{{else if lt $d (toInt 0)}}
			{{$d = (toInt 0)}}
		{{end}}
		{{if eq $d (toInt 0)}}
			{{$msg =" ▬ **Ultra critique !** "}}
		{{else if eq $d (toInt 1)}}
			{{$msg = " ▬ **Réussite critique !** "}}
		{{else if eq $d (toInt 10)}}
			{{$msg = " ▬ **Echec critique !** "}}
		{{else if eq $d (toInt 9)}}
			{{$msg = " ▬ **Echec...** "}}
		{{else}}
			{{$msg = " "}}
		{{end}}

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

			{{if eq (toFloat 1)  (toFloat (len .CmdArgs))}}
**{{$user}}**{{$msg}}
	[*Dé : {{$d}} ({{$v}}) | {{$i}} {{$arg1}}*]
			{{else}}
				{{if eq (toFloat (index .CmdArgs 1)) (toFloat 0)}}
**{{$user}}** : {{joinStr " " (slice .CmdArgs 1)}}{{$msg}}
	[*Dé : {{$d}} ({{$v}}) | {{$i}} {{$arg1}}*]
				{{else if ne (toFloat (index .CmdArgs 1)) (toFloat 0)}}
					{{$m:=(toInt (index .CmdArgs 1)) }}
					{{$d = add $d $m}}
					{{if eq $d (toInt 0)}}
						{{$msg =" ▬ **Ultra critique !** "}}
					{{else if eq $d (toInt 1)}}
						{{$msg = " ▬ **Réussite critique !** "}}
					{{else if eq $d (toInt 10)}}
						{{$msg = " ▬ **Echec critique !** "}}
					{{else if eq $d (toInt 9)}}
						{{$msg = " ▬ **Echec...** "}}
					{{else}}
						{{$msg = " "}}
					{{end}}
					{{if gt $d (toInt 10)}}
						{{$d = toInt 10}}
					{{else if lt $d (toInt 0)}}
						{{$d = (toInt 0)}}
					{{end}}

					{{if gt $m (toInt 0)}}
						{{if eq $m (toInt 1)}}
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

				{{if eq (toFloat 2) (toFloat (len .CmdArgs))}}
**{{$user}}**{{$msg}}
	[*Dé : {{$d}} ({{$v}}) | {{$i}} {{$arg1}} et {{$m}} {{$arg2}}*]
					{{else}}
**{{$user}}** : {{joinStr " " (slice .CmdArgs 2)}}{{$msg}}
	[*Dé : {{$d}} ({{$v}}) | {{$i}} {{$arg1}} et {{$m}} {{$arg2}}*]
				{{end}}
			{{end}}
		{{end}}

	{{else}}
		{{$v := $d}}
		{{if eq $v (toInt 0)}}
			{{$msg = " ▬ **Ultra critique !** "}}
		{{else if eq $v (toInt 1)}}
			{{$msg = " ▬ **Réussite critique !** "}}
		{{else if eq $v (toInt 10)}}
			{{$msg = " ▬ **Echec critique !** "}}
		{{else if eq $v (toInt 9)}}
			{{$msg = " ▬ **Echec...** "}}
		{{else}}
			{{$msg = " "}}
		{{end}}
**{{$user}}** : {{joinStr " " .CmdArgs}}{{$msg}}
[*Dé : {{$v}}*]
	{{end}}

{{else}}
	{{$v := $d}}
	{{if eq $v (toInt 0)}}
		{{$msg = " ▬ **Ultra critique !**"}}
	{{else if eq $v (toInt 1)}}
		{{$msg = " ▬ **Réussite critique !**"}}
	{{else if eq $v (toInt 10)}}
		{{$msg = " ▬ **Echec critique !**"}}
	{{else if eq $v (toInt 9)}}
		{{$msg = " ▬ **Echec...**"}}
	{{else}}
		{{$msg = " : "}}
	{{end}}
**{{$user}}**{{$msg}}
	[*Dé : {{$v}}*]
{{end}}

{{deleteTrigger 1}}
