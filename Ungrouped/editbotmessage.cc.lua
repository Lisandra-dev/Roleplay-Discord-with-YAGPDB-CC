{{$args := parseArgs 3 ""
  (carg "int" "channel-id")
  (carg "int" "message-id")
  (carg "string" "new-message")}}

{{editMessageNoEscape  ($args.Get 0) ($args.Get 1) ($args.Get 2)}}
{{deleteTrigger 1}}
