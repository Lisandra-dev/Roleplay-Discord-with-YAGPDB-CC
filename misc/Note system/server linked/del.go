{{$key := joinStr "" "notes_"  .StrippedMsg}}
{{$note := dbGet 0 $key}}
{{if $note}}
{{$strippedKey := slice $key 6 (len $key)}}
{{dbDel 0 $key}}
Note {{$strippedKey}} supprimée.
{{else}}Cette note n'existe pas :<{{end}}
{{deleteTrigger}}

 #37 - Command: noteperso
