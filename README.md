# Phone Numer Manger - Exam - JTI
Aplikasi Manajemen Nomor Telephone untuk ujian teknis calon karyawan JTI

## Tech
- Bahasa Pemgrograman: Go (https://go.dev/)
- Framewwork: Chi (https://github.com/go-chi/chi)
- Config Bundle: Apem (https://github.com/karincake/apem)
- Database ORM: Gorm (https://gorm.io/)
- Project Layout: Go Standar (https://github.com/golang-standards/project-layout/)
- Database: MySQL

## Installasi dan Konfigurasi
- Pastikan di PC anda sudah terinstall:
    - svc `git`
    - Bahasa Pemrorgaman `Go`
    - Database server `MySQL`
- Clone, silahkan clone dari repo dengan perintah: `git clone https://github.com/munaja/pnm-exam-jti-be`
- Konfigurasi, terdapat 2 bagian:
    - Buat database pada `MySQL` server yang telah terinstall
    - Migrasi:
        - masuk ke direktori 'pnm-exam-jti-be/cmd/adbmigration`
        - salin file 'config.yml.example' dan beri nama  `config.yml`
        - modifikasi file `config.yml` dan ubah `username`, `password`, dan `database-name` sesuai dnegan konfigurasi database anda.
        - jalankan perintah `go run .`
    - Server Customer:
        - masuk ke direktori 'pnm-exam-jti-be/cmd/customer`
        - salin file 'config.yml.example' dan beri nama  `config.yml`
        - modifikasi file `config.yml` dan sesuaikan bagian berikut:
            - httpConf, terdapat seting `host` dan `port`
            - dbConf, terdapat setting dsn untuk `username`, `password`, dan `database-name` sesuai dnegan konfigurasi database anda.
- Run Server Customer:
    - masuk ke direktori 'pnm-exam-jti-be/cmd/customer`
    - jalankan perintah `go run .` untuk running dengan mode debugging
    - jalankan perintah `go build .` untuk membuild program. Kusus mode ini program harus di ubah access filenya menjadi executable agar dapat di run. Setelah itu ketik `customer`.

