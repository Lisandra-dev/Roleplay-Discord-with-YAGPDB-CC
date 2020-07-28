{{$notes := dbGetPattern .User.ID "notes_%" 100 0}}
{{range $notes}}{{- $strippedKey := slice .Key 6 (len .Key)}}
{{sendDM "**Note** :\n :white_small_square:" $strippedKey}}
{{- else}}
Vous n'avez aucune note :<
{{end}}
{{deleteTrigger 0}}
