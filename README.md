# 🦆 Quotesly

> 🎓 **Projet pédagogique** - Une API REST simple pour découvrir le développement backend avec Go, Chi Router et GORM.

Une API de citations aléatoires qui permet de gérer et récupérer des citations inspirantes. Parfait pour apprendre les concepts CRUD et les bonnes pratiques REST API !

## 🚀 Fonctionnalités

- 📖 Récupérer toutes les citations
- 🎲 Obtenir une citation aléatoire
- 🏷️ Filtrer par catégorie
- ➕ Ajouter de nouvelles citations
- 💾 Base de données SQLite embarquée

## 🛠️ Technologies utilisées

- **Go** - Langage de programmation
- **Chi Router** - Routeur HTTP léger et rapide
- **GORM** - ORM pour Go
- **SQLite** - Base de données embarquée
- **Air** - Hot reload pour le développement

## 📋 Prérequis

- Go 1.22+ installé
- Git

## ⚡ Installation

1. **Cloner le repository**
   ```bash
   git clone https://github.com/Dilgo-dev/quotesly.git
   cd quotesly
   ```

2. **Installer les dépendances**
   ```bash
   go mod download
   ```

3. **Configurer l'environnement**
   ```bash
   cp .env.sample .env
   ```

4. **Lancer l'application**
   ```bash
   go run cmd/server/main.go
   ```

L'API sera disponible sur `http://localhost:3000` 🎉

## 🔥 Développement avec Hot Reload

Pour un développement plus confortable avec rechargement automatique :

```bash
# Installer Air
go install github.com/cosmtrek/air@latest

# Lancer avec hot reload
air
```

## 📡 Endpoints de l'API

| Méthode | Endpoint | Description |
|---------|----------|-------------|
| `GET` | `/api/quotes` | 📋 Récupère toutes les citations |
| `GET` | `/api/quotes/random` | 🎲 Récupère une citation aléatoire |
| `GET` | `/api/quotes/category/{category}` | 🏷️ Récupère les citations par catégorie |
| `POST` | `/api/quotes` | ➕ Ajoute une nouvelle citation |

## 💡 Exemples d'utilisation

### Récupérer une citation aléatoire
```bash
curl http://localhost:3000/api/quotes/random
```

```json
{
  "ID": 1,
  "CreatedAt": "2024-01-15T10:30:00Z",
  "UpdatedAt": "2024-01-15T10:30:00Z",
  "DeletedAt": null,
  "Quote": "The only way to do great work is to love what you do.",
  "Author": "Steve Jobs",
  "Category": "motivation"
}
```

### Ajouter une nouvelle citation
```bash
curl -X POST http://localhost:3000/api/quotes \
  -H "Content-Type: application/json" \
  -d '{
    "Quote": "La vie est un mystère qu'\''il faut vivre, et non un problème à résoudre.",
    "Author": "Gandhi",
    "Category": "philosophie"
  }'
```

### Récupérer par catégorie
```bash
curl http://localhost:3000/api/quotes/category/motivation
```

## 🐳 Docker

### Build et run avec Docker
```bash
# Build l'image
docker build -t quotesly .

# Lancer le container
docker run -p 3000:3000 quotesly
```

---

**Happy coding!** 💻✨