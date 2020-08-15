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

{{$commande := reFind `^\?(shop|(charg(é|e)(s|r)?)|(recharg(é|e)(s|r)?)|help|snippet|aide|all|d(é|e)s?|dices?|r(e|é)sum(e|é)s?|armes?|d(é|e)g(a|â)s?|store|PA|pa|jdr|recensement|recens|perso|char(list?)|template)` .Message.Content}}

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

	{{else if eq $commande "?jdr"}}
		→ https://www.jdr-system.ovh/

		{{else if eq $commande "?recensement" "?recens" "?perso" "?char" "?charlist"}}
			→ https://docs.google.com/spreadsheets/d/1k_7glSefzeAqWCFu9F3lWfYfCEw4cIq_ijWz2z-PwnU/edit?usp=sharing

		{{else if eq $commande "?template"}}
			→ https://docs.google.com/document/d/1CX4ye8loV4d34BOwmRlb-nOt8SMA4VvbFy4-_MHX5A8/edit?usp=sharing
			→ Fichier : Créer une copie
			→ __Remplacer l'image__ : clic droit (sur l'image) → Remplacer l'image.
			→ __ Remplir une case__ : Pot de peinture


	{{else if or (eq $commande "?help") (eq $commande "?snippet") (eq $commande "?all") }}
		{{$embed := cembed
			"title" "Liste des aides disponibles"
		"description" ":white_small_square: `?(armes|arme)`\n:white_small_square: `?résumé`\n:white_small_square: `?dés`\n:white_small_square: `?position`\n:white_small_square: `?(dégâts|dégât|dégat|dégats)`\n:white_small_square: `?charge`\n:white_small_square: `?shop`\n:white_small_square:`?(pa|PA)`\n\n:white_medium_square: **Pour afficher cette liste** : `?(all|snippet|help)`"}}
		{{sendMessage nil $embed}}
		{{deleteTrigger 1}}
	{{end}}
{{end}}
{{deleteTrigger 1}}
