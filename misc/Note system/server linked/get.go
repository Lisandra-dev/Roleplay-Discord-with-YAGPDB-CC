{{$key := joinStr "" "notes_"  .StrippedMsg}}
{{$note := dbGet 0 $key}}
{{if $note}}

{{$strippedKey := slice $key 6 (len $key)}}
**{{$strippedKey}}**:
{{$note.Value}}
{{else}}Cette note n'existe pas :<{{end}}
{{deleteTrigger 0}}
