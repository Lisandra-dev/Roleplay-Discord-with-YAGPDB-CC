{{$c := "count"}}
{{$t := "time"}}
{{$message := "message"}}
{{$amount := (toFloat (dbGet 0 "mgsc").Value)}}

{{if .CmdArgs}}

    {{if eq (index .CmdArgs 0) "cycle"}}
        {{if eq (len .CmdArgs) 2}}
            {{$floatArg := (toFloat (index .CmdArgs 1)) }}
            {{$arg1 := (toFloat (index .CmdArgs 1)) }}
            {{if ge (toFloat 4) $floatArg  }}
                {{dbSet 0 $t $arg1}}
                Nous sommes actuellement au cycle  `{{(dbGet 0 $t).Value}}`.
            {{else}}
              Le nombre est supérieur à 4.
            {{end}}
        {{else}}
            **Usage:** cycle <number>
        {{end}}

    {{else if eq (index .CmdArgs 0) "speed"}}
      {{if eq (len .CmdArgs) 2}}
          {{dbSet 0 "mgsc" (index .CmdArgs 1)}}
          Le nombre de message pour passer un cycle est maintenant fixé à : `{{(index .CmdArgs 1)}}`.
      {{else}}
          **Usage** : timer <number>
      {{end}}

    {{else if eq (index .CmdArgs 0) "day"}}
      {{if eq (len .CmdArgs) 2}}
        {{$floatArg := (toFloat (index .CmdArgs 1)) }}
        {{dbSet 0 "jour" (toFloat (index .CmdArgs 1)) }}
        {{$jour := (toString (toInt (dbGet 0 "jour").Value))}}
        Nous sommes maintenant au jour : `{{$jour}}`.
      {{else}}
        **Usage** day <number>
      {{end}}

    {{else if eq (index .CmdArgs 0) "msg"}}
        {{if eq (len .CmdArgs) 2}}
            {{$floatArg := (toFloat (index .CmdArgs 1)) }}
            {{$arg1 := (toFloat (index .CmdArgs 1)) }}
            {{if ge $amount $floatArg  }}
              {{dbSet 0 $c $arg1}}
	            {{dbSet 0 $message $arg1}}
              Il y a actuellement `{{(dbGet 0 $c).Value}}` messages dans le cycle.
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
