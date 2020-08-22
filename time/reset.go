{{/*All message database counter*/}}
{{$time := sdict }}
{{with (dbGet 0 "time")}}
	{{$time = sdict .Value}}
{{end}}
{{$desc := ""}}
{{$db := dbGet .User.ID "resettime"}}
{{if not $db}}
	{{$cslice := shuffle (cslice "a" "b" "c" "d" "e" "f" "g" "h" "i" "j" "k" "l" "m" "n" "o" "p" "q" "r" "s" "t" "u" "v" "w" "x" "y" "z")}}
	{{$random := ""}}
	{{range (seq 0 10)}}
		{{$r := index (shuffle (cslice 0 1)) 0}}
		{{if eq $r 1}}
			{{$random = print $random (upper (index $cslice .))}}
		{{else}}
			{{$random = print $random (index $cslice .)}}
		{{end}}
	{{end}}
	{{dbSetExpire .User.ID "resettime" (cslice $random) 60}}
	{{$desc = joinStr " " "Si vous êtes sûr de vouloir reset, merci de taper : `$reset" $random "`"}}

{{else}}
	{{$yes := (dbGet .User.ID "resettime").Value}}
	{{$acc := ""}}
	{{range $yes}}
		{{$acc = (print .)}}
	{{end}}
	{{if (reFind $acc .Message.Content)}}
		{{$desc = "Reset !"}}
		{{$time.Set "count" 0}}
		{{$time.Set "cycle" 1}}
		{{$time.Set "message" 0}}
		{{$time.Set "jour" 1}}
	{{end}}
{{end}}

{{$embed := cembed
	"title" "Reset !"
	"color" 0x670404
	"description" $desc
	"timestamp" .CurrentTime}}

{{dbSet 0 "time" $time}}
