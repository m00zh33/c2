package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"strings"

	e4 "gitlab.com/teserakt/e4common"

	"github.com/jinzhu/gorm"
)

// IDKey represents an Identity Key in the database given a unique device ID.
type IDKey struct {
	ID        int         `gorm:"primary_key:true"`
	E4ID      []byte      `gorm:"unique;NOT NULL"`
	Key       []byte      `gorm:"NOT NULL"`
	TopicKeys []*TopicKey `gorm:"many2many:idkeys_topickeys;"`
}

// TopicKey represents
type TopicKey struct {
	ID     int      `gorm:"primary_key:true"`
	Topic  string   `gorm:"unique;NOT NULL"`
	Key    []byte   `gorm:"NOT NULL"`
	IDKeys []*IDKey `gorm:"many2many:idkeys_topickeys;"`
}

// This function is responsible for the
func (s *C2) dbInitialize() error {
	s.logger.Log("msg", "Database Migration Started.")
	// TODO: better DB migration logic.
	// TODO: transactions?
	//tx := s.db.Begin()

	if result := s.db.AutoMigrate(&IDKey{}); result.Error != nil {
		//tx.Rollback()
		return result.Error
	}
	if result := s.db.AutoMigrate(&TopicKey{}); result.Error != nil {
		//tx.Rollback()
		return result.Error
	}
	/*if err := tx.Commit().Error; err != nil {
		return err
	}*/
	s.logger.Log("msg", "Database Migration Finished.")
	return nil
}

func (s *C2) dbInsertIDKey(id, key []byte) error {
	protectedkey, err := e4.Encrypt(s.keyenckey[:], nil, key)
	if err != nil {
		return err
	}

	var idkey IDKey
	s.db.Where(&IDKey{E4ID: id}).First(&idkey)
	if s.db.NewRecord(idkey) {
		idkey = IDKey{E4ID: id, Key: protectedkey}

		if result := s.db.Create(&idkey); result.Error != nil {
			return result.Error
		}
	} else {
		idkey.Key = protectedkey

		if result := s.db.Model(&idkey).Updates(idkey); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (s *C2) dbInsertTopicKey(topic string, key []byte) error {
	protectedKey, err := e4.Encrypt(s.keyenckey[:], nil, key)
	if err != nil {
		return err
	}

	var topicKey TopicKey
	s.db.Where(&TopicKey{Topic: topic}).First(&topicKey)

	if s.db.NewRecord(topicKey) {
		topicKey = TopicKey{Topic: topic, Key: protectedKey}
		if result := s.db.Create(&topicKey); result.Error != nil {
			return result.Error
		}
	} else {
		topicKey.Key = protectedKey
		if result := s.db.Model(&topicKey).Updates(topicKey); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (s *C2) dbGetIDKey(id []byte) ([]byte, error) {
	var idkey IDKey
	result := s.db.Where(&IDKey{E4ID: id}).First(&idkey)
	if gorm.IsRecordNotFoundError(result.Error) {
		// TODO: do we return an error for this?
		return nil, errors.New("ID not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	if !bytes.Equal(id, idkey.E4ID) {
		return nil, errors.New("Internal error: struct not populated but GORM indicated success")
	}
	clearkey, err := e4.Decrypt(s.keyenckey[:], nil, idkey.Key[:])
	if err != nil {
		return nil, err
	}
	return clearkey, nil
}

func (s *C2) dbGetTopicKey(topic string) ([]byte, error) {
	var topickey TopicKey
	result := s.db.Where(&TopicKey{Topic: topic}).First(&topickey)
	if gorm.IsRecordNotFoundError(result.Error) {
		return nil, errors.New("Topic not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	if strings.Compare(topickey.Topic, topic) != 0 {
		return nil, errors.New("Internal error: struct not populated but GORM indicated success")
	}
	clearkey, err := e4.Decrypt(s.keyenckey[:], nil, topickey.Key[:])
	if err != nil {
		return nil, err
	}
	return clearkey, nil
}

func (s *C2) dbDeleteIDKey(id []byte) error {
	var idkey IDKey

	if result := s.db.Where(&IDKey{E4ID: id}).First(&idkey); result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return errors.New("ID not found, nothing to delete")
		}
		return result.Error

	}
	// safety check:
	if !bytes.Equal(idkey.E4ID, id) {
		return errors.New("Single record not populated correctly; preventing whole DB delete")
	}
	tx := s.db.Begin()
	if result := tx.Model(&idkey).Association("TopicKeys").Clear(); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result := tx.Delete(&idkey); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (s *C2) dbDeleteTopicKey(topic string) error {
	var topicKey TopicKey
	if result := s.db.Where(&TopicKey{Topic: topic}).First(&topicKey); result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return errors.New("ID not found, nothing to delete")
		}
		return result.Error
	}
	if topicKey.Topic != topic {
		return errors.New("Single record not populated correctly; preventing whole DB delete")
	}
	tx := s.db.Begin()
	if result := s.db.Model(&topicKey).Association("IDKeys").Clear(); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result := s.db.Delete(&topicKey); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (s *C2) dbCountIDKeys() (int, error) {
	var idkey IDKey
	var count int
	if result := s.db.Model(&idkey).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (s *C2) dbCountTopicKeys() (int, error) {
	var topickey TopicKey
	var count int
	if result := s.db.Model(&topickey).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (s *C2) dbGetIDListHex() ([]string, error) {
	var idkeys []IDKey
	var hexids []string
	if result := s.db.Find(&idkeys); result.Error != nil {
		return nil, result.Error
	}

	for _, idkey := range idkeys {
		hexids = append(hexids, hex.EncodeToString(idkey.E4ID[0:]))
	}

	return hexids, nil
}

func (s *C2) dbGetTopicsList() ([]string, error) {
	var topickeys []TopicKey
	var topics []string
	if result := s.db.Find(&topickeys); result.Error != nil {
		return nil, result.Error
	}

	for _, topickey := range topickeys {
		topics = append(topics, topickey.Topic)
	}

	return topics, nil
}

/* -- M2M Functions -- */

// This function links a topic and an id/key. The link is created in both
// directions (IDkey to Topics, Topic to IDkeys).
func (s *C2) dbLinkIDTopic(id []byte, topic string) error {

	var idkey IDKey
	var topickey TopicKey

	if err := s.db.Where(&IDKey{E4ID: id}).First(&idkey).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("ID Key not found, cannot link to topic")
		}
		return err

	}
	if err := s.db.Where(&TopicKey{Topic: topic}).First(&topickey).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("ID Key not found, cannot link to IDkey")
		}
		return err
	}

	tx := s.db.Begin()

	if err := tx.Model(&idkey).Association("TopicKeys").Append(&topickey).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

// This function removes the relationship between a Topic and an ID, but
// does not delete the Topic or the ID.
func (s *C2) dbUnlinkIDTopic(id []byte, topic string) error {

	var idkey IDKey
	var topickey TopicKey

	if err := s.db.Where(&IDKey{E4ID: id}).First(&idkey).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("ID Key not found, cannot unlink from topic")
		}
		return err
	}
	if err := s.db.Where(&TopicKey{Topic: topic}).First(&topickey).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("ID Key not found, cannot unlink from IDkey")
		}
		return err
	}

	tx := s.db.Begin()

	if err := tx.Model(&idkey).Association("TopicKeys").Delete(&topickey).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(&IDKey{E4ID: id}).First(&idkey).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return errors.New("ID/Client appears to have been deleted, this is just an unlink")
		}
	}
	if err := tx.Where(&TopicKey{Topic: topic}).First(&topickey).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return errors.New("Topic appears to have been deleted, this is just an unlink")
		}
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (s *C2) dbCountTopicsForID(id []byte) (int, error) {

	var idkey IDKey

	if err := s.db.Where(&IDKey{E4ID: id}).First(&idkey).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return 0, errors.New("ID Key not found, cannot link to topic")
		}
		return 0, err
	}

	count := s.db.Model(&idkey).Association("TopicKeys").Count()
	return count, nil
}

