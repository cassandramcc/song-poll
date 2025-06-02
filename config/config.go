package config

type Config struct {
	MongoConfig *MongoConfig
}

type MongoConfig struct {
	Database          string
	ArtistsCollection string
	URI               string
}

func CreateConfig() *Config {
	return &Config{
		MongoConfig: &MongoConfig{
			Database:          "Song-Poll",
			ArtistsCollection: "artists",
			URI:               "mongodb://localhost:27017",
		},
	}
}

func CreateTestingConfig() *Config {
	return &Config{
		MongoConfig: &MongoConfig{
			Database:          "Song-Poll-Test",
			ArtistsCollection: "artists-test",
			URI:               "mongodb://localhost:27017",
		},
	}
}
