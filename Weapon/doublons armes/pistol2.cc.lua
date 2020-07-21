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


{{if not (dbGet .User.ID "pistol2")}}
 	{{dbSet .User.ID "pistol2" 0}}
 	{{dbSet 0 "chan" (toString .Channel.ID)}}
 	{{$incr := dbIncr .User.ID "pistol2" 1}}
 	{{$y := (dbGet .User.ID "pistol2").Value}}
	{{$x := sub 8 $y}}
 	{{if lt $y (toFloat 7)}}
		{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/8 charges dans votre deuxième pistolet !")}}
    {{ $id := sendMessageRetID nil $embed }}
  		{{deleteMessage nil $id 30}}
 	{{else}}
		{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" "Votre deuxième pistolet est vide."}}
		{{ $id := sendMessageRetID nil $embed }}
    {{deleteMessage nil $id 30}}
  	{{end}}

{{else}}
	{{$incr := dbIncr .User.ID "pistol2" 1}}
	{{$y := (dbGet .User.ID "pistol2").Value}}
  	{{$x := sub 8 $y}}
  	{{if lt $y (toFloat 8)}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/8 charges dans votre deuxième pistolet !")}}
    	{{ $id := sendMessageRetID nil $embed }}
    	{{deleteMessage nil $id 30}}
  	{{else}}
			{{ $embed := cembed
			 "author" (sdict "name" $user "icon_url" $img)
			 "color" 0x6CAB8E
			 "description" "Votre deuxième pistolet est vide."}}
    	{{ $id := sendMessageRetID nil $embed }}
    	{{deleteMessage nil $id 30}}
  	{{end}}
{{end}}
