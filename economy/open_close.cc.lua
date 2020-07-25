{{if .CmdArgs}}
	{{if eq (index .CmdArgs 0) "open"}}
		{{dbSet .Server.ID "store" "open"}}
		{{$embed := cembed
		"author" (sdict "name" "Sola-UI" "icon_url" "https://cdn.discordapp.com/attachments/726496591489400872/727978845185245283/download20200605012708.png")
		"description" "Un vaisseau marchand vous contact sur la ligne de communication. Il vous ouvre sa boutique dans un canal privé, n'hésitez pas à aller leur répondre."
		"color" 0xD5B882}}}}
		{{sendMessage nil $embed}}
	{{else}}
		{{dbSet .Server.ID "store" "close"}}
		{{dbSet .Server.ID "store" "open"}}
		{{$embed := cembed
		"author" (sdict "name" "Sola-UI" "icon_url" "https://cdn.discordapp.com/attachments/726496591489400872/727978845185245283/download20200605012708.png")
		"description" "Le vaisseau marchand vient de partir."
		"color" 0xD5B882}}
		{{sendMessage nil $embed}}
	{{end}}
{{end}}
{{$store := (dbGet .Server.ID "store").Value}}
