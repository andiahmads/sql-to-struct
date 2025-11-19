# 🛠️ SQLStruct — SQL → Go Struct Generator  
Konversi otomatis *CREATE TABLE* menjadi struct Golang, lengkap dengan **warna**, **CLI interface**, dan **file picker interaktif**.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![CLI](https://img.shields.io/badge/CLI-Tool-orange?style=for-the-badge)

---

## ✨ Fitur Utama
- ⚡ **Convert SQL → Go Struct** dari file SQL berisi `CREATE TABLE`
- 🎨 **CLI berwarna** menggunakan `fatih/color`
- 📁 **File Picker Interaktif** untuk memilih file SQL dengan daftar folder
- 🔍 Auto-detect tipe SQL → tipe Golang
- 🎯 Mendukung field menggunakan **backtick** atau tanpa backtick
- 📦 Dibangun menggunakan **Cobra**

---

## 🎥 Demo Singkat

```bash
$ sqlstruct convert

SQL File Path (press Enter to browse): ./your-migration-folder

? Choose file in ./your-migration-folder:
> users.sql
  roles.sql
  logs.sql

==== Generated Struct ====
type Users struct {
    Id        string `json:"id"`
    FullName  string `json:"full_name"`
    Email     string `json:"email"`
    RoleId    string `json:"role_id"`
    Password  string `json:"password"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

-----


## 📦 Instalasi

### 1. Clone Repository
```bash
git clone https://github.com/your-username/sqlstruct.git
cd sqlstruct



🚀 Cara Menjalankan
./sqlstruct convert

