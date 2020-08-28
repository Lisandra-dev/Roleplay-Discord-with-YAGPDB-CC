{{/* Channel */}}
{{$chan:= 716988208205791342}} {{/* Change the channel ID you want send the embed */}}

{{/*Thumbnail*/}}
{{$nuit:="https://i.imgur.com/e04keB7.png"}}
{{$matin:="https://i.imgur.com/ZB5yT5s.png"}}
{{$midi:="https://i.imgur.com/AFOj90o.png"}}
{{$soir:="https://i.imgur.com/xSDYgqD.png"}}

{{/*All message database counter*/}}
{{$time := sdict }}
{{with (dbGet 0 "time")}}
	{{$time = sdict .Value}}
{{end}}

{{/* Variable */}}
{{$msgc := (toFloat ($time.Get "mgsc"))}}
{{$message := toFloat ($time.Get "message")}}
{{$number:= ($time.Get "time")}}
{{$jour := $time.Get "jour"}}
{{$cycle := $time.Get "cycle"}}

{{/*Check if massage sent is in (brackets)*/}}
{{ $matches := (index (reFindAllSubmatches `\((.*)\)|(^\$(.*))|(^\!\!(.+))` .Message.Content)) }} {{/* exclude my bot bracket from the count */}}

{{/*If not*/}}
{{if not $matches}}
  {{$x := (toFloat ($time.Get "count"))}}
  {{$time.Set "count" (add $x 1)}}
  {{$time.Set "message" (add $message 1)}}

 {{/*Amount of messages needed for cycle, starts from 0*/}}
    	{{if eq $x $msgc }} {{/*Change here for the amount of message */}}
       		{{$time.Set "cycle" (add $cycle 1)}}

{{/*reset count for new cycle*/}}
  		{{$time.Set "count" 0}}
  		{{$time.Set "message" 0}}
      {{$y := toFloat ($time.Get "cycle")}}

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
        {sendMessage $chan $embed}}

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
  			{{$time.Set "count" 0 }}
  			{{$time.Set "cycle" 1}}
  			{{$time.Set "jour" (add $jour 1)}}
  			{{$jour:= $time.Get "jour"}}
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
{{dbSet 0 "time" $time}}
