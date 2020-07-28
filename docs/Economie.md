Merci à Shadownight#9025 de me permettre d'utiliser ses commandes. 

Toutes les commandes pour acheter, vendre, ouvrir l'inventaire et le porte-monnaie peuvent être utiliser avec `#name` à **la fin de la commande**
Pour les admins, l'user peut être un reroll, en ajoutant `#`à la place de `user`.

Toutes ses commandes utilisent des dictionnaires ou des listes, sauf pour l'ouverture / fermeture du shop, qui est une simple entrée de BD. 

# ADMIN

**Trigger : ** Command

**Usage :**

* Money : `-admin -money -(add|remove|reset|set) user amount`
 > L'utilisateur sera averti si un admin change son porte-monnaie. 

* Symbol : `-admin -symbol <symbol>`

* Manage inventory : `-admin -manageinv -(add|remove|set) user item amount`

* Reset/clean inventory : `-admin -manageinv -(reset|clean) user`

  > Reset détruit complètement l'inventaire du joueur, quand clean le vide seulement.



# Item

**Trigger : ** Command

**Usage :**

* Add : `$item -add "name" price (none|sellprice) (inf|stockage)(true|false)"description"`
* Delete : `$item -delete "name"`
* Information : `$item -info "name"`

> - Attention à la casse. 
>
> - SII veut dire si l'objet peut apparaître ou non dans l'inventaire du joueur. 

* `$item -edit "name" -(price|sell|stock|sii|desc|rbuy|rsell) "value"`

Type que doit être `value` : 
* <u>Prix : </u> :  nombre. 
* <u>Vendre </u> : `0` ou un nombre.
* <u>Stock</u> :  `inf` ou un nombre
* <u>SII</u> : true ou false
* <u>desc</u> : Chaîne de caractère.
* <u> rbuy & rsell </u> : Ce sont les deux seules commandes qui ont besoin de deux valeurs : une borne "start" et "end" qui permet de choisir aléatoirement un nombre entre ses deux valeurs. 

# Store

**Trigger :** Command

**Usage :** `$store (optional page)`

> Montre la boutique. 
> Que la boutique soit fermée ou ouverte, un message sera affiché. 
> S'il y a beaucoup d'objet, l'utilisateur peut indiquer une page. 

# Open & close

**Trigger : ** Regex `\$(close|open)`
**Usage : ** `$close` ou `$open`

> - Ouvre ou ferme la boutique.
> - Si la boutique est fermée, les utilisateurs ne peuvent ni vendre ou acheter. 
> - La boutique aura un message spécial si la boutique est fermée.
> - Un message est envoyé dans un channel pour indiqué l'ouverture/fermeture de la boutique. 
>
> 

# Inventory

**Trigger :** Command

**Usage :** `$inv (optional page)`



# Sell

**Trigger : ** Command

**Usage :** `$sell "name" amount`



# Buy

**Trigger :** Command

**Usage :** `$buy "name" amount`



# Money

**Trigger :** Command

**Usage : ** `$money`

> Affiche le porte monnaie de l'utilisateur. 

# VN

**Trigger :** Command

**Usage : ** `$vn -(add|use) (-armes?|modules?|BC|LC|SF|CU|bc|lc|cb|sf|cu) "value"`

**Indication sur la valeur : ** 
* <u>Armes & module</u>: nom
* <u> Implant :</u> Quantité 

> La fonction va mettre à jour le message de l'inventaire à chaque fois que quelqu'un utilise la commande pour éditer cet inventaire. 
> Le message de l'inventaire doit **obligatoirement** être fait par YAGPD. 

# Craft

**Trigger : ** Command

**Usage : ** `$(craft|use) (#nom) (quantité) "nom"`

Pour Use, quantité peut être `all` ce qui permet d'utiliser tous les objets directement. 



Pour les rerolls, il faut ajouter `#nom` avant d'indiquer la quantité et le nom. 

> Permet à l'utilisateur de se rajouter à lui ou à un reroll un objet qui ne se trouve pas forcément dans le store. Si l'objet se trouve déjà dans son inventaire, alors il augmente la quantité. 
>
> Cette dernière est d'ailleurs optionnelle. Non précisé, elle sera de 1. La commande `$use` permet de retirer l'objet de l'inventaire, en précisant un montant ou non. 
>
> Le bot est sensible aux fautes d'orthographes et de frappes. Il est donc possible que vous ayez un affichage doublé pour le même objet si vous avez fait une erreur d''orthographe ou de frappe. 
>
> *Il est impossible de vérifier toutes les orthographes possibles pour chaque objet, faites donc attention.*
>
> Le bot n'est pas sensible à la casse : une fonction permet de mettre automatiquement la première lettre en majuscule, c'est ce qui est fait ici. 