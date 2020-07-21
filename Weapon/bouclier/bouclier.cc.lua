{{$icon := "https://i.imgur.com/FdD2UqK.png"}}
{{if not (dbGet 0 "plasma")}}
      {{dbSet 0 "plasma" 1}}
	{{ $embed := cembed
		"color" 0x88D9EE
		"description" (joinStr "" "Début du rechargement bouclier plasma")
		"author" (sdict "name" "Vaisseau Nucleus" "icon_url" $icon)}}
	{{ $id := sendMessageRetID 701395027896565810 $embed }}
	{{dbSet 0 "bouclierID" (str $id)}}
    {{else}}
     	{{ $embed := cembed
			"color" 0x88D9EE
			"description" "Le bouclier est déjà en cours de rechargement"
			"author" (sdict "name" "Vaisseau Nucleus" "icon_url" $icon)}}
	{{ sendMessage nil $embed }}
    {{end}}
