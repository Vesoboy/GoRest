package DataContext

import (
	"fmt"
	"log"
	"restApp/Models"

	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AddWallet(db *gorm.DB, Id uuid.UUID, allSum float64) error {

	wallet := Models.Wallets{ValletId: Id, AllSum: allSum}

	if err := db.Create(&wallet).Error; err != nil {
		return fmt.Errorf("Ошибка создания кошелька: %v", err)
	}
	return nil

}

func UpdateWallet(db *gorm.DB, Id uuid.UUID, Operation string, amount float64, allsum float64) error {

	wallet := Models.Wallets{ValletId: Id, OperationType: Operation, Amount: amount, AllSum: allsum}

	if err := db.Save(&wallet).Error; err != nil {
		return fmt.Errorf("Ошибка при обновлении кошелька: %v", err)
	}
	return nil
}

func GetWallets(db *gorm.DB, path string) (*Models.Wallets, error) {

	var wallet Models.Wallets
	result := db.First(&wallet, "vallet_id = ?", path)
	if result.Error != nil {
		return nil, fmt.Errorf("Кошелька нет в бд: %v", result.Error)
	}
	return &wallet, nil
}

func DataContextDB() *gorm.DB {

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}

	dsn := GormDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST_DOCKER"),
		os.Getenv("DB_PORT"), os.Getenv("DB_DBNAME"), os.Getenv("DB_SSLMODE"))
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		dsn = GormDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"), os.Getenv("DB_DBNAME"), os.Getenv("DB_SSLMODE"))
		gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Через Docker подключение не удалось, но локальное подключение к база данных подключена успешно!")
			return gormDB
		}
		log.Fatalf("Ошибка подключения к базе через Docker и локальное подключение: %v", err)
	}

	// Миграция
	err = AutoMigrate(gormDB)
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	fmt.Println("Миграция выполнена успешно!")
	return gormDB
}

func GormDB(db_user, db_password, db_host, db_port, db_dbname, db_sslmode string) string {

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s ",
		db_host, db_user, db_password, db_port, db_dbname, db_sslmode)

	return dsn
}
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Models.Wallets{})
}
