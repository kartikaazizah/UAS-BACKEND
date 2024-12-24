package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// TemplateModul defines the structure of the template_modul collection
type TemplateModul struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IDJenisUser primitive.ObjectID `bson:"id_jenis_user" json:"id_jenis_user"`
	IDModul     primitive.ObjectID `bson:"id_modul" json:"id_modul"`
}
