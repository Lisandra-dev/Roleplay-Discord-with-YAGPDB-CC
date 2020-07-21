{{/* Each time the bot sees the trigger, it will count until it reaches the value set in the "if lt".
It will also count the number of balls used, and will return this message to tell the user that it has no more balls.

If you change the value of the if, you must change the value in the "$x := sub".  */}}

{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{if $name}}
	{{$user = $name}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$img := "https://i.imgur.com/YeIsRmw.png"}}

{{if not (dbGet .User.ID "fusil")}}
	{{dbSet .User.ID "fusil" 0}}
 	{{$incr := dbIncr .User.ID "fusil" 1}}
	{{$y := (dbGet .User.ID "fusil").Value}}
	{{$x := sub 12 $y}}
	{{if lt $y (toFloat 12)}}
		{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/12 charges de fusil !")}}
		{{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
	{{else}}
		{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" "Votre fusil est vide..."}}
		{{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
	{{end}}
{{else}}
 	{{$incr := dbIncr .User.ID "fusil" 1}}
	{{$y := (dbGet .User.ID "fusil").Value}}
	{{$x := sub 12 $y}}
	{{if lt $y (toFloat 12)}}
	{{ $embed := cembed
		"author" (sdict "name" $user "icon_url" $img)
		"color" 0x6CAB8E
		"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/12 charges de fusil !")}}
		{{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
	{{else}}
		{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" "Votre fusil est vide..."}}
		{{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
	{{end}}
{{end}}
