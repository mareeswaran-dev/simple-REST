package app

import (
	"fmt"

	"user-crud-api/config"
	"user-crud-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App holds all the application dependencies and configurations
type App struct {
	Config *config.Config
	DB     *gorm.DB
}

// New creates a new application instance
func New(cfg *config.Config) (*App, error) {
	app := &App{
		Config: cfg,
	}

	// Initialize database connection
	if err := app.initDB(); err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	return app, nil
}

// initDB initializes the database connection
func (a *App) initDB() error {
	var err error
	a.DB, err = gorm.Open(postgres.Open(a.Config.GetDBConnString()), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Auto migrate the schema
	if err := a.DB.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("error migrating database: %v", err)
	}

	fmt.Println("Successfully connected to database!")
	return nil
}

// Close closes all application resources
func (a *App) Close() error {
	sqlDB, err := a.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
