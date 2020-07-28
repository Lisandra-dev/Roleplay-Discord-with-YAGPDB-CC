All these command use dictionnary, so it will not have much space. YEAY !

# Delreroll :

**Trigger** : command
**Usage** : `-delchar name/ID`

> Delete in the database, with using the name for a reroll, and an ID for another player.

# Getchar :

**Trigger** : Command
**Usage** : `-getchar (-all|stats|implant) (-ID id)`
> Arguments are optional
> Give the stats in the DB

# Getreroll :

**Trigger** : Command
**Usage** : `-getreroll name (-stats|-implant)`
> Name is required
> Other arguments aren't. Without, it will give all the stats (implant + stats)

# Setreroll :

**Trigger** : Command
**Usage** :
* All : `-setreroll name -all force implant_force endurance implant_force agilité implant_agilité précision implant_précision intelligence implant_intelligence karma`
* Stats : `-setreroll name -stats force endurance agilité précision intelligence karma`
* Implant : `-setreroll name -implant force endurance agilité précision intelligence`
> Set in the DB a "reroll", using his name like an ID (convert in int)

#Setchar :

**Trigger** : Command
**Usage** :
* All : `-setchar -all (-ID id) force implant_force endurance implant_force agilité implant_agilité précision implant_précision intelligence implant_intelligence karma `
* Stats : `-setchar -stats (-ID id) force endurance agilité précision intelligence karma`
* Implant : `-setchar -implant (-ID id) force endurance agilité précision intelligence`
> Set in the DB someone
> `-ID` is optional : without, it's for the user who trigger the command
