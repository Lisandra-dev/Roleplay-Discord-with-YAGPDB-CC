{{$chancommande := 701373579593252944}}
{{$combat := 701373579593252944}}
{{$col := 16777215}}
{{$p := 0}}
{{$r := .Member.Roles}}
{{range .Guild.Roles}}
	{{if and (in $r .ID) (.Color) (lt $p .Position)}}
	{{$p = .Position}}
	{{$col = .Color}}
	{{end}}
{{end}}
{{$link := ""}}

{{$commande := reFind `^\?(shop|boucliers?|(charg(é|e)(s|r)?)|(recharg(é|e)(s|r)?)|horloges?|times?|help|snippet|aide|all|tickets?|notes?|d(é|e)s?|dices?|r(e|é)sum(e|é)s?|armes?|d(é|e)g(a|â)s?|store|PA|pa)` .Message.Content}}

{{if .CmdArgs}}
		{{if or (eq $commande "?arme") (eq $commande "?armes")}}
	 		{{$message := getMessage $combat 733363513988218952 }}
			{{$link = print "(https://discordapp.com/channels/" .Guild.ID "/" $combat "/733363513988218952)"}}
			{{$embed := cembed
    		"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
				"color" $col
				"description" (joinStr "" $message.Content "\n\n [➥ Original]" $link " \n")
				"footer" (sdict "text" (joinStr "" "Cité par : " .User.String)
				"icon_url" (.User.AvatarURL "512"))}}
				{{sendMessage nil $embed}}
				{{deleteTrigger 1}}

		{{else if or (eq $commande "?résumé") (eq $commande "?resume") (eq $commande "?résume") (eq $commande "?resumé") (eq $commande "?résumés") (eq $commande "?resumes") (eq $commande "?résumes") (eq $commande "?resumés")}}
			{{$embed := cembed
				"title" "Commande de base"
				"color" $col
				"description" ":white_small_square: Pour un dé simple : `$d Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n :white_small_square: Pour un dé avec bonus extérieur à un implant : `$d -bonus Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n :white_small_square: Pour un dé avec malus : `$d malus Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n :white_small_square: Pour un dé avec malus, bonus : `$d -bonus malus Statistique - Description de l'action (cible, rang, type d'arme, module...)`\n:\nwhite_small_square **Autre type d'action :** Pour les utiliser, remplacer `$d` par la commande suivante : \n <:tr:724626754282717194> :small_blue_diamond: `$poison` : poison\n <:tr:724626754282717194> :small_blue_diamond: `$soin` : Soin\n <:tr:724626754282717194> :small_blue_diamond: `$malus` : Malus (de statistiques)"}}
			{{sendMessage nil $embed}}
			{{deleteTrigger 1}}

	{{else if or (eq $commande "?dés") (eq $commande "?dice") (eq $commande "?dé") (eq $commande "?de") (eq $commande "?dices")}}
		{{$link = (print "(https://discordapp.com/channels/" .Guild.ID "/" $combat "/727992973131776020)")}}
		{{$message := getMessage $combat 727992973131776020 }}
		{{$embed := cembed
    	"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"description" (joinStr "" $message.Content "\n\n [➥ Original]" $link " \n")
			"footer" (sdict "text" (joinStr "" "Cité par : " .User.String)
			"icon_url" (.User.AvatarURL "512"))}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "dégât") (eq $commande "?dégat") (eq $commande "?degât") (eq $commande "?degat") (eq $commande "?dégâts") (eq $commande "?dégats") (eq $commande "?degâts") (eq $commande "?degats")}}
		{{$link = (print "(https://discordapp.com/channels/" .Guild.ID "/" $combat "/7701374206268538960)")}}
		{{$message := getMessage $combat 701374206268538960}}
		{{$embed := cembed
    	"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"description" (joinStr "" $message.Content "\n\n [➥ Original]" $link "\n")
			"footer" (sdict "text" (joinStr "" "Cité par : " .User.String)
			"icon_url" (.User.AvatarURL "512"))}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "?position") (eq $commande "?positions")}}
		{{$embed :=cembed
		"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
		"color" $col
		"description" ":white_small_square:  *Abrité* : +3 en précision pour les attaquants, mais ne peut ni attaquer, ni esquiver.\n:white_small_square: *A découvert* : -1 aux jets d'attaques (pour les attaquants).\n:white_small_square:  *Duel* : -1 aux jets d'attaques (pour les attaquants), et +2 en précision si visée sur les autres rangs."
		"footer" (sdict "text" (joinStr "" "Cité par : " .User.String)
		"icon_url" (.User.AvatarURL "512"))}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "?recharge") (eq $commande "?recharger") (eq $commande "?charger") (eq $commande "?rechargé") (eq $commande "?chargé")  (eq $commande "?charge") (eq $commande "?charger")}}
	{{$link = (print "(https://discordapp.com/channels/" .Guild.ID "/" 734838748721840188 "/735869281698447463)")}}
		{{$message := getMessage 734838748721840188 735869281698447463}}
		{{$embed := cembed
    	"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"description" (joinStr "" $message.Content "\n\n [➥ Original]" $link "\n")
			"icon_url" (.User.AvatarURL "512")}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "?ticket") (eq $commande "?tickets")}}
	{{$embed := cembed
		"title" "Ticket"
		"color" 0x51CDEF
		"description" "Les tickets permettent de créer des channels privés, où seul la personne ayant créer le ticket, et celles invitées peuvent lire et discuter. Cependant, l'admin du serveur y a accès. L'avantage est évidemment la prise en compte par le bot de vos messages et donc, de faire avancer le timer. En tant qu'administrateur, je ne compte pas lire vos messages, et la catégorie est donc mute et invisible. Au besoin, vous pouvez me ping.\n\n:white_small_square: *Ouvrir un ticket* : `$ticket create <nom du channel>`\n:white_small_square: *Inviter des utilisateur* : `$ticket addUser @pseudo`\n:white_small_square: *Fermer un ticket* : `$ticket close <raison>`\n> Il est à noter que même si l'utilisateur n'apparaît pas dans la liste, si vous tapez son pseudo complet (type `@Mara`) discord le reconnaîtra sans soucis, mais au besoin vous pouvez utiliser l'identifiant)."}}
	{{sendMessage nil $embed}}
	{{deleteTrigger 1}}

	{{else if or (eq $commande "?note") (eq $commande "?notes")}}
		{{$embed := cembed
			"title" "Bloc-notes"
			"color" 0x51CDEF
			"fields" (cslice
			(sdict "name" "Bloc-note du serveur" "value" ":white_small_square: *Créer une note* : `$note <titre> <contenu>`\n:white_small_square: *Accéder à une note* : `$get <titre>`\n:white_small_square: *Supprimé une note* : `$del <titre>`\n:white_small_square: *Accéder à la liste des notes du serveur* : `$list`\n" "inline" false)
			(sdict "name" "_ _" "value" "_ _" "inline" false)
			(sdict "name" "Bloc-note personnel" "value" "*Attention, les résultats sont envoyés en message privé.*\n:white_small_square: *Créer une note* : `$noteperso <titre> <contenu>`\n:white_small_square: *Accéder à une note* : `$getperso <titre>`\n:white_small_square: *Supprimé une note* : `$delperso <titre>`\n:white_small_square: *Accéder à la liste des notes du serveur* : `$listperso`" "inline" false))}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "?horloge") (eq $commande "?horloges") (eq $commande "?time") (eq $commande "?times")}}
		{{$embed := cembed
			"title" "Horloge"
			"color" 0x51CDEF
			"description" ":white_small_square: *Avoir l'heure* : `$time`\n:white_small_square: *Historique du temps* : <#716988208205791342>\n:white_small_square: *Paramètres actuels* : `$settings`"}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "?bouclier") (eq $commande "?boucliers")}}
	{{$link = (print "(https://discordapp.com/channels/" .Guild.ID "/" $combat "/726465578809688135)")}}
		{{$message := getMessage $combat 726465578809688135}}
		{{$embed := cembed
			"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"description" (joinStr "" $message.Content "\n\n [➥ Original]" $link "\n")
			"icon_url" (.User.AvatarURL "512")}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "?shop") (eq $commande "?store")}}
		{{$link = (print "(https://discordapp.com/channels/" .Guild.ID "/734838748721840188/736713337571901570)")}}
		{{$message := getMessage 734838748721840188 736713337571901570}}
		{{$embed := cembed
			"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"description" (joinStr "" $message.Content "\n\n [➥ Original]" $link "\n")
			"icon_url" (.User.AvatarURL "512")}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if eq $commande "?pa" "?PA"}}
		{{$link = print "(https://discordapp.com/channels/" .Guild.ID "/" $combat "/739903306683646105)"}}
		{{$message := getMessage $combat 739903306683646105}}
		{{$embed := cembed
			"author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "512"))
			"color" $col
			"description" (joinStr "" $message.Content "\n\n [➥ Original]" $link "\n")
			"icon_url" (.User.AvatarURL "512")}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}

	{{else if or (eq $commande "?help") (eq $commande "?snippet") (eq $commande "?all") }}
		{{$embed := cembed
			"title" "Liste des aides disponibles"
		"description" ":white_small_square: `?(armes|arme)`\n:white_small_square: `?résumé`\n:white_small_square: `?dés`\n:white_small_square: `?position`\n:white_small_square: `?(dégâts|dégât|dégat|dégats)`\n:white_small_square: `?notes`\n:white_small_square: `?horloge`\n:white_small_square: `?bouclier`\n:white_small_square:`?ticket`\n:white_small_square: `?charge`\n:white_small_square: `?shop`\n:white_small_square:`?(pa|PA)`\n\n:white_medium_square: **Pour afficher cette liste** : `?(all|snippet|help)`"}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}
	{{end}}
{{end}}
