{{$msg:= (toString (toInt (dbGet 0 "count").Value))}}
{{$cycle := (toString (toInt (dbGet 0 "time").Value))}}
{{$day := (toString (toInt (dbGet 0 "jour").Value))}}
{{$speed := (toString (toInt (dbGet 0 "mgsc").Value))}}

{{$embed := cembed
"title" "Paramètre actuel des cycles"
"description" (joinStr " " ":white_small_square: **Cycle actuel :**" $cycle "\n :white_small_square: **Nombre de message dans le cycle :**" $msg "\n :white_small_square: **Jour actuel :**" $day "\n :white_small_square: **Nombre de message fixé à :**" $speed)
"color" 0x1B3175}}
{{sendMessage nil $embed}}

{{deleteTrigger 1}}
