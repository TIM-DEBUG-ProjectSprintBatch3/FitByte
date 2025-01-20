# Desain Sistem

Karena mempraktikan `microservice` pada Projek 2 FitByte ini, kemudian memisahkan service menjadi 3 bagian

1. Readonly (Hanya menangani operasi `GET` atau query `SELECT` sahaja) dengan port 8081
2. Manipulasi Data (Menangani operasi `PATCH` dan `DELETE` atau query `UPDATE` dan `DELETE`) dengan port 8082
3. File (Menangani operasi file) dengan port 8083
4. API Gateway Port 8080

Dalam praktiknya, dapat diaplikasikan pada:

1. 1 Service - 1 Instance / Server
2. 3 Service - 1 Instance / Server
Untuk opsi nomor dua, kemudian akan disediakan berupa API Gateway yang akhirnya akan tetap menjadi 1 Port 8080

## Melakukan debugging pada suatu service

1. Ganti direktori pada folder sesuai dengan service yang dituju, `cd [ur_awesome_service]`
2. Lakukan `go run main.go`

Jika ingin melakukan multiple services, mohon build terlebih dahulu tiap service

```sh
# For build, run this command
go build -o .build/<name-of-build.extension>

# NOTE: it is important to put the build inside of the .build folder
# to ensure the gitignore caught up with the files

# After build go application
cd .build/<name-of-build.extension>
```

Selanjutnya buka aplikasi tersebut manual satu-persatu

## Jalankan aplikasi dengan docker-compose

1. Kembali pada direktori `root` projek FitByte
2. Pastikan port tiap service sudah diarahkan
3. Lakukan `docker-compose up -d` atau jika ingin rebuild tambahkan flag `--build`
4. --belum ada swagger, lewat ke nomor 5--Jika ingin menambahkan swagger, pada rute swagger yang ada di `route.go` buat menjadi seperti ini `swagger/readonly`
5. Dengan aplikasi seperti Postman, arahkan pada `localhost:8080/[ur_awesome_api_route]`

## Info lebih lanjut

Sila mengunjungi pada [GogoManager_Project1](https://github.com/prasasdi/projeksprint_p1)
