Blockchain-go

│  
├── main.go                # File utama yang berisi kode untuk blockchain dan komunikasi antar node  
│  
├── blockchain/            # Direktori untuk logika blockchain  
│   ├── block.go           # Definisi struktur blok dan fungsi terkait blok  
│   └── blockchain.go      # Definisi struktur node dan fungsi terkait blockchain  
│  
├── network/               # Direktori untuk logika jaringan  
│   ├── server.go          # Kode untuk server yang mendengarkan koneksi dari node lain  
│   └── client.go          # Kode untuk klien yang mengirim data ke node lain  
│  
├── go.mod                 # File modul Go  
│  
└── README.md              # File dokumentasi proyek  
