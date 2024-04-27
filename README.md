# Tubes2_GoD

Tubes 2 Strategi Algoritma 2024 : Permainan WikiRace menggunakan Algoritma BFS dan IDS

<div style="text-align: justify;">
Breadth-First Search (BFS):
BFS adalah algoritma pencarian yang mengunjungi simpul (node) terdekat terlebih dahulu sebelum mengunjungi simpul yang lebih jauh.
Algoritma ini menggunakan pendekatan "level by level", dimulai dari simpul awal (start node), lalu mengunjungi semua simpul tetangga sebelum melanjutkan ke simpul-simpul yang lebih jauh.
BFS biasanya digunakan ketika mencari solusi yang terdekat atau ketika ingin menemukan jalur terpendek dari simpul awal ke simpul tujuan.

Iterative Deepening Search (IDS):
IDS adalah modifikasi dari Depth-First Search (DFS), tetapi dengan keuntungan dari tidak memerlukan memori yang lebih banyak untuk menyimpan informasi untuk seluruh simpul yang belum diperiksa.
Algoritma ini secara berulang-ulang menjalankan DFS dengan batasan kedalaman (depth limit) yang semakin bertambah pada setiap iterasi.
IDS efektif untuk ruang pencarian yang besar atau tidak terbatas dengan menyatukan keunggulan dari DFS (penggunaan memori yang minimal) dan BFS (penemuan solusi yang dangkal terlebih dahulu).
</div>
   
## Requirement :

1. Go
2. Nodejs

## Cara instalasi dan nge-Run :

1. clone repo </br>
   `git clone https://github.com/rafimaliki/Tubes2_GoD`
2. masuk ke root dir </br>
   `cd Tubes2_GoD`
3. setup **BACKEND** </br>
   `cd backend` </br>
   `go run main.go` buat nyalain backend nya </br> kalo berhasil nge run dia bakal ada promptnya dan ada "Listening and serving HTTP on :8080" di terminal
4. setup **FRONTEND** </br>
   buka terminal baru di root dir repo </br>
   `cd frontend` </br>
   `npm i` buat install node module (ini cuman sekali aja di awal)</br>
   `npm run dev` buat jalanin front end
5. akses front end di </br>
   `http://localhost:5175/`

## Author

1. Muhammad Zaki (13522136)
2. Ahmad Rafi Maliki (13522137)
3. Muhammad Dzaki Arta (13522149)
