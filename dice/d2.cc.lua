{{$user:=.Member.Nick}}
{{if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$d:= (randInt 1 10)}}
{{$implant:= "implant"}}
{{$msg:= ""}}
{{ if .CmdArgs}}
	{{if ne (toFloat (index .CmdArgs 0)) (toFloat 0)}}
		{{$i:=(toInt (index .CmdArgs 0)) }}
		{{if ge (toFloat $i) (toFloat 2)}}
			{{$implant = "implants"}}
		{{end}}
		{{$v :=sub $d $i}}
		{{if eq $v (toInt 0)}}
			{{$msg ="**Ultra réussite critique !**"}}
		{{else if eq $v (toInt 1)}}
			{{$msg = "**Réussite critique !**"}}
		{{else if eq $v (toInt 10)}}
			{{$msg = "**Echec critique !**"}}
		{{else if eq $v (toInt 9)}}
			{{$msg = "**Echec...**"}}
		{{else}}
			{{$msg = ""}}
		{{end}}
		{{if eq (toFloat 1)  (toFloat (len .CmdArgs))}}
**{{$user}}** : {{$msg}}
	Vous obtenez : {{$v}} avec {{$i}} {{$implant}} *({{$d}})*
		{{else}}
**{{$user}}** ▬ {{joinStr " " (slice .CmdArgs 1)}} : {{$msg}}
	Vous obtenez : {{$v}} en utilisant {{$i}} {{$implant}} *({{$d}})*
		{{end}}
	{{else}}
		{{$v := $d}}
		{{if eq $v (toInt 0)}}
			{{$msg = "**Ultra réussite critique !**"}}
		{{else if eq $v (toInt 1)}}
			{{$msg = "**Réussite critique !**"}}
		{{else if eq $v (toInt 10)}}
			{{$msg = "**Echec critique !**"}}
		{{else if eq $v (toInt 9)}}
			{{$msg = "**Echec...*"}}
		{{else}}
			{{$msg = ""}}
		{{end}}
**{{$user}}** ▬ {{joinStr " " .CmdArgs }} : {{$msg}}
	Vous obtenez : {{$v}}
	{{end}}
{{else}}
	{{$v := $d}}
	{{if eq $v (toInt 0)}}
		{{$msg = "**Ultra réussite critique !**"}}
	{{else if eq $v (toInt 1)}}
		{{$msg = "**Réussite critique !**"}}
	{{else if eq $v (toInt 10)}}
		{{$msg = "**Echec critique !**"}}
	{{else if eq $v (toInt 9)}}
		{{$msg = "**Echec...**"}}
	{{else}}
		{{$msg = ""}}
	{{end}}
**{{$user}}** : {{$msg}}
	Vous obtenez : {{$v}}
{{end}}
{{deleteTrigger 1}}
