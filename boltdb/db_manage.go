package boltdb


import (
	"log"
	"github.com/boltdb/bolt"
	"fmt"
	"os"
)

const DB_NAME  = "BLC.db"
const DB_TABLE_NAME_BLOCKS  = "blocks"
const DB_TABLE_NAME_TIP  = "tip"

const DB_TABLE_NAME_TIP_KEY_LASTBLOCKHASH = "lastBlcHash"//用于存放最后那块hash

/**
	bolt 数据库管理
 */
type DBManger struct {
	boltdb *bolt.DB
}

func OpenDB(dbFile string) *DBManger  {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	dbManager := new(DBManger)
	dbManager.boltdb = db
	return dbManager
}

func (dbManager *DBManger)IsExistDBTable(tableName string) bool  {
	var isExisted bool
	err := dbManager.boltdb.Update(func(tx *bolt.Tx) error {

		//获取表，看是否存在
		table := tx.Bucket([]byte(tableName))
		if table != nil {
			isExisted = true
		}
		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}
	return isExisted

}

/**
	插入新数据
	key: 以hash值作为key
 */

func (dbManager *DBManger)InsertData(tableName string, key[]byte, data[] byte ) {

	var err error
	if dbManager == nil {
		return;
	}
	db := dbManager.boltdb
	err = db.Update(func(tx *bolt.Tx) error {

		//获取表，看是否存在
		table := tx.Bucket([]byte(tableName))

		if table == nil {
			// 创建BlockBucket表
			table, err = tx.CreateBucket([]byte(tableName))
			if err != nil {
				return fmt.Errorf("db create table: %s", err)
			}

		//	fmt.Printf("db create a new table %s successfully \n ",tableName)
		}

		// 往表里面存储数据
		err := table.Put(key,data)
		if err != nil {
			log.Panic("db insert data failed......")
		}

		// 返回nil，以便数据库处理相应操作
		return nil
	})

	//更新失败
	if err != nil {
		log.Panic(err)
	}
}

/**
	查询数据
	key: 以hash值作为key
 */
func (dbManager *DBManger) SelectData(tableName string, key[]byte) []byte {
	db := dbManager.boltdb

	var getData[] byte
	err := db.View(func(tx *bolt.Tx) error {
		//获取表，看是否存在
		table := tx.Bucket([]byte(tableName))
		if table == nil {
			fmt.Println("数据库中不存在表:  ",tableName)
			return nil
		}
		getData = table.Get(key)
		return nil
	})

	if err != nil{
		fmt.Println("failed to select data")
		return nil
	}
	return getData
}

func (dbManager *DBManger) CloseDB() {
	fmt.Printf("db close \n")
	dbManager.boltdb.Close()
}

func DBIsExist() bool  {
	_, err := os.Stat(DB_NAME);
	return os.IsExist(err)
}