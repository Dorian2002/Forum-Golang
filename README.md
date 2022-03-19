# Forum

## Notre Groupe : 
* Fabio Garcia
* Fonseca Dorian
* Adonaï Killian
* D'Autheville Adam

## Notre Projet : 
* Le Forum est un projet de site web en golang, html, css et SQLite.
* Son but est de créer un forum d'utilisateurs avec des intéractions via des postes.
* Début du projet : vendredi 12 février.
* Fin du projet : dimanche 02 mai.

## Procédure d'installation du projet  en local : 
* Télécharger la totalité du REPOSITORY avec les fichiers du projet.
* Ouvrir Visual Code Studio sur votre ordinateur (installez le si vous ne l'avez pas).
* Cliquez sur ```File``` en haut à gauche du logiciel.
* Sélectionner ```Add Folder to Workspace```.
* Selectionner le REPOSITORY préalablement téléchargé.
* Ouvrir un terminal de commande en haut à gauche du logiciel.
* Assurez vous d'être dans le chemin d'accès ```forum\src``` car c'est là que ce trouve le code source.
* Tapez la commande ```go run main.go``` dans le terminal pour exécuter le programme.
* Ouvrez une page web et dans la barre de recherche tapez l'adresse suivante : ```localhost:1111```

Le projet devrait être installé et lancé.

## Exemple d'utilisation : 
* Un fois sur la page d'accueil du site web. Vous trouverez un menu sur la gauche de l'écran.
* En haut à droite vous pourrez vous Login ou vous Register, c'est ensuite dans le menu de gauche que vous pourrez vous Logout grâce au bouton du même nom.
* Dans le menu de gauche vous touverez aussi un bouton ```profil```, il vous permettra de voir les données de votre profil. Si vous n'êtes pas connecté, ce bouton vous redirigera vers la page Login.
* Vous trouverez juste en dessous du bouton profil, un bouton ```new post``` grace auquel vous accederez à la page de création de postes.
* Les trois derniers boutons de ce menu ne sont pas fonctionnels.
* L'affichage des postes n'est pas fonctionnel, mais ils sont bien enregistrés dans la base de donnée.
* Enfin les catégories des différents posts ne sont pas fonctionnelles  non plus.