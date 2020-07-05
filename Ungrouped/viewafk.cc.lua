{{/*
	This command manages mentioning of users who are AFK.

	Recommended trigger: Regex trigger with trigger `<@!?\d+>`
*/}}
{{ $id := reFind `\d+` .Cmd | toInt64 }}
{{ with (dbGet $id "afk") }}
	{{ $user := userArg .UserID }}
	{{ $eta := "" }}
	{{ if gt .ExpiresAt.Unix 0 }} {{ $eta = humanizeDurationSeconds (.ExpiresAt.Sub currentTime) | printf "*%s sera de retour dans : %s.*" $user.Username }} {{ end }}
	{{ sendMessage nil (cembed
		"author" (sdict "name" (printf "%s est AFK" $user.String) "icon_url" ($user.AvatarURL "256"))
		"description" (joinStr "\n\n" $eta .Value "\n" (dbGet $id "afk"))
		"color" (randInt 0 16777216)
		"footer" (sdict "text" "Away since")
		"timestamp" .UpdatedAt
	) }}
{{ end }}
