{{$message := "message"}}
{{$time := "time"}}
{{$jour := "jour"}}
{{$number:= (dbGet 0 $time).Value}}
{{/* $msg := div (dbGet 0 $message).Value 2 */}} {{/* for 200 messages, create a real pourcentage. */}}
{{$msg := (dbGet 0 $message).Value}}

{{$y := (toFloat (dbGet 0 $time).Value)}}
{{$txt := (dbGet 0 $message).Value }}
{{$val := (joinStr " " (toString (toInt $txt)) "message(s) dans le cycle")}} {{/* Footer message */}}
{{$day := (toString (toInt (dbGet 0 $jour).Value))}}

{{/* Thumbnail */}}
{{$nuit:="https://cdn.discordapp.com/attachments/717422988038897744/717423036713795625/Citycons_night_icon-icons.com_67936.png"}}
{{$matin:="https://cdn.discordapp.com/attachments/717422988038897744/717423356550447204/morning-3867006_640.png"}}
{{$midi:="https://cdn.discordapp.com/attachments/717422988038897744/717423822705262662/vippng.com-soleil-png-3861240.png"}}
{{$soir:="https://cdn.discordapp.com/attachments/717422988038897744/717424480959332452/-night-scenery_89739.png"}}

{{if eq $y (toFloat 1) }}
	{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" (toString (toInt $msg)) " % du cycle est passé...")
  		"color" 0x1B3175
 		"thumbnail" (sdict "url" $nuit)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}

{{else if eq $y (toFloat 2) }}
		{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" (toString (toInt $msg)) " % du cycle est passé...")
  		"color" 0xDD99DF
 		"thumbnail" (sdict "url" $matin)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}

{{else if eq $y (toFloat 3) }}
	{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" (toString (toInt $msg)) " % du cycle est passé...")
  		"color" 0xF0B535
 		"thumbnail" (sdict "url" $midi)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}

{{/* If you want add cycle, add it here ! */}}

{{else}}
	{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" (toString (toInt $msg)) " % du cycle est passé...")
  		"color" 0x9593E8
 		"thumbnail" (sdict "url" $soir)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}
{{end}}

{{deleteTrigger 0}}
