{{$count := "count"}}
{{$time := "time"}}
{{$day := "jour"}}
{{$message := "message"}}
{{$embed := cembed
    "title" "Reset ! "
    "color" 0x670404
    "timestamp" .Message.Timestamp}}
{{ sendMessage nil $embed}}

{{dbSet 0 $count 0}}
{{dbSet 0 $time 1}}
{{dbSet 0 $message 0}}
{{dbSet 0 $day 1}}
