{{/* v1 */}}
{{ $matches := reFindAllSubmatches `https://discordapp.com/channels\/(\d+)\/(\d+)\/(\d+)` .Message.Content }}
{{$msg := getMessage (index (index $matches 0) 2) (index (index $matches 0) 3) }}
{{if not $msg}}

{{else}}
{{ $avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString $msg.Author.ID) "/" $msg.Author ".png") }}

{{$embedRaw := sdict
"description" (joinStr "" $msg.Content "\n" "*[➥ Lien vers l'original](" (index (index $matches 0) 0) ")*  ")
"color" 4645612
"author" (sdict "name"  $msg.Author.Username "icon_url" ($msg.Author.AvatarURL "64"))
"footer" (sdict "text" (joinStr "" "Cité par "  .Message.Author.Username ""))
"timestamp" $msg.Timestamp  }}

{{if $msg.Attachments}}
{{$embedRaw.Set "image" (sdict "url" (index $msg.Attachments 0).URL) }}
{{end}}

{{ sendMessage nil (cembed $embedRaw) }}

{{/* delete the trigger if it only contained a link and nothing more */}}
{{if eq (len (index (index $matches 0) 0)) (len .Message.Content) }} {{deleteTrigger 1}} {{end}}
{{end}}
