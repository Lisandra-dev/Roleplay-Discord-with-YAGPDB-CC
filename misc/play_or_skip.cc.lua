{{if and (eq .ReactionMessage.Author.ID 204255221017214977) (eq .ReactionMessage.ID 735106303877185556)}}
	{{if .ReactionAdded}}
		{{if eq .Reaction.Emoji.Name "⏯️" }}
			{{dbSet .User.ID "userplay" 1}}
		{{else if eq .Reaction.Emoji.Name "⏩"}}
			{{dbSet .User.ID "userskip" 1}}
		{{end}}
	{{else if not .ReactionAdded}}
		{{if (dbGet .User.ID "userskip")}}
			{{dbDel .User.ID "userskip"}}
		{{else if (dbGet .User.ID "userplay")}}
			{{dbDel .User.ID "userplay"}}
		{{end}}
	{{end}}
{{end}}

{{$s := ""}}
{{$p := ""}}
{{range (dbTopEntries "userskip" 100 0)}}
	{{$s = (joinStr "" $s " <:tr:724626754282717194>:small_blue_diamond:" .User.Mention "\n")}}
{{end}}

{{range (dbTopEntries "userplay" 100 0)}}
	{{$p = (joinStr "" $p " <:tr:724626754282717194>:small_blue_diamond:" .User.Mention "\n")}}
{{end}}

{{$embed := cembed
"title" "Absence"
"description" (joinStr "" "⏩ | **Skip** : Autorise les MJ à sauter son tour\n" $s "\n ⏯️ | **Jouer** : Autorise les MJ à jouer son personnage\n" $p)
"color" 0x8062BC
"footer" (sdict "text" "Action que les joueurs autorisent aux MJ lorsqu'ils sont absents | Edit")
"thumbnail" (sdict "url" "https://i.imgur.com/XUQJXyO.png")
"timestamp" currentTime}}

{{editMessage nil .ReactionMessage.ID (complexMessageEdit "embed" $embed "content" "")}}
