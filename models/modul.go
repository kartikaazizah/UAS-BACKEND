// package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

// type Modul struct {
//     ID         primitive.ObjectID `bson:"_id,omitempty"`
//     NamaModul  string             `bson:"nm_modul"`
//     Kategori   primitive.ObjectID `bson:"kategori"`
//     IsAktif    bool               `bson:"is_aktif"`
//     GbrIcon    string             `bson:"gbr_icon"`
// }

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Modul struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	nm_modul string             `bson:"nm_modul" json:"nm_modul"`
	Kategori primitive.ObjectID `bson:"kategori" json:"kategori"`
	IsAktif  bool               `bson:"is_aktif" json:"is_aktif"`
	gbr_icon string             `bson:"gbr_icon" json:"gbr_icon"`
}
