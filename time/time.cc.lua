{{$message := "message"}}
{{$time := "time"}}
{{$jour := "jour"}}
{{$number:= (dbGet 0 $time).Value}}

{{$msgc := (toFloat (dbGet 0 "mgsc").Value)}}
{{$pourcent := div $msgc 100}}
{{if lt $pourcent (toFloat 1)}}
	{{$pourcent = (toFloat 1)}}
{{end}}
{{ $msg := (json (div (dbGet 0 $message).Value $pourcent)) }}


{{$y := (toFloat (dbGet 0 $time).Value)}}
{{$txt := (dbGet 0 $message).Value }}
{{$val := (joinStr " " (toString (toInt $txt)) "message(s) dans le cycle")}} {{/* Footer message */}}
{{$day := (toString (toInt (dbGet 0 $jour).Value))}}

{{/* Thumbnail */}}
{{$nuit:="https://i.imgur.com/e04keB7.png"}}
{{$matin:="https://i.imgur.com/ZB5yT5s.png"}}
{{$midi:="https://i.imgur.com/AFOj90o.png"}}
{{$soir:="https://i.imgur.com/xSDYgqD.png"}}

{{if eq $y (toFloat 1) }}
	{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" $msg " % du cycle est passé...")
  		"color" 0x1B3175
 		"thumbnail" (sdict "url" $nuit)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}

{{else if eq $y (toFloat 2) }}
		{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" $msg " % du cycle est passé...")
  		"color" 0xDD99DF
 		"thumbnail" (sdict "url" $matin)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}

{{else if eq $y (toFloat 3) }}
	{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" $msg " % du cycle est passé...")
  		"color" 0xF0B535
 		"thumbnail" (sdict "url" $midi)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}

{{/* If you want add cycle, add it here ! */}}

{{else}}
	{{$embed := cembed
		"title" (joinStr "" "Jour : " $day)
  		"description"  (joinStr "" "**Cycle** : " (toString (toInt $number)) "\n" $msg " % du cycle est passé...")
  		"color" 0x9593E8
 		"thumbnail" (sdict "url" $soir)
		"footer" (sdict "text" $val)
		"timestamp" .Message.Timestamp}}
	{{ sendMessage nil $embed}}
{{end}}

{{deleteTrigger 0}}
