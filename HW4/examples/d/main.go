package main

import (
	"log"

	"github.com/boltdb/bolt"
)

// Самое интересное, что это 90% функционала. Еще 9% -- это просто ручное владение транзакциями,
// ничего нового для нас в этом нет
func main() {
	// Открытие или создание базы данных BoltDB
	db, err := bolt.Open("my.db", 0666, nil) // И тот самый 1% -- это опции, по типу таймаутов и прав на доступ
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Panic(err)
		}
	}()

	// Запись данных
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return err
		}

		err = bucket.Put([]byte("key"), []byte("value"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Чтение данных
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("MyBucket"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		value := bucket.Get([]byte("key"))
		log.Println("value:", string(value)) // |out| value: value
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
