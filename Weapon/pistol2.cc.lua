{{/* Each time the bot sees the trigger, it will count until it reaches the value set in the "if lt".
It will also count the number of balls used, and will return this message to tell the user that it has no more balls.

If you change the value of the if, you must change the value in the "$x := sub".  */}}



{{if not (dbGet .User.ID "pistol2")}}
 	{{dbSet .User.ID "pistol2" 0}}
 	{{dbSet 0 "chan" (toString .Channel.ID)}}
 	{{$incr := dbIncr .User.ID "pistol2" 1}}
 	{{$y := (dbGet .User.ID "pistol2").Value}}
	{{$x := sub 8 $y}}
 	{{if lt $y (toFloat 7)}}
      		{{ $embed := cembed
       			"description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $x)) " charges dans votre pistolet n°2 !")
		}}
      		{{ $id := sendMessageRetID nil $embed }}
  		{{deleteMessage nil $id 30}}
 	{{else}}
    		{{ $embed := cembed
    			"description" "Votre pistolet N°2 est vide."
    		}}
    		{{ $id := sendMessageRetID nil $embed }}
    		{{deleteMessage nil $id 30}}
  	{{end}}

{{else}}
	{{$incr := dbIncr .User.ID "pistol2" 1}}
	{{$y := (dbGet .User.ID "pistol2").Value}}
  	{{$x := sub 8 $y}}
  	{{if lt $y (toFloat 8)}}
    		{{ $embed := cembed
      			"description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $x)) " charges de pistolet n°2 !")
    		}}
    		{{ $id := sendMessageRetID nil $embed }}
    		{{deleteMessage nil $id 30}}
  	{{else}}
   		 {{ $embed := cembed
    			"description" (joinStr "" .User.Mention ", votre pistolet n°2 est vide. ")
   	 	}}
    		{{ $id := sendMessageRetID nil $embed }}
    		{{deleteMessage nil $id 30}}
  	{{end}}
{{end}}
