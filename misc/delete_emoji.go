{{ $matches := reFindAllSubmatches `((<a?:[\w~]{2,32}:\d{17,19}>)|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})` .Message.Content }}
{{if eq (len (index (index $matches 0) 0)) (len .Message.Content) }}
	{{deleteTrigger 1}}
{{end}}