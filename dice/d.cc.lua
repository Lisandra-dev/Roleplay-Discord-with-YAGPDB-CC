{{$col := 0xD2D1F5}}
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
			{{$col = 0x716DE9}}
		{{else if eq $v (toInt 1)}}
			{{$msg = "**Réussite critique !**"}}
			{{$col = 0x538DAD}}
		{{else if eq $v (toInt 10)}}
			{{$msg = "**Echec critique !**"}}
			{{$col = 0x951414}}
		{{else if eq $v (toInt 9)}}
			{{$msg = "**Echec...**"}}
			{{$col = 0x895A5A}}
		{{else}}
			{{$msg = (joinStr "" "Résultat : " (toString (toInt $d)))}}
			{{$col = 0xD2D1F5}}
		{{end}}
		{{$real:=(toString $d)}}
		{{$d := (joinStr "" "Valeur du dé : " (toString (toInt $d))) }}
		{{if eq (toFloat 1)  (toFloat (len .CmdArgs))}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" (.User.AvatarURL "256"))
				"description" (joinStr " " $msg "*(avec " $i " " $implant ")*" )
				"footer" (sdict "text" $d)
				"color" $col}}
			{{sendMessage nil $embed}}
		{{else}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" (.User.AvatarURL "256"))
				"description" (joinStr " " "**" (slice .CmdArgs 1) "** \n" $msg "*(avec " $i " " $implant ")*" )
				"footer" (sdict "text" (joinStr " " "Valeur du dé :" $v "|" $real))
				"color" $col}}
			{{sendMessage nil $embed}}
		{{end}}
	{{else}}
		{{$v := $d}}
		{{if eq $v (toInt 0)}}
			{{$msg = "**Ultra réussite critique !**"}}
			{{$col = 0x716DE9}}
		{{else if eq $v (toInt 1)}}
			{{$msg = "**Réussite critique !**"}}
			{{$col = 0x538DAD}}
		{{else if eq $v (toInt 10)}}
			{{$msg = "**Echec critique !**"}}
			{{$col = 0x951414}}
		{{else if eq $v (toInt 9)}}
			{{$msg = "**Echec...*"}}
			{{$col = 0x895A5A}}
		{{else}}
			{{$msg = (joinStr "" "Résultat : " (toString (toInt $d)))}}
			{{$col = 0xD2D1F5}}
		{{end}}
		{{$d := (joinStr "" (toString $d)) }}
		{{$embed := cembed
			"author" (sdict "name" $user "icon_url" (.User.AvatarURL "256"))
			"description" (joinStr " " "**" .CmdArgs "**" "\n" $msg)
			"footer" (sdict "text" (joinStr " " "Valeur du dé :" $d ))
			"color" $col}}
		{{sendMessage nil $embed}}
	{{end}}
{{else}}
	{{$v := $d}}
	{{if eq $v (toInt 0)}}
		{{$msg = "**Ultra réussite critique !**"}}
		{{$col = 0x716DE9}}
	{{else if eq $v (toInt 1)}}
		{{$msg = "**Réussite critique !**"}}
		{{$col = 0x538DAD}}
	{{else if eq $v (toInt 10)}}
		{{$msg = "**Echec critique !**"}}
		{{$col = 0x951414}}
	{{else if eq $v (toInt 9)}}
		{{$msg = "**Echec...*"}}
		{{$col = 0x895A5A}}
	{{else}}
		{{$msg = (joinStr "" "Résultat : " $d) }}
		{{$col = 0xD2D1F5}}
	{{end}}
	{{$embed := cembed
		"author" (sdict "name" $user "icon_url"  (.User.AvatarURL "256"))
		"description" (joinStr " " $msg)
		"footer" (sdict  "text" (joinStr " " "Valeur du dé :" (toString $d)))
		"color" $col}}
	{{sendMessage nil $embed}}
{{end}}
{{deleteTrigger 1}}
