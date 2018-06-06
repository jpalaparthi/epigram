package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	Con    string
	Sess   *mgo.Session
	DBName string
}

func New(conn string, db string) (dbs *Session, err error) {
	dbs = &Session{}
	dbs.Sess, err = mgo.Dial(conn)

	if err != nil {
		return nil, err
	} else {
		dbs.DBName = db
	}
	return dbs, nil
}

//start mongo session

func (s *Session) Insert(col string, data interface{}) error {
	err := s.Sess.DB(s.DBName).C(col).Insert(data)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s *Session) UpdateByID(col string, id interface{}, data interface{}) error {
	err := s.Sess.DB(s.DBName).C(col).UpdateId(id, data)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s *Session) DeleteByID(col string, id string) error {
	_id := bson.ObjectIdHex(id)
	err := s.Sess.DB(s.DBName).C(col).RemoveId(_id)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s *Session) DeleteAll(col string) error {
	_, err := s.Sess.DB(s.DBName).C(col).RemoveAll(nil)
	if err != nil {
		return err
	} else {
		return nil
	}

}

// FindByID fetches only one record based on the id provided
func (s *Session) FindByID(col string, id string) (map[string]interface{}, error) {
	_id := bson.ObjectIdHex(id)
	var result map[string]interface{}
	err := s.Sess.DB(s.DBName).C(col).FindId(_id).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//FindByQuery fetches one record based on the query
func (s *Session) FindByQuery(col string, query map[string]interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.Sess.DB(s.DBName).C(col).Find(bson.M(query)).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//FindByQueryAll fetches all records based on the query provided. Query is key value piars
func (s *Session) FindByQueryAll(col string, query map[string]interface{}) ([]interface{}, error) {
	var result []interface{}
	err := s.Sess.DB(s.DBName).C(col).Find(bson.M(query)).Sort("-_id").All(&result)

	if err != nil {
		return result, err
	}
	return result, nil
}

// FindByRegEx fetches all records based on the supplied reg expression
func (s *Session) FindByRegEx(col, key, value string) ([]interface{}, error) {
	var result []interface{}
	err := s.Sess.DB(s.DBName).C(col).Find(bson.M{key: &bson.RegEx{Pattern: value, Options: "i"}}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FindAll fetches all records in the database
func (s *Session) FindAll(col string) ([]interface{}, error) {
	var result []interface{}
	err := s.Sess.DB(s.DBName).C(col).Find(nil).Sort("-_id:").All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FindAllByLimitOffset fetches all records based on the given limit and offset values.
func (s *Session) FindAllByLimitOffset(col string, skip, limit int) ([]interface{}, error) {
	var result []interface{}
	err := s.Sess.DB(s.DBName).C(col).Find(nil).Skip(skip).Limit(limit).Sort("-_id:").All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
