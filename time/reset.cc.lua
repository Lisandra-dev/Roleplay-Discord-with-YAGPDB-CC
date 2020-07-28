{{$db := dbGet .User.ID "resettime"}}
{{if not $db}}
	{{$cslice := shuffle (cslice "a" "b" "c" "d" "e" "f" "g" "h" "i" "j" "k" "l" "m" "n" "o" "p" "q" "r" "s" "t" "u" "v" "w" "x" "y" "z")}}
	{{$random := ""}}{{range (seq 0 10)}}{{$r := index (shuffle (cslice 0 1)) 0}}{{if eq $r 1}}{{$random = print $random (upper (index $cslice .))}}{{else}}{{$random = print $random (index $cslice .)}}{{end}}{{end}}
	{{dbSetExpire .User.ID "resettime" (cslice $random) 60}}
	Si vous êtes sûr de vouloir reset, merci de taper : `$reset {{$random}} `
{{else}}
	{{$yes := (dbGet .User.ID "resettime").Value}}
	{{$acc := ""}}
	{{range $yes}}
		{{$acc = (print .)}}
	{{end}}
	{{if (reFind $acc .Message.Content)}}
		{{$count := "count"}}
		{{$time := "time"}}
		{{$day := "jour"}}
		{{$message := "message"}}
		{{$embed := cembed
		    "title" "Reset ! "
		    "color" 0x670404
		    "timestamp" .Message.Timestamp}}
		{{ sendMessage nil $embed}}
		{{dbSet 0 $count 0}}
		{{dbSet 0 $time 1}}
		{{dbSet 0 $message 0}}
		{{dbSet 0 $day 1}}
	{{end}}
{{end}}
