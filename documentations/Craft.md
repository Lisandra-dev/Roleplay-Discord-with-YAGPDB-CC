Les commandes de craft sont en lien avec celles de l'économie, mais sont profondément lié au fonctionnement du RP.

# Recipe :

## Ajout :

**Trigger** : Regex `\$recipe \-add`

**Usage** : `$recipe "nom" -(bc|lc|sf|cb) (-(bc|lc|sf|cb)) (-(bc|lc|sf|cb)) (-(bc|lc|sf|cb)) (-bdg)`

> Permet de rajouter une recette
> Tous les arguments sont optionnels, sauf les deux premiers.
> BDG permet d'indiquer le "bas de gamme"

## Info :

**Trigger** : Regex `\$recipe \-info`

**Usage** : `$recipe -info "nom"`

> Permet d'afficher la recette d'un objet.

## Edit :

**Trigger** : Regex `\$recipe \-edit`

**Usage** : Usage : `$recipe -edit "Nom" -(bc|lc|sf|cb)`

> Permet d'éditer rapidement une recette

# Craft :

**Trigger** : Commande

**Usage** : `$craft "objet" (q[1-100]) (+VN|-case) (-bdg) (-cu) (#reroll)`

L'objet est automatiquement mis dans le casier du crafteur.
> - Permet de créer un objet.
> - Les arguments entre parenthèses sont optionnels.
> - `q` est un regex, pour indiquer une quantité elle doit être écriture comme suit : `q20` ou encore `q10`.
> - `+VN `permet d'utiliser à la fois votre inventaire et celui du vaisseau, d'où le "+"
> `-case` permet d'utiliser uniquement l'inventaire du Nucleus.
> `-bdg` permet d'indiquer que l'objet est un objet bas de gamme, pour rajouter une écriture précise sur ces objets.
> De même, si le nom de l'objet contient "chargeur", il sera réécrit de manière précise.
> `-cu` permet d'utiliser un composant universel pour crafter automatiquement l'objet.

*Note : La commande est séparée en deux à cause de la longueur. La commande contenant uniquement VN aura pour trigger un regex : `^\$craft`*

# Craft - BDG : 

**Trigger** : Commande

**Usage** : `$bdg "objet" (q[1-100]) (-cu) (#reroll)`
> Si l'utilisateur de la commande a l'autorisation, il peut fabriquer un objet à partir des composants de son inventaire.
> L'objet se met donc dans l'inventaire de l'utilisateur.
> Fabrique **uniquement** des objets "bas de gamme".

# Recycle :

**Trigger** : Commande

**Usage** : `$recycle "objet" (q[1-100]) (+VN) (#Reroll)`

> - Permet de recycler un objet : on récupère tous les objets qui ont permis à le créer.
> - `+VN` met les composants dans l'inventaire du Nucleus.
> - Par défaut, la quantité est de 1.
