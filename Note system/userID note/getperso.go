{{$key := joinStr "" "notes_"  .StrippedMsg}}
{{$note := dbGet .User.ID $key}}
{{if $note}}

{{$strippedKey := slice $key 6 (len $key)}}
{{sendDM  $strippedKey ":\n" $note.Value}}

{{else}}Cette note n'existe pas :<{{end}}
