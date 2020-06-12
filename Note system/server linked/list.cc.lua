{{$notes := dbGetPattern 0 "notes_%" 100 0}}
**Notes présentes sur le serveur** :{{range $notes}}{{- $strippedKey := slice .Key 6 (len .Key)}}
> :white_small_square: {{$strippedKey}}
{{- else}}
Il n'y a pas de note sur le serveur.
{{end}}
