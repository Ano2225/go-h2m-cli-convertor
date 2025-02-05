# 🚀 H2M - Markdown ↔ HTML Converter  

## 📌 Description  
H2M est un outil en ligne de commande (CLI) permettant de convertir des fichiers :  
✅ **Markdown (.md) → HTML (.html)**  
✅ **HTML (.html) → Markdown (.md)**  

Il est conçu pour être rapide, simple et efficace.  

---

## 📥 Installation  
### 🔧 Prérequis  
- Go installé (`>= 1.18`)  

### 📌 Étapes  
Clone le dépôt et build le projet :  
```sh
git clone https://github.com/Ano2225/go-h2m-cli-convertor
cd go-h2m-cli-convertor
go build -o h2m
```
---

## 🚀 Utilisation  

### 🔹 Conversion Markdown → HTML  
```sh
./h2m md2html -i document.md -o resultat.html
```

### 🔹 Conversion HTML → Markdown  
```sh
./h2m html2md -i page.html -o resultat.md
```

### 🔹 Mode interactif  
Lance l'outil sans argument pour entrer dans le mode interactif :  
```sh
./h2m
```
Tape `exit` ou utilise `Ctrl+C` pour quitter.  

---

## ⚡ Exemples  
#### 📄 Fichier `sample.md` :  
```md
# Hello World
This is a **Markdown** file.
```
#### 🔄 Conversion :  
```sh
./h2m -i sample.md -o sample.html
```
#### 📄 Résultat `sample.html` :  
```html
<h1>Hello World</h1>
<p>This is a <strong>Markdown</strong> file.</p>
```
---

## 🤝 Contribution
Forke, modifie, et envoie une Pull Request !


## 🛠️ Mainteneur  
👤 **Arouna Ouattara**  
🔗 [LinkedIn](https://www.linkedin.com/in/arouna-ouattara/)  

🚀 **Happy coding!** 😃