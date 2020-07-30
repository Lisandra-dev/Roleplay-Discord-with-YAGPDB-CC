{{$chan := reFind `\-chan` .Message.Content}}
{{$id := 0}}
{{$channel := "nil"}}
{{if $chan}}
	{{$channel = (toInt (index .CmdArgs 1))}}
	{{$id = (toInt (index .CmdArgs 2)) }}
{{else}}
	{{$channel = .Channel.ID}}
	{{$id = (toInt (index .CmdArgs 0))}}
{{end}}
{{$msg:= getMessage $channel $id}}
{{if $msg.Embeds}}
```
{{(index $msg.Embeds 0).Description}}
```
{{else}}
```
{{$msg.Content}}
```
{{end}}
{{deleteTrigger 1}}
{{deleteResponse 180}}
