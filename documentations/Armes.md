Toutes les commandes listées peuvent être utiliser avec un reroll, en indiquant `#nom` à la fin de la commande.



# PA

Le centre du système d'arme.

**Trigger : ** Coutains `:pa:`

> Compte les PA d'un utilisateur à chaque trigger. A quatre, l'utilisateur a utilisé tous ses PA. A chaque fois que la commande est lancé, elle recherche dans la base de données si l'utilisateur recharge quelque chose (arme ou compétence). Si oui, elle va faire `+1`. Si cette valeur atteint le seuil, alors elle va indiquer le rechargement complet puis supprimer la clé de la base de données. En outre, pour les armes, elle va reset le compte et indiquer donc le maximum de charge.

# Recharge

**Trigger : ** Command

**Usage : ** 

* <u> Arme :</u>`$recharge -(fusil|fusil2|pistolet|pistolet2|canon)`
* <u> Compétence :</u> `$cd -(court|long) "name"`

Il faut :

* Pistolet : 2 PA
* Fusil : 4 PA
* Canon : 8 PA
* Compétence d'attaque : 8 PA
* Compétence de support : 12 PA (les altération durent 6 PA, et le cooldown 6PA.)


# Fusil, fusil2, pistolet, pistolet2, canon

**Trigger : ** Contains :

* Pistolet : `:pistol:` & `:pistol2:`

* Fusil : `:fusil` & `:fusil2`

* Canon : `:canon:`

  > Le bot compte les balles utilisés par l'utilisateur quand il l'indique avec l'émoji associé.

  # Bullet stock

**Trigger :** Command

**Usage : ** `$stock -(pistolet|pistol2|canon|fusil|fusil2)`

> Indique le nombre de charge restante dans l'arme. Si l'utilisateur n'a pas utilisé son arme, indique le nombre de charge maximal.

# Turn

**Trigger : **Command

**Usage : ** `$turn (random argument optional)`

> Reset pour tout les membres du groupe les PA, et rajoute +1 au nombre de turn, sauf si un argument est passé en paramètre. Dans ce cas, le tour retourne à 1. 

# Delweapon

**Trigger : ** Command

**Usage : ** `$delweapon <mention|#name>`



> Permet de supprimer dans la base de données toutes les données des armes relative à la personne passée en paramètre.

