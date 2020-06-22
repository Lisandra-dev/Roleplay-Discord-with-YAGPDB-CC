{{if .CmdArgs}}
  {{if eq (index .CmdArgs 0) "fusil"}}
    {{if (dbGet .User.ID "fusil")}}
      {{$y := (dbGet .User.ID "fusil").Value}}
      {{$x := sub 12 $y}}
      {{if lt $y (toFloat 12)}}
        {{ $embed := cembed
          "description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $x)) " charges de fusil !")}}
        {{ sendMessage nil $embed }}
      {{else}}
        {{ $embed := cembed
          "description" "Votre fusil est vide."
        }}
        {{ sendMessage nil $embed }}
      {{end}}
  {{else}}
    {{ $embed := cembed
      "description" (joinStr "" .User.Mention ", il vous reste 12 charges de fusil !")}}
    {{ sendMessage nil $embed }}
  {{end}}


  {{else if eq (index .CmdArgs 0) "pistolet"}}
    {{if (dbGet .User.ID "pistol")}}
      {{$a := (dbGet .User.ID "pistol").Value}}
      {{$b := sub 8 $a}}
      {{if lt $a (toFloat 7)}}
        {{ $embed := cembed
            "description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $b)) " charges de pistolet !")
        }}
        {{ sendMessage nil $embed }}
      {{else}}
        {{ $embed := cembed
          "description" "Votre pistolet est vide."}}
        {{ sendMessage nil $embed }}
      {{end}}
    {{else}}
      {{$embed := cembed
          "description" (joinStr "" .User.Mention ", il vous reste 8 charges de pistolet !")}}
      {{ sendMessage nil $embed }}
    {{end}}

    {{else if eq (index .CmdArgs 0) "canon"}}
      {{if (dbGet .User.ID "canon")}}
        {{$c := (dbGet .User.ID "canon").Value}}
        {{$d := sub 20 $c}}
        {{if lt $c (toFloat 20)}}
          {{ $embed := cembed
            "description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $d)) " charges de canon !")}}
          {{ sendMessage nil $embed }}
        {{else}}
          {{ $embed := cembed
            "description" (joinStr "" .User.Mention ", votre canon est vide !")}}
          {{ sendMessage nil $embed }}
        {{end}}
      {{else}}
        {{ $embed := cembed
          "description" (joinStr "" .User.Mention ", il vous reste 20 charges de canon !")}}
        {{sendMessage nil $embed }}
      {{end}}

      {{else}}
      **Usage : ** `$etat <canon|pistolet|fusil>`
      {{end}}
    {{else}}
    **Usage : ** `$etat <canon|pistolet|fusil>`
    {{end}}
    {{deleteTrigger 1}}
    {{deleteResponse 5}}
