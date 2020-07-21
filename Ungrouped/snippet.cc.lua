{{$chan := 701373579593252944}}

{{if .CmdArgs}}
		{{if or (eq (index .CmdArgs 0) "armes") (eq (index .CmdArgs 0) "arme")}}
	 		{{$message := getMessage $chan 733363513988218952 }}
			{{$embed := cembed
    		"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
				"color" (randInt 111111 999999)
				"description" (joinStr "" $message.Content "\n\n [➥ Original](https://discordapp.com/channels/701368588690522112/701373579593252944/710817488392290305) \n")
				"footer" (sdict "text" (joinStr "" "Cité par : " .User.String)
				"icon_url" (.User.AvatarURL "512"))}}
				{{sendMessage nil $embed}}

		{{else if eq (index .CmdArgs 0) "résumé"}}
			{{$embed := cembed
				"title" "Commande de base"
				"color" (randInt 111111 999999)
				"description" ":white_small_square: Pour un dé simple : `$d Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n :white_small_square: Pour un dé avec bonus extérieur à un implant : `$d -bonus Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n :white_small_square: Pour un dé avec malus : `$d malus Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n :white_small_square: Pour un dé avec malus, bonus : `$d -bonus malus Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n:\nwhite_small_square **Autre type d'action :** Pour les utiliser, remplacer `$d` par la commande suivante : \n <:tr:724626754282717194> :small_blue_diamond: `$poison` : poison\n <:tr:724626754282717194> :small_blue_diamond: `$soin` : Soin\n <:tr:724626754282717194> :small_blue_diamond: `$malus` : Malus (de statistiques)\n<:tr:724626754282717194> :small_blue_diamond: `$esquive` : Dé d'esquive"}}
			{{sendMessage nil $embed}}

	{{else if eq (index .CmdArgs 0) "dés"}}
		{{$message := getMessage 701373579593252944 727994181926256640}}
		{{$embed := cembed
    	"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" (randInt 111111 999999)
			"description" (joinStr "" $message.Content "\n\n [➥ Original](https://discordapp.com/channels/701368588690522112/701373579593252944/727992973131776020) \n")
			"footer" (sdict "text" (joinStr "" "Cité par : " .User.String)
			"icon_url" (.User.AvatarURL "512"))}}
		{{sendMessage nil $embed}}

	{{else if or (eq (index .CmdArgs 0) "dégats")  (eq (index .CmdArgs 0) "dégâts")  (eq (index .CmdArgs 0) "dégat") (eq (index .CmdArgs 0) "dégât")}}
		{{$message := getMessage $chan 701374206268538960}}
		{{$embed := cembed
    	"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" (randInt 111111 999999)
			"description" (joinStr "" $message.Content "\n\n [➥ Original](https://discordapp.com/channels/701368588690522112/701373579593252944/701374206268538960) \n")
			"footer" (sdict "text" (joinStr "" "Cité par : " .User.String)
			"icon_url" (.User.AvatarURL "512"))}}
		{{sendMessage nil $embed}}

	{{else if (eq (index .CmdArgs 0) "position")}}
		{{$embed :=cembed
		"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
		"color" (randInt 111111 999999)
		"description" ":white_small_square:  *Abrité* : +3 en précision pour les attaquants, mais ne peut ni attaquer, ni esquiver.\n:white_small_square: *A découvert* : -1 aux jets d'attaques (pour les attaquants).\n:white_small_square:  *Duel* : -1 aux jets d'attaques (pour les attaquants), et +2 en précision si visée sur les autres rangs."
		"footer" (sdict "text" (joinStr "" "Cité par : " .User.String)
		"icon_url" (.User.AvatarURL "512"))}}
		{{sendMessage nil $embed}}

	{{else if (eq (index .CmdArgs 0) "charge")}}
		{{$message := getMessage $chan 720919240181415976}}
		{{$embed := cembed
    	"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" (randInt 111111 999999)
			"description" (joinStr "" $message.Content)
			"icon_url" (.User.AvatarURL "512")}}
		{{sendMessage nil $embed}}

	{{else if eq (index .CmdArgs 0) "ticket"}}
	{{$embed := cembed
		"title" "Ticket"
		"color" 0x51CDEF
		"description" "Les tickets permettent de créer des channels privés, où seul la personne ayant créer le ticket, et celles invitées peuvent lire et discuter. Cependant, l'admin du serveur y a accès. L'avantage est évidemment la prise en compte par le bot de vos messages et donc, de faire avancer le timer. En tant qu'administrateur, je ne compte pas lire vos messages, et la catégorie est donc mute et invisible. Au besoin, vous pouvez me ping.\n\n:white_small_square: *Ouvrir un ticket* : `$ticket create <nom du channel>`\n:white_small_square: *Inviter des utilisateur* : `$ticket addUser @pseudo`\n:white_small_square: *Fermer un ticket* : `$ticket close <raison>`\n> Il est à noter que même si l'utilisateur n'apparaît pas dans la liste, si vous tapez son pseudo complet (type `@Mara`) discord le reconnaîtra sans soucis, mais au besoin vous pouvez utiliser l'identifiant)."}}
	{{sendMessage nil $embed}}

	{{else if eq (index .CmdArgs 0) "notes"}}
		{{$embed := cembed
			"title" "Bloc-notes"
			"color" 0x51CDEF
			"fields" (cslice
			(sdict "name" "Bloc-note du serveur" "value" ":white_small_square: *Créer une note* : `$note <titre> <contenu>`\n:white_small_square: *Accéder à une note* : `$get <titre>`\n:white_small_square: *Supprimé une note* : `$del <titre>`\n:white_small_square: *Accéder à la liste des notes du serveur* : `$list`\n" "inline" false)
			(sdict "name" "_ _" "value" "_ _" "inline" false)
			(sdict "name" "Bloc-note personnel" "value" "*Attention, les résultats sont envoyés en message privé.*\n:white_small_square: *Créer une note* : `$noteperso <titre> <contenu>`\n:white_small_square: *Accéder à une note* : `$getperso <titre>`\n:white_small_square: *Supprimé une note* : `$delperso <titre>`\n:white_small_square: *Accéder à la liste des notes du serveur* : `$listperso`" "inline" false))}}
		{{sendMessage nil $embed}}

	{{else if eq (index .CmdArgs 0) "horloge"}}
		{{$embed := cembed
			"title" "Horloge"
			"color" 0x51CDEF
			"description" ":white_small_square: *Avoir l'heure* : `$time`\n:white_small_square: *Historique du temps* : <#716988208205791342>\n:white_small_square: *Paramètres actuels* : `$settings`"}}
		{{sendMessage nil $embed}}

	{{else if eq (index .CmdArgs 0) "bouclier"}}
		{{$message := getMessage $chan 726465578809688135}}
		{{$embed := cembed
			"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" (randInt 111111 999999)
			"description" (joinStr "" $message.Content)
			"icon_url" (.User.AvatarURL "512")}}
		{{sendMessage nil $embed}}

	{{else if or (eq (index .CmdArgs 0) "snippet") (eq (index .CmdArgs 0) "all") (eq (index .CmdArgs 0) "help") }}
		{{$embed := cembed
			"title" "Liste des aides disponibles"
		"description" ":white_small_square: `!(armes|arme)`\n:white_small_square: `!résumé`\n:white_small_square: `!dés`\n:white_small_square: `!position`\n:white_small_square: `!(dégâts|dégât|dégat|dégats)`\n:white_small_square: `!notes`\n:white_small_square: `!horloge`\n:white_small_square: `!bouclier`\n:white_small_square:`!ticket`\n:white_small_square: `!charge`\n\n:white_medium_square: **Pour afficher cette liste** : `!(all|snippet|help)`"}}
		{{sendMessage nil $embed}}
	{{end}}
{{end}}
{{deleteTrigger 1}}
