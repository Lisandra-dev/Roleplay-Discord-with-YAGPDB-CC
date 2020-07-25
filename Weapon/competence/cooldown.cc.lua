
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
{{$user = title $user}}

{{$imga:="https://i.imgur.com/zNofnyh.png"}}
{{$imgs :="https://i.imgur.com/9iRdtbM.png" }}
{{$imgm := "https://i.imgur.com/FCy00x2.png"}}

{{if .CmdArgs}}
	{{if eq (index .CmdArgs 0) "attaque"}}
		{{$nom:= (index .CmdArgs 1)}}
		{{dbSet $id "comp" $nom}}
		{{$arg:= (dbGet $id "comp").Value}}
		{{if not (dbGet $id $arg)}}
			{{dbSet $id $arg 1}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imga)
				"description" (joinStr "" "Début du cooldown pour la compétence " $arg)
				"color" 0xDFAA58}}
			{{ $idM := sendMessageRetID 735938256038002818 $embed }}

		{{else}}
			{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $imga)
			"description" (joinStr "" "La compétence " $arg " a déjà été utilisée")
			"color" 0xDFAA58}}
			{{ $idM := sendMessageRetID 735938256038002818 $embed }}
		{{end}}
	{{else if eq (index .CmdArgs 0) "support"}}
		{{$nom:= (index .CmdArgs 1)}}
		{{dbSet $id "aide" $nom}}
		{{$arg:= (dbGet $id "aide").Value}}
		{{if not (dbGet $id $arg)}}
			{{dbSet $id $arg 1}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imgs)
				"description" (joinStr "" "Début du cooldown pour la compétence " $arg)
				"color" 0xB57CA3}}
			{{ $idM := sendMessageRetID 735938256038002818 $embed }}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imgs)
				"description" (joinStr "" "La compétence " $arg " a déjà été utilisée")
			"color" 0xB57CA3}}
			{{ $idM := sendMessageRetID 735938256038002818 $embed }}
		{{end}}
	{{else if eq (index .CmdArgs 0) "manuel"}}
		{{if lt (len .CmdArgs) 2}}
			{{"ERREUR"}}
		{{else}}
			{{$nom := (index .CmdArgs 1)}}
			{{dbSet $id "manuel" $nom}}
			{{$cd := (index .CmdArgs 2)}}
			{{dbSet $id "cdmanuel" $cd}}
			{{$arg := (dbGet $id "manuel").Value}}
			{{if not (dbGet $id $arg)}}
				{{dbSet $id $arg 1}}
				{{$embed := cembed
				"author" (sdict "name" $user "icon_url" $imgm)
				"description" (joinStr "" "Début du cooldown pour la compétence " $arg)
				"color" 0x977ED0}}
			{{$idM := sendMessageRetID 735938256038002818 $embed}}

			{{else}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $imgm)
					"description" (joinStr "" "La compétence " $arg " a déjà été utilisée")
					"color" 0x977ED0}}
				{{ $idM := sendMessageRetID 735938256038002818 $embed }}


			{{end}}
		{{end}}


	{{else}}
		{{ $embed := cembed
			"description" "Merci d'indiquer le type de cooldown que vous utilisez.\n\n **Usage** : `$cooldown <attaque|support> nom`"}}
		{{ $idM := sendMessageRetID 735938256038002818 $embed }}

	{{end}}
{{else}}
		{{ $embed := cembed
		"description" "Merci d'indiquer le type de compétence ainsi que son nom. \n\n **Usage** : `$cooldown <attaque|support> nom`" }}
		{{ $idM := sendMessageRetID 735938256038002818 $embed }}

{{end}}
{{deleteTrigger 1}}
