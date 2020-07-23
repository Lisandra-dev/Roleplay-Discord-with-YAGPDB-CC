{{ $matches := (index (reFindAllSubmatches `^\((.*)\)|(^\$(.*))|(^\!\!(.+))|((<a?:[\w~]{2,32}:\d{17,19}>)|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})` .Message.Content)) }}
{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id:= .User.ID}}
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

{{$arg := (dbGet $id "comp").Value}}
{{$arg2 := (dbGet $id "aide").Value}}
{{$arg3 := (dbGet $id "manuel").Value}}

{{$imga:="https://i.imgur.com/zNofnyh.png"}}
{{$imgs :="https://i.imgur.com/9iRdtbM.png" }}
{{$imgm := "https://i.imgur.com/FCy00x2.png"}}


{{if not $matches}}
	{{if (dbGet $id $arg)}}
		{{if and (lt (dbGet $id $arg).Value (toFloat 2)) (gt (len .Message.Content) 15) }}
				{{$incr := dbIncr $id $arg 1}}
		{{else if (eq (dbGet $id $arg).Value (toFloat 2))  }}
				{{dbDel $id "comp"}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $imga)
					"description" (joinStr "" "Votre compétence " $arg " est de nouveau utilisable")
					"color" 0xDFAA58}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
				{{dbDel $id $arg}}
		{{end}}

	{{else if (dbGet $id $arg2)}}
		{{if and (lt (dbGet $id $arg2).Value (toFloat 5)) (gt (len .Message.Content) 15) }}
			{{$incr := dbIncr $id $arg2 1}}
		{{else if (eq (dbGet $id $arg2).Value (toFloat 5))}}
			{{dbDel $id "aide"}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imgs)
				"description" (joinStr "" "Votre compétence " $arg2 " est de nouveau utilisable")
				"color" 0xB57CA3}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
			{{dbDel $id $arg2}}
		{{end}}

	{{else if (dbGet $id $arg3)}}
		{{$cd := (dbGet $id "cdmanuel").Value}}
		{{if and (lt (dbGet $id $arg3).Value (toFloat $cd)) (gt (len .Message.Content) 15)}}
			{{$incr := dbIncr $id $arg3 1}}
		{{else if eq (dbGet $id $arg3).Value (toFloat $cd)}}
			{{dbDel $id "cdmanuel"}}
			{{dbDel $id "manuel"}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imgm)
				"description" (joinStr "" "Votre compétence " $arg3 " est de nouveau utilisable")
				"color" 0xB57CA3}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
			{{dbDel $id $arg3}}
	{{end}}
	{{else}}
	{{end}}
{{end}}
