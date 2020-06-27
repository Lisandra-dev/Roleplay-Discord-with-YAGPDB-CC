{{if .CmdArgs}}
	{{$nom:= (index .CmdArgs 0)}}
	{{dbSet .User.ID "comp" $nom}}
	{{$arg:= (dbGet .User.ID "comp").Value}}
	{{if not (dbGet .User.ID $arg)}}
		{{dbSet .User.ID $arg 1}}
		{{dbSet .User.ID "run" (toString .Channel.ID)}}
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
	"description" "Merci d'indiquer la compétence utilisée"}}
	{{ $id := sendMessageRetID nil $embed }}
	{{deleteMessage nil $id 30}}
{{end}}
{{deleteTrigger 1}}
