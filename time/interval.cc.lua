{{/*Cycles counter*/}}
{{$time := "time"}}
{{/*Get amound of cycles*/}}
{{$number:= (dbGet 0 $time).Value}}
{{ editChannelName 716988208205791342 (joinStr "" "Cycle " (toString (toInt $number)) ".") }}
{{/* Don't forget the channel ID !}}
