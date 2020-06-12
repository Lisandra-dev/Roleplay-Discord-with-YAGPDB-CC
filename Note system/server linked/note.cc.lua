{{$args := parseArgs 2 ""
  (carg "string" "key")
  (carg "string" "value")}}

{{dbSet 0 (joinStr "" "notes_" ($args.Get 0)) ($args.Get 1)}}
*Note `{{$args.Get 0}}` sauvegardée.*
{{deleteTrigger 2}}
{{deleteResponse 2}}
