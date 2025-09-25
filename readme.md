# Mini CRM Verdiane — Go

Gestionnaire de contacts clients développé en **Go**.

## Fonctionnalités

Ce CRM propose un menu interactif avec les options suivantes : 

- Ajout d'un nouveau client
- Affichage de tous les clients
- Suppression d'un client par son ID
- Mise à jour d'un client
- Ajout d'un client via des flags CLI (`--name`, `--email`)

## Installation

1. **Cloner le projet :**

```bash
git clone https://github.com/Verdiane123/minicrm_.git
cd minicrm

2. **Exécuter le programme :**
go run main.go

3. **Ajout d'un client via les flags CLI :**
go run main.go --name="Verdiane" --email="verdiane@verdiane.fr"

4. Évolutions : branche persJson

Cette version du CRM est modulaire et utilise la persistance de données via un fichier JSON.
Pour tester cette branche, faire :
- git checkout persJson
- go run cmd/crm/main.go
** Résulatat : Les nouveaux clients seront ajoutés dans un fichier JSON.

5. Évolutions : branche bddGorm
Cette version intègre l'ajout d'une base de données avec Gorm
Pour tester faire
- git checkout bddGorm
- go run ./cmd/crm add --name "Le_nom-du_client" --email "son@dresse.com"
Résultat : Client ajouté: &{ID Le_nom-du_client son@dresse.com}