func (s *C2) dbGetTopicsForID(id []byte, offset int, count int) ([]string, error) {

	var idkey IDKey
	var topickeys []TopicKey
	var topics []string

	if err := s.db.Where(&IDKey{E4ID: id}).First(&idkey).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("ID Key not found, cannot link to topic")
		}
		return nil, err
	}

	if err := s.db.Model(&idkey).Offset(offset).Limit(count).Related(&topickeys, "TopicKeys").Error; err != nil {
		return nil, err
	}

	for _, topickey := range topickeys {
		topics = append(topics, topickey.Topic)
	}

	return topics, nil
}

func (s *C2) dbCountIDsForTopic(topic string) (int, error) {
	var topickey TopicKey

	if err := s.db.Where(&TopicKey{Topic: topic}).First(&topickey).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return 0, errors.New("Topic key not found, cannot link to ID")
		}
		return 0, err
	}

	count := s.db.Model(&topickey).Association("IDKeys").Count()
	return count, nil
}

func (s *C2) dbGetIdsforTopic(topic string, offset int, count int) ([]string, error) {

	var topickey TopicKey
	var idkeys []IDKey
	var hexids []string

	if err := s.db.Where(&TopicKey{Topic: topic}).First(&topickey).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("ID Key not found, cannot link to topic")
		}
		return nil, err
	}

	if err := s.db.Model(&topickey).Offset(offset).Limit(count).Related(&idkeys, "IDKeys").Error; err != nil {
		return nil, err
	}

	for _, idkey := range idkeys {
		hexids = append(hexids, hex.EncodeToString(idkey.E4ID[:]))
	}

	return hexids, nil
}