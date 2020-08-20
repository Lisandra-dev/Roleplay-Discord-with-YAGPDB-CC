{{$chan := reFind `\-chan` .Message.Content}}
{{$id := 0}}
{{$channel := "nil"}}
{{$champ := reFind `\-(footer|title|field(s?)|color|author)` .Message.Content }}
{{if $chan}}
	{{$channel = (toInt (index .CmdArgs 1))}}
	{{$id = (toInt (index .CmdArgs 2)) }}
{{else}}
	{{$channel = .Channel.ID}}
	{{$id = (toInt (index .CmdArgs 0))}}
{{end}}
{{$msg:= getMessage $channel $id}}
{{if $msg.Embeds}}
  {{with (index $msg.Embeds 0)}}
	{{if not $champ}}
```
{{.Description}}
```
	{{else if eq $champ "-field" "fields"}}
		{{ range $i, $field := .Fields }} 
			{{ add $i 1 }}
```
{{ $field.Name }} : {{ $field.Value }}
```
		{{ end }}
	{{else if eq $champ "-footer"}}
```
{{.Footer.Text}}
```
	{{else if eq $champ "-title"}}
```
{{.Title}}
```
	{{else if eq $champ "-color"}}
```
{{.Color}}
```
{{else if eq $champ "-author"}}
```
{{.Author.Name}}
```
	{{end}}
  {{end}}
{{end}}
{{if $msg.Content}}
```
{{$msg.Content}}
```
{{end}}
{{deleteTrigger 1}}
{{deleteResponse 180}}