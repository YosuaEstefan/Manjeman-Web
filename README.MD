# Golang Deployment Web API

Sistem Manajemen Barang menggunakan **Golang + Gin** yang telah berhasil dideploy ke Vercel.

## 📂 **Struktur Proyek**
```
Mnajemen Sistem/
│── main.go
│── go.mod
│── go.sum
│── vercel.json
│── README.md
│── controllers/
│   ├── belanjaController.go
│── models/
│   ├── belanja.go
│── api/
│   ├── handlers.go
```
---

## 📌 **1. Dependencies / Library yang Dibutuhkan**
Jalankan perintah berikut untuk menginstall dependencies yang dibutuhkan:

```sh
go mod init tugas-deployment
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/mux
npm install -g vercel
vercel --version
vercel login
```

---

## 📌 **2. Testing API dengan Postman**

### 🟢 **Get data belanjaan**
- **Method:** GET  
- **URL:** `localhost:8080/api/belanja`
- **Screenshoot:** [Google Drive](https://drive.google.com/drive/folders/1TGf9IiP0FU7JhQabVY_ur2VXYESW4yfa?usp=sharing)

### 🟢 **Get data belanjaan berdasarkan ID**
- **Method:** GET  
- **URL:** `localhost:8080/api/belanja/13`
- **Screenshoot:** [Google Drive](https://drive.google.com/drive/folders/1DtZvOjnOiVfdJFIsuEP9LnL_aWeKdTaj?usp=drive_link)

### 🟢 **Post data belanjaan**
- **Method:** POST  
- **URL:** `localhost:8080/api/belanja`
- **Body:**
  ```json
  {
      "id": 15,
      "nama": "Minyak Kelapa",
      "jenis": "Makanan",
      "jumlah": 4
  }
  ```
- **Screenshoot:** [Google Drive](https://drive.google.com/drive/folders/1MZSyxb_SmDOcQ8m26EYIJEyMN63EhomE?usp=drive_link)

---

## 📌 **3. Testing API dengan Vercel**

### 🟢 **Link Deployment**
🔗 [Vercel Deployment](https://vercel.com/yosua-estefan-silitongas-projects/manejemenbarang/6RM3bTL5iqHu6Dkcm4LG4gJuHn5e)

### 🟢 **Get data belanjaan**
- **Method:** GET  
- **URL:** `https://manejemenbarang-menz1xgg8-yosua-estefan-silitongas-projects.vercel.app/api/belanja`
- **Screenshoot:** [Google Drive](https://drive.google.com/drive/folders/1J7NnJwYfRYQ72v82MWktbgno3XbNFqFi?usp=drive_link)

### 🟢 **Get data belanjaan berdasarkan ID**
- **Method:** GET  
- **URL:** `https://manejemenbarang-menz1xgg8-yosua-estefan-silitongas-projects.vercel.app/api/belanja/9`
- **Screenshoot:** [Google Drive](https://drive.google.com/drive/folders/1Lsg4vGKE3JaJxOgq7B7AVP1CiYvhJnQC?usp=drive_link)

### 🟢 **Post data belanjaan**
- **Method:** POST  
- **URL:** `https://manejemenbarang.vercel.app/api/belanja`
- **Body:**
  ```json
  {
      "id": 18,
      "nama": "Minyak gula",
      "jenis": "Makanan",
      "jumlah": 20
  }
  ```
- **Screenshoot:** [Google Drive](https://drive.google.com/drive/folders/1aeS4L8ZNsqTafyblU2Fv4PhK6Q-Tvyp7?usp=drive_link)

---

## 📌 **4. Kesimpulan**
Dengan mengikuti langkah-langkah di atas, Anda telah berhasil **mendeploy API Golang ke Vercel** dan mengujinya menggunakan Postman. Semoga bermanfaat dan selamat mencoba! 🚀

