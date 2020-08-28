{{/*All message database counter*/}}
{{$time := sdict }}
{{with (dbGet 0 "time")}}
	{{$time = sdict .Value}}
{{end}}

{{$amount := toFloat ($time.Get "mgsc")}}

{{if .CmdArgs}}

    {{if eq (index .CmdArgs 0) "cycle"}}
        {{if eq (len .CmdArgs) 2}}
            {{$floatArg := (toFloat (index .CmdArgs 1)) }}
            {{$arg1 := (toFloat (index .CmdArgs 1)) }}
            {{if ge (toFloat 4) $floatArg }}
                {{$time.Set "cycle" $arg1}}
                Nous sommes actuellement au cycle  `{{$time.Get "cycle"}}`.
            {{else}}
              Le nombre est supérieur à 4.
            {{end}}
        {{else}}
            **Usage:** cycle <number>
        {{end}}

    {{else if eq (index .CmdArgs 0) "speed"}}
      {{if eq (len .CmdArgs) 2}}
          {{$time.Set "mgsc" (index .CmdArgs 1)}}
          Le nombre de message pour passer un cycle est maintenant fixé à : `{{(index .CmdArgs 1)}}`.
      {{else}}
          **Usage** : timer <number>
      {{end}}

    {{else if eq (index .CmdArgs 0) "day"}}
      {{if eq (len .CmdArgs) 2}}
        {{$floatArg := (toFloat (index .CmdArgs 1)) }}
        {{$time.Set "jour" (toFloat (index .CmdArgs 1)) }}
        {{$jour := (toString (toInt ($time.Get "jour")))}}
        Nous sommes maintenant au jour : `{{$jour}}`.
      {{else}}
        **Usage** day <number>
      {{end}}

    {{else if eq (index .CmdArgs 0) "msg"}}
        {{if eq (len .CmdArgs) 2}}
            {{$floatArg := (toFloat (index .CmdArgs 1)) }}
            {{$arg1 := (toFloat (index .CmdArgs 1)) }}
            {{if ge $amount $floatArg  }}
              {{$time.Set "count" $arg1}}
	            {{$time.Set "message" $arg1}}
              Il y a actuellement `{{$time.Get "count"}}` messages dans le cycle.
            {{else}}
            Nombre supérieur au nombre de message fixé.
            {{end}}
        {{else}}
        **Usage:** msg <number>
        {{end}}
    {{end}}
{{else}}
**Usage:** <cycle|msg|day|speed> <number>
{{end}}
{{dbSet 0 "time" $time}}
