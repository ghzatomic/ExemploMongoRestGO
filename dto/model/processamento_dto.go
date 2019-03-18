package model

import "gopkg.in/mgo.v2/bson"

type Processamento struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	TicketId    string        `bson:"ticketId" json:"ticketId"`
	UsuarioId   string        `bson:"usuarioId" json:"usuarioId"`
}