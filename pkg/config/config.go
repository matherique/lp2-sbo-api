package config

import (
	"fmt"
	"os"
)

type config struct {
	APP_PORT          string
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_DBNAME   string
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
}

func Read() (*config, error) {
	c := new(config)
	c.APP_PORT = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	c.DATABASE_HOST = os.Getenv("DATABASE_HOST")
	c.DATABASE_PORT = os.Getenv("DATABASE_PORT")
	c.DATABASE_DBNAME = os.Getenv("DATABASE_DBNAME")
	c.DATABASE_USERNAME = os.Getenv("DATABASE_USERNAME")
	c.DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *config) validate() error {
	switch true {
	case c.APP_PORT == "":
		return fmt.Errorf("missing APP_PORT env variable")
	case c.DATABASE_HOST == "":
		return fmt.Errorf("missing DATABASE_HOST env variable")
	case c.DATABASE_DBNAME == "":
		return fmt.Errorf("missing DATABASE_DBNAME env variable")
	case c.DATABASE_USERNAME == "":
		return fmt.Errorf("missing DATABASE_USERNAME env variable")
	case c.DATABASE_PASSWORD == "":
		return fmt.Errorf("missing DATABASE_PASSWORD env variable")
	}

	return nil
}
