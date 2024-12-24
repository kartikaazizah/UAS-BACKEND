package main

import (
    "log"
    "BEKEN_UAS_PRAK/database"
    "BEKEN_UAS_PRAK/routes"
)

func main() {
    // Koneksi ke MongoDB
    database.Connect()

    // Atur rute
    r := routes.SetupRouter()

    // Jalankan server
    log.Println("Server running on http://localhost:5000")
    err := r.Run(":5000")
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}