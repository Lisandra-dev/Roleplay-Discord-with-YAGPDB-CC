{{/*Cycles counter*/}}
{{$time := sdict }}
{{with (dbGet 0 "time")}}
	{{$time = sdict .Value}}
{{end}}
{{$cycle := $time.Get "cycle"}}
{{/*Get amound of cycles*/}}
{{ editChannelName 716988208205791342 (joinStr " " "Cycle" (toString (toInt $cycle))) }}
{{/* Don't forget the channel ID !*/}}
