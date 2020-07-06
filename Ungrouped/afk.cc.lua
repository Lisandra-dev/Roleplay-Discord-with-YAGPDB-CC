{{/*
	This command allows users to set an AFK message with optional duration.
	Usage:

	$afk -d <duration> <reason>

	Recommended trigger: Command trigger with trigger `-afk`
*/}}

{{$duration := 0}}
{{$msg := ""}}
{{$id := ""}}
{{$idrole := }}
{{$chan := }}


{{ if .CmdArgs }}
	{{if eq (toFloat (len .CmdArgs)) (toFloat 1)}}
		{{$msg = .CmdArgs}}
	{{else if gt (toFloat (len .CmdArgs)) (toFloat 1)}}
		{{if eq (toFloat (len .CmdArgs)) (toFloat 2)}}
			{{if eq (index .CmdArgs 0) "-d"}}
				{{$duration = (index .CmdArgs 1)}}
				{{if eq (toInt $duration) (toInt 0)}}
					"Mauvaise durée d'absence."
				{{end}}
				{{dbSetExpire .User.ID "afk" "" (toInt $duration)}}
			{{else}}
				{{$msg = (joinStr " " "car :" .CmdArgs)}}
			{{end}}
		{{else}}
			{{if eq (index .CmdArgs 0) "-d"}}
				{{$duration = (index .CmdArgs 1)}}
				{{if eq (toInt $duration) (toInt 0)}}
					"Mauvaise durée d'absence."
				{{end}}
				{{$msg = (joinStr " " "car :" (slice .CmdArgs 2))}}
				{{dbSetExpire .User.ID "afk" $msg (toInt $duration)}}
			{{end}}
		{{end}}
	{{end}}
		{{$embed := cembed
			"color" (randInt 111111 999999)
			"author" (sdict "name" (printf "%s est AFK" .User.Username) "icon_url" (.User.AvatarURL "256"))
			"description" (joinStr " " "Durant"  (toDuration (joinStr "" $duration "s" )) $msg)}}
		{{$id = sendMessageRetID $chan $embed }}
		{{deleteMessage $chan $id $duration}}
	{{addRoleID $idrole}}
	{{removeRoleID $idrole $duration}}
{{end}}

{{if hasRoleID $idrole}}
	{{removeRoleID $idrole}}
	Vous n'êtes plus AFK
	{{deleteMessage $chan $id $duration}}
	{{dbDel .User.ID "afk"}}
{{end}}