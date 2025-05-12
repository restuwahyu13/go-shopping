# GO Checkout API

Berikut adalah simple simulasi api menggunakan golang, cara ini masih dikembangkan kembali agar bisa lebih bagus lagi terkait flow konsepnya.

## Penjelasan Database

Terkait database ini hanyalah contoh sederhana dan masih bisa di kembangkan kembali, mungkin perlu penambahan table baru contoh seperti: voucher,  promosi, address, shipping dan lain - lain.

- User: Digunakan untuk menyimpan data pengguna setelah berhasil melakukan proses registrasi.
- Bank: Digunakan untuk menyimpan data master bank, yang diperlukan untuk menentukan apakah bank tersebut merupakan rekanan atau bukan dalam proses pembayaran.
- Courier: Digunakan untuk menyimpan data master kurir, yang diperlukan untuk menentukan apakah kurir tersebut merupakan rekanan atau bukan dalam proses pengiriman barang.
- Promotion Rules: Digunakan untuk menetapkan standar pola promosi yang akan diterapkan pada suatu produk.
- Product: Digunakan untuk menyimpan data master produk, yang berfungsi untuk menampilkan daftar produk.
- Product Item: Digunakan untuk menyimpan data master item produk, yang masih terkait dengan modul Product karena adanya hubungan keterkaitan antara keduanya.
- Product Config: Digunakan untuk menerapkan konfigurasi promosi, baik untuk semua produk maupun spesifik item produk tertentu.
- Payment: Digunakan untuk menyimpan data pembayaran pengguna setelah melakukan pemesanan produk terkait.
- Order: Digunakan untuk menyimpan data pesanan, mencakup produk-produk yang pernah dibeli oleh pengguna.
- Order Item: Digunakan untuk menyimpan data item produk yang pernah dibeli oleh pengguna tersebut.

## Flow Proses

1. Login menggunakan email & password menggunakan api -> `/api/v1/auth/login`
2. Jika ingin melakukan checkout produk, bisa menungganakn api -> `/api/v1/checkout`, api ini digunakan untuk menambah checkout barang,
menghapus checkout barang dan membuat order pesanan berdasarkan barang yang dicheckout, better buat api terpisah untuk order, karena ini hanya simulasi jadi saya gabungkan.
3. Jika anda ingin menerapkan promosi untuk suatu produk anda bisa lihat di rules promosion, untuk cara pengunaannya.

## Cara Menjalankan Docker

```sh
$ docker-compose up -d | make up
```

## Cara Menjalankan Aplikasi

```sh
$ go run --race -v ./internal/cmd | make dev
```

## Promotion Rules
Rules ini bersifat generic jadi bisa anda dunakan untuk keperluan seperti pembuatan promo, diskon, gift, beli gratis 1 dan lain - lain, Pada sebuah produk.

```js
  {
    name: 'Beli MacBook Pro gratis Rasberry Pi B',
    promotion_rules: JSON.stringify([
      { key: 'B', value: 1 },
      { key: 'EQ', value: '=' },
      { key: 'F', value: productItems[3]?.id },
    ])
    },
    {
    name: 'Beli 3 Google Home hanya bayar 2 harga',
    promotion_rules: JSON.stringify([
      { key: 'B', value: 3 },
      { key: 'EQ', value: '=' },
      { key: 'PR', value: 2 },
    ])
    },
    {
    name: 'Beli 3 Alexa Speaker dapatkan discount 10% di semua Alexa Speaker',
    promotion_rules: JSON.stringify([
      { key: 'B', value: 3 },
      { key: 'EQ', value: '=' },
      { key: 'D', value: '10' },
      { key: 'PC', value: '%' },
      { key: 'P', value: 'SAME_BRAND_SIMILAR_CATEGORY' },
    ]),
    }
```

## Run migration Database error
Jalankan fungsi ini  menggunakan postgres

```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```