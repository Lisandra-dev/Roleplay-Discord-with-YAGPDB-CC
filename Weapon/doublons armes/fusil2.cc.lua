{{/* Each time the bot sees the trigger, it will count until it reaches the value set in the "if lt".
It will also count the number of balls used, and will return this message to tell the user that it has no more balls.

If you change the value of the if, you must change the value in the "$x := sub".  */}}



{{if not (dbGet .User.ID "fusil2")}}
	{{dbSet .User.ID "fusil2" 0}}
 	{{$incr := dbIncr .User.ID "fusil2" 1}}
	{{$y := (dbGet .User.ID "fusil2").Value}}
	{{$x := sub 12 $y}}
	{{if lt $y (toFloat 12)}}
		{{ $embed := cembed
		"description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $x)) " charges dans votre deuxième fusil !")}}
		{{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
	{{else}}
		{{ $embed := cembed
		"description" "Votre fusil2 est vide."}}
		{{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
	{{end}}
{{else}}
 	{{$incr := dbIncr .User.ID "fusil2" 1}}
	{{$y := (dbGet .User.ID "fusil2").Value}}
	{{$x := sub 12 $y}}
	{{if lt $y (toFloat 12)}}
		{{ $embed := cembed
			"description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $x)) " charges dans votre deuxième fusil !")}}
		{{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
	{{else}}
		{{ $embed := cembed
		"description" (joinStr "" .User.Mention ", votre deuxième fusil est vide !")}}
		{{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
	{{end}}
{{end}}
