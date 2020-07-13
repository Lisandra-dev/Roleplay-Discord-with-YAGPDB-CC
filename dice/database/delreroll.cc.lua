{{$name := toRune (lower (index .CmdArgs 0))}}
{{$user :=""}}
{{range $name}}
{{- $user = (print $user .)}}
{{- end}}
{{$user = (toInt $user)}}

{{dbDel $user "force"}}
{{dbDel $user "i_force"}}
{{dbDel $user "agi"}}
{{dbDel $user "i_agi"}}
{{dbDel $user "preci"}}
{{dbDel $user "i_preci"}}
{{dbDel $user "intelligence"}}
{{dbDel $user "i_intel" }}
{{dbDel $user "karma"}}
