{{if .CmdArgs}}
	{{if eq (index .CmdArgs 0) "attaque"}}
		{{$nom:= (index .CmdArgs 1)}}
		{{dbSet .User.ID "comp" $nom}}
		{{$arg:= (dbGet .User.ID "comp").Value}}
		{{if not (dbGet .User.ID $arg)}}
			{{dbSet .User.ID $arg 1}}
			{{ $embed := cembed
				"description" (joinStr "" "Début du cooldown pour la compétence " $arg)}}
			{{ $id := sendMessageRetID nil $embed }}
			{{deleteMessage nil $id 30}}
		{{else}}
			{{ $embed := cembed
			"description" (joinStr "" "La compétence " $arg " a déjà été utilisée")}}
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
				"description" (joinStr "" "Début du cooldown pour la compétence " $arg)}}
			{{ $id := sendMessageRetID nil $embed }}
			{{deleteMessage nil $id 30}}
		{{else}}
			{{ $embed := cembed
				"description" (joinStr "" "La compétence " $arg " a déjà été utilisée")}}
			{{ $id := sendMessageRetID nil $embed }}
			{{deleteMessage nil $id 30}}
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
bite
