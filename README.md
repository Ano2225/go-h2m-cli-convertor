# 🚀 H2M - Markdown ↔ HTML Converter  

## 📌 Description  
H2M is a command-line tool (CLI) for converting files:  
✅ **Markdown (.md) → HTML (.html)**  
✅ **HTML (.html) → Markdown (.md)**  

Fast, simple, and efficient.  

---

## 📥 Installation  
### 🔧 Requirements  
- Go installed (`>= 1.18`)  

### 📌 Steps  
Clone the repository and build the project:  
```sh
git clone https://github.com/Ano2225/go-h2m-cli-convertor
cd go-h2m-cli-convertor
go build -o h2m
```
---

## 🚀 Usage  

### 🔹 Convert Markdown → HTML  
```sh
./h2m md2html -i document.md -o result.html
```

### 🔹 Convert HTML → Markdown  
```sh
./h2m html2md -i page.html -o result.md
```

### 🔹 Interactive Mode  
Run the tool without arguments to enter interactive mode:  
```sh
./h2m
```
Type `exit` or press `Ctrl+C` to quit.  

---

### Help
```sh
./h2m --help 
```

## ⚡ Example  
#### 📄 `sample.md` file:  
```md
# Hello World
This is a **Markdown** file.
```
#### 🔄 Convert:  
```sh
./h2m -i sample.md -o sample.html
```
#### 📄 Output `sample.html`:  
```html
<h1>Hello World</h1>
<p>This is a <strong>Markdown</strong> file.</p>
```
---

## 🤝 Contributing  
Fork, modify, and submit a pull request!  

## 🛠️ Maintainer  
👤 **Arouna Ouattara**  
🔗 [LinkedIn](https://www.linkedin.com/in/arouna-ouattara/)  

🚀 **Happy coding!** 😃