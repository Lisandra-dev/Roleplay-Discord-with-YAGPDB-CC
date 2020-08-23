Toutes les commandes listées peuvent être utiliser avec un reroll, en indiquant `#nom` à la fin de la commande. En outre, les emoji peuvent être indiqué à la fin de la commande.


# PA et PA2

Le centre du système d'arme.

**Trigger :** Coutains `:pa:`

**Trigger : ** Countains `:2pa:`

> Compte les PA d'un utilisateur à chaque trigger. A quatre, l'utilisateur a utilisé tous ses PA. A chaque fois que la commande est lancé, elle recherche dans la base de données si l'utilisateur recharge quelque chose (arme ou compétence). Si oui, elle va faire `+1`. Si cette valeur atteint le seuil, alors elle va indiquer le rechargement complet puis supprimer la clé de la base de données. En outre, pour les armes, elle va reset le compte et indiquer donc le maximum de balle.

# recharge automatique des armes

Fun : c'est la seule commande qui n'utilise pas de dictionnaire avec le temps, son utilisation nécessite donc un premium car il y a de nombreux `dbGet`. Cette commande permet d'activer la recharge des armes et des compétences.

**Trigger :** Commande

**Usage :**

* <u> Arme :</u>`$recharge (fusil|fusil2|pistolet|pistolet2|canon)`
* <u> Compétence :</u> `$competence (long|court) "name"`

Il faut :

* Pistolet : 2 PA
* Fusil : 4 PA
* Canon : 8 PA
* Compétence d'attaque : 4 PA
* Compétence de support : 8 PA (les altération durent 4 PA, et le cooldown 4PA.)

# recharge manuelle des armes :

* **Il faut absolument que l'utilisateur ait des balleurs dans son inventaire et que les noms soient indiquent à la fois dans l'inventaire, la boutique, et la fonction.**

 > Pour des raisons de praticités, les balleurs portent le nom de l'arme. Si dans la boutique, des armes (différente d'un balleur) doivent y être mis, je conseille de rajouter un indicateur pour les ARMES (**et non les balleurs**), tel que [A] ou [E] par exemple.
 > Le nom peut être cependant changé, dans ce cas, il faut changer la ligne indiqué en indiquant le nom exacte des balleurs se trouvant dans la boutique.

* Les fichiers `pa` et `2pa` ne contiennent la recharge que pour les compétences.
* Le fichier `competence` ne change pas entre les deux versions.
* Note : recharger coûte 2PA, la fonction le vérifie donc.
* Le compte des `2pa` peut être indiqué directement dans la commande.

**Trigger :** Commande

**Usage :** `$recharge (fusil|fusil2|pistolet|pistolet2|canon)`


# Fusil, fusil2, pistolet, pistolet2, canon

**Trigger :** Contains :

* Pistolet : `:pistol:` & `:pistol2:`

* Fusil : `:fusil` & `:fusil2`

* Canon : `:canon:`

  > Le bot compte les balles utilisés par l'utilisateur quand il l'indique avec l'émoji associé.
  > Les emoji peuvent être indiqué directement dans la commande de dés.

  # Bullet stock

**Trigger :** Command

**Usage :** `$stock -(pistolet|pistol2|canon|fusil|fusil2)`

> Indique le nombre de balle restante dans l'arme. Si l'utilisateur n'a pas utilisé son arme, indique le nombre de balle maximal.

# Turn

**Trigger :**Command

**Usage :** `$turn (random argument optional)`

> Reset pour tout les membres du groupe les PA, et rajoute +1 au nombre de turn, sauf si un argument est passé en paramètre. Dans ce cas, le tour retourne à 1.

# Delweapon

**Trigger :** Command

**Usage :** `$delweapon <mention|#name>`


> Permet de supprimer dans la base de données toutes les données des armes relative à la personne passée en paramètre.
