# URC/RC :

**Trigger** : Regex `^\-(rc|urc)`

**Usage** : `-rc|rcc <value>`

> Calculate the RC / URC of a value.

# dice info :

**Trigger** : Regex `^\-(malus|poison|soin)`

**Usage** : `-malus|poison|soin <value>`

> Give the information about the result for some dice.

# d :

**Trigger** : command

**Usage** : `-d (implant|malus) (implant|malus) (commentaire)`

> Same as above, but the message is a simple markdown.
> Argument are optional.

# choose

**Trigger** : command

**Usage** : `-choose arg1 arg2...`

> Choose one random argument in a list.

# Carac

**Trigger** : Command

**Usage** : `-carac`

> Same as above, but the list is "Force", "Endurance", "Agilité", "Précision", "Intelligence", "Karma"

# Malus | Poison | Soin

**Trigger** : Command

**Usage** `-malus (implant|malus) (implant|malus) (commentaire)`

> Arg are optional.
> Same as "dice info" but in simpler markdown and roll the dice directly.

# Set :

**Trigger** : command

**Usage** : `-set (id) Force Endurance Agilité Précision Karma`

> Set in the YAGPDB database the char value for a User ID (can be set by another people by it's ID or just using USER.ID if there are no ID set)

# Get :

**Trigger** : command

**Usage** : `-get characteristics`

> Get the char value (debugging only)
