package database

import "gorm.io/gorm"

type DBInstance struct {
	DB *gorm.db
}

var DB DBInstance

func ConnectDb(){
	dsn := fmt.SPrintf("host=db userpostgres password=%s dbname=%s port=5432 sslmode=disble TimeZone=Europe/Paris", os.Getenv("postgres_password"), os.Getenv("postgres_db"))
	db, err := gormOpen(postgres.Open(dsn), &gorm.Config{Logger : logger.Default.LogMode(logger.Info),
	})
	if err != nill {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected to database !")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.AutoMigrate(&models.Message{})

	DB = DBInstance {
		DB: db,
	}
}
