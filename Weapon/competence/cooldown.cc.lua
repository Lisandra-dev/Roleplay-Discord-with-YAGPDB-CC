
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

{{if .CmdArgs}}
	{{if eq (index .CmdArgs 0) "attaque"}}
		{{$nom:= (index .CmdArgs 1)}}
		{{dbSet .User.ID "comp" $nom}}
		{{$arg:= (dbGet .User.ID "comp").Value}}
		{{if not (dbGet .User.ID $arg)}}
			{{dbSet .User.ID $arg 1}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imga)
				"description" (joinStr "" "Début du cooldown pour la compétence " $arg)
				"color" 0xDFAA58}}
			{{ $id := sendMessageRetID nil $embed }}
			{{deleteMessage nil $id 30}}
		{{else}}
			{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $imga)
			"description" (joinStr "" "La compétence " $arg " a déjà été utilisée")
			"color" 0xDFAA58}}
			{{ $id := sendMessageRetID nil $embed }}
			{{deleteMessage nil $id 30}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "support"}}
		{{$nom:= (index .CmdArgs 1)}}
		{{dbSet .User.ID "aide" $nom}}
		{{$arg:= (dbGet .User.ID "aide").Value}}
		{{if not (dbGet .User.ID $arg)}}
			{{dbSet .User.ID $arg 1}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imgs)
				"description" (joinStr "" "Début du cooldown pour la compétence " $arg)
				"color" 0xB57CA3}}
			{{ $id := sendMessageRetID nil $embed }}
			{{deleteMessage nil $id 30}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imgs)
				"description" (joinStr "" "La compétence " $arg " a déjà été utilisée")
			"color" 0xB57CA3}}
			{{ $id := sendMessageRetID nil $embed }}
			{{deleteMessage nil $id 30}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "manuel"}}
		{{if lt (len .CmdArgs) 2}}
			{{"ERREUR"}}
		{{else}}
			{{$nom := (index .CmdArgs 1)}}
			{{dbSet .User.ID "manuel" $nom}}
			{{$cd := (index .CmdArgs 2)}}
			{{dbSet .User.ID "cdmanuel" $cd}}
			{{$arg := (dbGet .User.ID "manuel").Value}}
			{{if not (dbGet .User.ID $arg)}}
				{{dbSet .User.ID $arg 1}}
				{{$embed := cembed
				"author" (sdict "name" $user "icon_url" $imgm)
				"description" (joinStr "" "Début du cooldown pour la compétence " $arg)
				"color" 0x977ED0}}
			{{$id := sendMessageRetID nil $embed}}
			{{deleteMessage nil $id 30}}
			{{else}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $imgm)
					"description" (joinStr "" "La compétence " $arg " a déjà été utilisée")
					"color" 0x977ED0}}
				{{ $id := sendMessageRetID nil $embed }}
				{{deleteMessage nil $id 30}}
			{{end}}
		{{end}}


	{{else}}
		{{ $embed := cembed
			"description" "Merci d'indiquer le type de cooldown que vous utilisez.\n\n **Usage** : `$cooldown <attaque|support> nom`"}}
		{{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
	{{end}}
{{else}}
		{{ $embed := cembed
		"description" "Merci d'indiquer le type de compétence ainsi que son nom. \n\n **Usage** : `$cooldown <attaque|support> nom`" }}
		{{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
{{end}}
{{deleteTrigger 1}}
