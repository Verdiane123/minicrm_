    *******************CRM VERDIANE*******************************
Les fonctionnalités et les readme varient d'une branche à une autre.

Faire des checkout et voir leur readme pour tester le fonctionnement du code

branche main : script basique de création des clients via un menu interactif et des flags

branche persJson : CRM modulaire avec stockage des clients dans un fichier Json

branche bddGorm : ajout d'une BDD avec Gorm

branche cliCobra : création d'une cli professionnelle avec Cobra et Viper


# Mini CRM Verdiane — Go

Gestionnaire de contacts clients avec GO

## Fonctionnalités
Menu interactif avec les options suivantes : 
- Ajout d'un nouveau client 
- Affichage de tous les clients
- Suppression d'un client avec son ID
- Mise à jour d'un client
Ajouter d'un client via des flags (`-name`, `-email`)

## Installation
```bash
git clone https://github.com/Verdiane123/minicrm_.git
git checkout main
Exécuter en lançant la commande "go run main.go"
Utiliser les flags en faisant ''go run main.go -name="Verdiane" -email="verdiane@verdiane.fr"''
Résultat: le client est créé soit via un menu interactif soit en ligne de commande












