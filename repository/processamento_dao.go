package repository

import (
	. "ExemploMongoRest1/dto/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ProcessamentoDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "processamentoDTO"
)


func (m *ProcessamentoDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *ProcessamentoDAO) GetAll() ([]Processamento, error) {
	var processamentos []Processamento
	err := db.C(COLLECTION).Find(bson.M{}).All(&processamentos)
	return processamentos, err
}

func (m *ProcessamentoDAO) GetByID(id string) (Processamento, error) {
	var processamento Processamento
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&processamento)
	return processamento, err
}

func (m *ProcessamentoDAO) Create(processamento Processamento) error {
	err := db.C(COLLECTION).Insert(&processamento)
	return err
}

func (m *ProcessamentoDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *ProcessamentoDAO) Update(id string, processamento Processamento) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &processamento)
	return err
}