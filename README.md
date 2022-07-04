Info Test BE
1. Dengan NodeJS (atau Golang), agar saudara membuat Rest API CRUD User dan User Login.
2. Jika menggunakan NodeJS maka disarankan menggunakan ExpressJS. Database bebas, tetapi disarankan MongoDB.
3. User Login digunakan user (username, password) untuk mengakses API CRUD (token, tetapi mendapatkan nilai tambahan jika menggunakan refresh token).
4. Bikin 2 users dengan role: 1 Admin, 1 User.
5. Admin bisa melakukan/mengakses semua API CRUD, sedangkan User hanya bisa mengakses data user bersangkutan saja (Read)
6. Implementasi arstektur Microservices, menggunakan Kubernetes dengan Docker container deploy di VPS (1 node dengan beberapa pod di dalamnya). Bagi yang belum memiliki VPS, maka cukup (a) menyiapkan semua YML agar aplikasi bisa dijalankan secara containerize dan siap di deploy di Kubernetes dan (b) di-deploy di lokal dan sertakan screenshoot.
7. Upload source code ke Github beserta script YML Kubernetes.
8. Bikin dokumentasi API nya (Postman atau Swagger) yang bisa diakses ke server Rest API nya.
9. Bikin diagram arstektur nya yang menjelaskan flow API CRUD dan Login.
10. Lampirkan credential Admin di Readme.

## TODO
- [x] User Login digunakan user (username, password) untuk mengakses API CRUD (token, tetapi mendapatkan nilai tambahan jika menggunakan refresh token).
- [x] Bikin 2 users dengan role: 1 Admin, 1 User.
- [x] Admin bisa melakukan/mengakses semua API CRUD, sedangkan User hanya bisa mengakses data user bersangkutan saja (Read)
- [x] Implementasi arstektur Microservices, menggunakan Kubernetes dengan Docker container deploy di VPS (1 node dengan beberapa pod di dalamnya). Bagi yang belum memiliki VPS, maka cukup (a) menyiapkan semua YML agar aplikasi bisa dijalankan secara containerize dan siap di deploy di Kubernetes dan (b) di-deploy di lokal dan sertakan screenshoot.
- [x] Bikin dokumentasi API nya (Postman atau Swagger) yang bisa diakses ke server Rest API nya.
- [x] Bikin diagram arstektur nya yang menjelaskan flow API CRUD dan Login.
- [x] Lampirkan credential Admin di Readme.

Next: Add UserModel

## User credential

BASE_URL: http://103.171.84.154:30183

username: admin
password: secret

username: user
password: secret

## Documentation

- API docs is available at `golang_simple_crud.postman_collection.json` in this repo
- Diagram is available at `testBackend.drawio` and `testBackend.drawio.pdf` for the pdf version
