# URC/RC

**Trigger** : Regex `^\$(rc|urc)`

**Usage** : `$rc|rcc <value>`

> Calcule la valeur d'un URC/RC.

# dice info

**Trigger** : Regex `^\$(malus|poison|soin)`

**Usage** : `$-malus|poison|soin <value>`

> Donne le résultat d'un poison/soin/malus pour le résulat donné. Pratique si on a lancé le mauvais dés !

# d | poison | soin | malus

**Trigger** : command

**Usage** :
* d : `$d (-bonus|malus) (-bonus|malus) (#name) (commentaire)`
* poison : `$poison (-bonus|malus) (-bonus|malus) (#name) (commentaire)`
* soin : `$soin (-bonus|malus) (-bonus|malus) (#name) (commentaire)`
* malus : `$malus (-bonus|malus) (-bonus|malus) (#name) (commentaire)`
* Soutien : `$soutien (-bonus|malus) (-bonus|malus) (#name) (commentaire)`


> * Lance un dé correspondant à la commande, et affiche le résultat. Les commandes spéciales "malus", "poison", "soin" et "esquive" affiche donc les résultats d'une action spécifique.
> * Les arguments sont optionnels, cependant, les bonus sont obligatoirement en négatif, sinon ils ne sont pas reconnus.
> * Si la personne veut prendre la bonne statistiques stockée dans la base de donnée, elle doit l'indiquer, sinon, le dé prend les valeurs par défauts (seuil de 8, bonus de 0.)
> * `#name` permet d'indiquer un autre personnage que celui du joueur, notamment un reroll. Le bot va donc utiliser ce nom pour fouiller la base de donnée si ce dernier est enregistré. Attention à utiliser le même nom ! Elle n'est cependant pas sensible à la casse.

# Choose

**Trigger** : command

**Usage** : `$choose arg1 arg2...`

> Choisit un argument dans une liste, aléatoirement.

# Carac

**Trigger** : Command

**Usage** : `$carac`

> Comme la commande précédente, sauf que la liste est intégré au bot, et permet de choisir aléatoirement une caractéristique (hormis le karma, puisqu'il est impossible d'altérer la chance de quelqu'un).
