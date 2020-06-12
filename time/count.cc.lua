{{/* Channel */}}
{{$chan:= 716988208205791342}} {{/* Change the channel ID you want send the embed */}}

{{/*Thumbnail*/}}
{{$nuit:="https://cdn.discordapp.com/attachments/717422988038897744/717423036713795625/Citycons_night_icon-icons.com_67936.png"}}
{{$matin:="https://cdn.discordapp.com/attachments/717422988038897744/717423356550447204/morning-3867006_640.png"}}
{{$midi:="https://cdn.discordapp.com/attachments/717422988038897744/717423822705262662/vippng.com-soleil-png-3861240.png"}}
{{$soir:="https://cdn.discordapp.com/attachments/717422988038897744/717424480959332452/-night-scenery_89739.png"}}

{{/*All message database counter*/}}
{{$count := "count"}}

{{/*Message counter*/}}
{{$message:= "message"}}

{{/*Cycles counter*/}}
{{$time := "time"}}

{{/*Get amound of cycles*/}}
{{$number:= (dbGet 0 $time).Value}}

{{/*Check if massage sent is in (brackets)*/}}
{{ $matches := (index (reFindAllSubmatches `\((.*?)\)` .Message.Content)) }}

{{/*If not*/}}
{{if not $matches}}
    	{{$x := (toFloat (dbGet 0 $count).Value)}}
    	{{$silent := dbIncr 0 $count 1}}
      {{$msg := dbIncr 0 $message 1}}

 {{/*Amount of messages needed for cycle, starts from 0*/}}
    	{{if eq $x (toFloat 100) }} {{/*Change here for the amount of message */}}
       		{{$silent := dbIncr 0 $time 1}}

{{/*reset count for new cycle*/}}
		{{dbSet 0 $count 0}}
		{{dbSet 0 $message 0}}
    		{{$y := (toFloat (dbGet 0 $time).Value)}}

{{/*thumbnail cycle*/}}
		{{if eq $y (toFloat 1)}}
			{{$embed := cembed
  				"title" "Changement de cycle"
  				"description" (joinStr "" "Nous sommes maintenant au cycle 1.")
  				"color" 0x1B3175
				"thumbnail" (sdict "url" $nuit)
				"timestamp" .Message.Timestamp}}
        		{{sendMessage $chan $embed}}

		{{else if eq $y (toFloat 2)}}
			{{$embed := cembed
  				"title" "Changement de cycle"
  				"description" (joinStr "" "Nous sommes maintenant au cycle 2.")
  				"color" 0xDD99DF
				"thumbnail" (sdict "url" $matin)
				"timestamp" .Message.Timestamp}}
        		{{sendMessage $chan $embed}}

		{{else if eq $y (toFloat 3)}}
			{{$embed := cembed
  				"title" "Changement de cycle"
  				"description" (joinStr "" "Nous sommes maintenant au cycle 3.")
  				"color" 0xF0B535
				"thumbnail" (sdict "url" $midi)
				"timestamp" .Message.Timestamp}}
        		{{sendMessage $chan $embed}}

		{{else if eq $y (toFloat 4)}}
			{{$embed := cembed
  				"title" "Changement de cycle"
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
			{{$embed := cembed
 		 		"title" "Nouveau jour !"
				"description" "Nous sommes maintenant au cycle 1"
  				"color" 0x1B3175
				"thumbnail" (sdict "url" $nuit)
				"timestamp" .Message.Timestamp}}
			{{sendMessage $chan $embed}}
    		{{end}}
	{{end}}
{{else}}
{{end}}
