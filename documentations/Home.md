Toutes les commandes sont en français, et ont été crée pour un RP de science-fiction avec un gameplay particulier (vous pouvez retrouver plus d'info à ce sujet [ici](https://www.jdr-system.ovh/readme.html). Cependant, vous pouvez les adapter et certaines commandes (notamment celle du groupe "ungrouped") peuvent être utilisée sans édit.

A noter que j'utilise ici `$` pour le "trigger" des commandes du bot. Vous devez donc penser à la changer sur les commandes pour l'adapter à votre serveur.



A noter que la majorité des ID ici sont celles du serveurs, donc penser bien à les changer. 



Nous avons donc :

* Le groupe "misc", qui contient des commandes "fun" et d'aide, avec des commandes d'administration sans aucun trigger (notamment des commandes qui suppriment automatiquement, ou basé sur les réactions à des messages). 

* Le dossier pour les notes, qui a été pris entièrement depuis le serveurs officiels de YAGPD. Il est composé de deux parties : 

  * Un système de notes sur le serveurs et où tous les utilisateurs peuvent y accéder. 
  * Un système de note "privé", où les notes sont envoyés en DM.

* Le système d'arme, qui doit être pris entièrement, se basant sur un système de PA. Les armes et compétences se rechargent donc en fonction du nombre de PA que l'utilisateur utilise. De plus, le PA et les armes ont une fonction qui permet de compter les utilisations, et la fonction `:pa:` est très centrale ici. Dans ce dossier se trouve aussi les fonctions reliés au bouclier, qui prend en compte le nombre de message pour recharger ce dernier.

  > Il se trouve aussi l'ancienne version du gameplay, où tout était basé sur la longueur des messages et le nombre de message pour recharger les armes et compétences.

* Le système de temps permet, au final, d'avoir une sorte de temps interne au serveur, basé sur le nombre de message envoyé dans les catégories de RP. Ce temps peut être facilement modifié, et à chaque nouveau cycle, un nom de channel est modifié, et un embed envoyé. Actuellement, il y a quatre cycles. A la fin du quatrième, une nouvelle journée commence. Il est suffisamment facile de modifier le nombre de cycle pour une journée. Une commande spécifique a été créée pour modifier facilement les paramètres de cette commande. Une commande permet d'accéder au temps actuel, avec le nombre de message présent dans le cycle. 

* Le dossier dés contient tout le système de dés du RP, où le bot tire un dé et affiche un message avec le résultat. Un système de bonus/malus permet de modifier se résultat, et les messages sont basés sur ce résultat "final". Des commandes pour certaines attaques spéciales ont été faite et indique directement le résultat de l'action (nombre de PV soigné, nombre de pv retiré par un poison, malus, réussite d'esquive...). Les dés sont tirés en fonctions des statistiques indiquées par le joueur et le reliant à la base de donnée. De plus, on peut indiquer un nom de personnage pour tirer un dé sur ce "reroll".

* Le dossier "database" contient les fonctions qui permettent d'intégrer dans la base de données les statistiques d'un reroll ou d'un personnage. 

* Le système d'économie permet d'avoir une boutique et un système monétaire, et les joueurs  (ou reroll) peuvent vendre ou acheter des objets. La boutique peut être ouverte ou fermée. De plus, une commande permet de mettre à jour l'inventaire du vaisseau en fonction de ce qui a été créée.



