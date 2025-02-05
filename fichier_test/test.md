# go-markdown-to-html-cli

## 📌 Description
Un outil CLI permettant de convertir du **Markdown** en **HTML** et inversement.

## 🚀 Installation

### 1️⃣ Cloner le projet
```sh
git clone https://github.com/tonuser/go-markdown-to-html-cli.git
cd go-markdown-to-html-cli
```

### 2️⃣ Installer les dépendances
```sh
go mod tidy
```

### 3️⃣ Compiler le programme
```sh
go build -o h2m
```
> Sur Windows, l'exécutable sera `h2m.exe`

## 🛠️ Utilisation

### 🔹 Afficher l'aide générale
```sh
./h2m --help   # Linux/macOS
h2m.exe --help # Windows
```

### 🔹 Convertir un fichier HTML en Markdown
```sh
./h2m html2md --input fichier.html --output fichier.md
```
> Remplace `fichier.html` par ton fichier source et `fichier.md` par le fichier de sortie.

### 🔹 Convertir un fichier Markdown en HTML
```sh
./h2m md2html --input fichier.md --output fichier.html
```

### 🔹 Tester le programme sans fichier de sortie
Si tu ne spécifies pas `--output`, le résultat s'affichera dans la console :
```sh
./h2m html2md --input fichier.html
```

## 🔍 Exemple de Test
Crée un fichier `test.md` avec ce contenu :
```markdown
# Titre Principal

- Point 1
- Point 2
- Point 3
```
Puis exécute :
```sh
./h2m md2html --input test.md --output test.html
```
Cela générera un fichier `test.html` contenant le HTML correspondant.

## 📜 Licence
Ce projet est sous licence MIT.

