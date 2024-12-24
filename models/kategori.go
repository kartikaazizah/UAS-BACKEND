package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Kategori struct {
    ID           primitive.ObjectID `bson:"_id,omitempty"`
    KategoriNama string             `bson:"kategori_nama"`
}


