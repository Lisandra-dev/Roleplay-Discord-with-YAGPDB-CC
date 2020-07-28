{{/* Channel */}}
{{$chan:= 716988208205791342}} {{/* Change the channel ID you want send the embed */}}

{{/*Thumbnail*/}}
{{$nuit:="https://i.imgur.com/e04keB7.png"}}
{{$matin:="https://i.imgur.com/ZB5yT5s.png"}}
{{$midi:="https://i.imgur.com/AFOj90o.png"}}
{{$soir:="https://i.imgur.com/xSDYgqD.png"}}

{{/*All message database counter*/}}
{{$count := "count"}}

{{/*Timer message */}}
{{$msgc := (toFloat (dbGet 0 "mgsc").Value)}}

{{/*Message counter*/}}
{{$message:= "message"}}

{{/*Cycles counter*/}}
{{$time := "time"}}

{{/*Get amound of cycles*/}}
{{$number:= (dbGet 0 $time).Value}}

{{/*Check if massage sent is in (brackets)*/}}
{{ $matches := (index (reFindAllSubmatches `\((.*)\)|(^\$(.*))|(^\!\!(.+))` .Message.Content)) }} {{/* exclude my bot bracket from the count */}}

{{/*If not*/}}
{{if not $matches}}
  {{$x := (toFloat (dbGet 0 $count).Value)}}
  {{$silent := dbIncr 0 $count 1}}
  {{$msg := dbIncr 0 $message 1}}
  {{$jour:= (toString (toInt (dbGet 0 "jour").Value))}}

 {{/*Amount of messages needed for cycle, starts from 0*/}}
    	{{if eq $x $msgc }} {{/*Change here for the amount of message */}}
       		{{$silent := dbIncr 0 $time 1}}

{{/*reset count for new cycle*/}}
  		{{dbSet 0 $count 0}}
  		{{dbSet 0 $message 0}}
      {{$y := (toFloat (dbGet 0 $time).Value)}}

{{/*thumbnail cycle*/}}
  		{{if eq $y (toFloat 1)}}
  			{{$embed := cembed
    				"title" (joinStr "" "Jour : " $jour " ▬ Changement de cycle")
    				"description" (joinStr "" "Nous sommes maintenant au cycle 1.")
    				"color" 0x1B3175
  				"thumbnail" (sdict "url" $nuit)
  				"timestamp" .Message.Timestamp}}
          		{{sendMessage $chan $embed}}

  		{{else if eq $y (toFloat 2)}}
  			{{$embed := cembed
    				"title" (joinStr "" "Jour : " $jour " ▬ Changement de cycle")
    				"description" (joinStr "" "Nous sommes maintenant au cycle 2.")
    				"color" 0xDD99DF
  				"thumbnail" (sdict "url" $matin)
  				"timestamp" .Message.Timestamp}}
          		{{sendMessage $chan $embed}}

  		{{else if eq $y (toFloat 3)}}
  			{{$embed := cembed
    				"title" (joinStr "" "Jour : " $jour " ▬ Changement de cycle")
    				"description" (joinStr "" "Nous sommes maintenant au cycle 3.")
    				"color" 0xF0B535
  				"thumbnail" (sdict "url" $midi)
  				"timestamp" .Message.Timestamp}}
          		{{sendMessage $chan $embed}}

  		{{else if eq $y (toFloat 4)}}
  			{{$embed := cembed
    				"title" (joinStr "" "Jour : " $jour " ▬ Changement de cycle")
    				"description" (joinStr "" "Nous sommes maintenant au cycle 4.")
    				"color" 0x9593E8
  				"thumbnail" (sdict "url" $soir)
  				"timestamp" .Message.Timestamp}}
         			{{sendMessage $chan $embed}}

              {{/*you can add cycle with just else if eq $y (toFloat number) */}}

  {{/*Nouveau jour*/}}
     		{{else}}
  			{{dbSet 0 $count 0 }}
  			{{dbSet 0 $time 1}}
  			{{$sil:= dbIncr 0 "jour" 1}}
  			{{$jour:= (toString (toInt (dbGet 0 "jour").Value))}}
  			{{$embed := cembed
   		 		"title" (joinStr "" "Début du Jour : " $jour)
  				"description" "Nous sommes maintenant au cycle 1"
    				"color" 0x1B3175
  				"thumbnail" (sdict "url" $nuit)
  				"timestamp" .Message.Timestamp}}
  			{{sendMessage $chan $embed}}
    	{{end}}
	{{end}}
{{else}}
{{end}}
