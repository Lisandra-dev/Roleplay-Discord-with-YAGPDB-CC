# URC/RC :

**Trigger** : Regex `^\-(rc|urc)`

**Usage** : `-rc|rcc <value>`

> Calculate the RC / URC of a value.

# dice info :

**Trigger** : Regex `^\-(malus|poison|soin)`

**Usage** : `-malus|poison|soin <value>`

> Give the information about the result for some dice.

# d | Esquive | poison | soin | malus

**Trigger** : command

**Usage** :
* d : `-d (-bonus|malus) (-bonus|malus) (#charname) (commentaire)`
* esquive : `-esquive (-bonus|malus) (-bonus|malus) (#charname) (commentaire)`
* poison : `-poison (-bonus|malus) (-bonus|malus) (#charname) (commentaire)`
* soin : `-soin (-bonus|malus) (-bonus|malus) (#charname) (commentaire)`
* malus : `-malus (-bonus|malus) (-bonus|malus) (#charname) (commentaire)`

> Same as above, but the message is a simple markdown.
> Argument are optional.
> If #charname, take from the "reroll database". If it not exists, use the default value. Without, take the stats value with the User.ID

# choose

**Trigger** : command

**Usage** : `-choose arg1 arg2...`

> Choose one random argument in a list.

# Carac

**Trigger** : Command

**Usage** : `-carac`

> Same as above, but the list is "Force", "Endurance", "Agilité", "Précision", "Intelligence", "Karma"
