All these command use dictionnary, so it will not have much space. YEAY !

# Delreroll

**Trigger** : command
**Usage** : `$delchar name/ID`

> Supprime un personnage (joueur ou reroll) de la base de donnée.

# Getchar

**Trigger** : Command
**Usage** : `$getchar (-all|stats|implant) (-ID id)`

> - Les arguments sont optionnels. 
> - Donne les statistiques demandé, sans, donne toutes les statistiques. Attention, si un ID est indiqué, `-all` est obligatoire pour obtenir le même effet. 

# Getreroll

**Trigger** : Command
**Usage** : `$getreroll name (-stats|-implant)`

> - Nom obligatoire. 
> - Donne les statistiques demandées. Sans, elle va donner 

# Setreroll

**Trigger** : Command
**Usage** :
* All : `$setreroll name -all force implant_force endurance implant_force agilité implant_agilité précision implant_précision intelligence implant_intelligence karma`
* Stats : `$setreroll name -stats force endurance agilité précision intelligence karma`
* Implant : `$setreroll name -implant force endurance agilité précision intelligence`
> - Met les statistiques d'un "reroll" dans la base de donnée. 
>
> - Le nom est obligatoire. 

# Setchar

**Trigger** : Command
**Usage** :
* All : `$setchar -all (-ID id) force implant_force endurance implant_force agilité implant_agilité précision implant_précision intelligence implant_intelligence karma `
* Stats : `$setchar -stats (-ID id) force endurance agilité précision intelligence karma`
* Implant : `$setchar -implant (-ID id) force endurance agilité précision intelligence`
> - Met un personnage joueur (or reroll) dans la base de données. 
> - L'argument `-ID` est optionnel. Sans, il va utiliser l'ID de la personne qui a lancé la commande. 
