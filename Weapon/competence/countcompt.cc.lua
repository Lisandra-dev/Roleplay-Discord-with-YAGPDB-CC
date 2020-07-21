{{ $matches := (index (reFindAllSubmatches `^\((.*)\)|(^\$(.*))|(^\!\!(.+))|((<a?:[\w~]{2,32}:\d{17,19}>)|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})` .Message.Content)) }}
{{$arg := (dbGet .User.ID "comp").Value}}
{{$arg2 := (dbGet .User.ID "aide").Value}}
{{$arg3 := (dbGet .User.ID "manuel").Value}}

{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{if $name}}
	{{$user = $name}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}

{{$imga:="https://i.imgur.com/zNofnyh.png"}}
{{$imgs :="https://i.imgur.com/9iRdtbM.png" }}
{{$imgm := "https://i.imgur.com/FCy00x2.png"}}


{{if not $matches}}
	{{if (dbGet .User.ID $arg)}}
		{{if and (lt (dbGet .User.ID $arg).Value (toFloat 2)) (gt (len .Message.Content) 15) }}
				{{$incr := dbIncr .User.ID $arg 1}}
		{{else if (eq (dbGet .User.ID $arg).Value (toFloat 2))  }}
				{{dbDel .User.ID "comp"}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $imga)
					"description" (joinStr "" "Votre compétence " $arg " est de nouveau utilisable")
					"color" 0xDFAA58}}
				{{ $id := sendMessageRetID nil $embed }}
				{{deleteMessage nil $id 30}}
				{{dbDel .User.ID $arg}}
		{{end}}

	{{else if (dbGet .User.ID $arg2)}}
		{{if and (lt (dbGet .User.ID $arg2).Value (toFloat 5)) (gt (len .Message.Content) 15) }}
			{{$incr := dbIncr .User.ID $arg2 1}}
		{{else if (eq (dbGet .User.ID $arg2).Value (toFloat 5))}}
			{{dbDel .User.ID "aide"}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imgs)
				"description" (joinStr "" "Votre compétence " $arg2 " est de nouveau utilisable")
				"color" 0xB57CA3}}
			{{ $id := sendMessageRetID nil $embed }}
			{{deleteMessage nil $id 30}}
			{{dbDel .User.ID $arg2}}
		{{end}}

	{{else if (dbGet .User.ID $arg3)}}
		{{$cd := (dbGet .User.ID "cdmanuel").Value}}
		{{if and (lt (dbGet .User.ID $arg3).Value (toFloat $cd)) (gt (len .Message.Content) 15)}}
			{{$incr := dbIncr .User.ID $arg3 1}}
		{{else if eq (dbGet .User.ID $arg3).Value (toFloat $cd)}}
			{{dbDel .User.ID "cdmanuel"}}
			{{dbDel .User.ID "manuel"}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imgm)
				"description" (joinStr "" "Votre compétence " $arg3 " est de nouveau utilisable")
				"color" 0xB57CA3}}
			{{ $id := sendMessageRetID nil $embed }}
			{{deleteMessage nil $id 30}}
			{{dbDel .User.ID $arg3}}
	{{end}}
	{{else}}
	{{end}}
{{end}}
