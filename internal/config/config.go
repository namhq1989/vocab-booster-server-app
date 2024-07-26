package config

import "errors"

type (
	Server struct {
		RestPort string
		GRPCPort string

		AppName      string
		Environment  string
		IsEnvRelease bool
		Debug        bool

		// Authentication
		FirebaseServiceAccount string
		AccessTokenSecret      string
		AccessTokenTTL         int // seconds

		// MongoDB
		MongoURL    string
		MongoDBName string

		// Redis
		CachingRedisURL string

		// Queue
		QueueRedisURL    string
		QueueUsername    string
		QueuePassword    string
		QueueConcurrency int

		// Sentry
		SentryDSN     string
		SentryMachine string

		// Ably
		AblyAPIKey string

		// Endpoint
		EndpointEnglishHub     string
		EndpointExerciseGrpc   string
		EndpointVocabularyGrpc string
	}
)

func Init() Server {
	cfg := Server{
		RestPort: ":3020",
		GRPCPort: ":3021",

		AppName:     getEnvStr("APP_NAME"),
		Environment: getEnvStr("ENVIRONMENT"),
		Debug:       getEnvBool("DEBUG"),

		FirebaseServiceAccount: getEnvStr("FIREBASE_SERVICE_ACCOUNT"),
		AccessTokenSecret:      getEnvStr("ACCESS_TOKEN_SECRET"),
		AccessTokenTTL:         getEnvInt("ACCESS_TOKEN_TTL"),

		MongoURL:    getEnvStr("MONGO_URL"),
		MongoDBName: getEnvStr("MONGO_DB_NAME"),

		CachingRedisURL: getEnvStr("CACHING_REDIS_URL"),

		QueueRedisURL:    getEnvStr("QUEUE_REDIS_URL"),
		QueueUsername:    getEnvStr("QUEUE_USERNAME"),
		QueuePassword:    getEnvStr("QUEUE_PASSWORD"),
		QueueConcurrency: getEnvInt("QUEUE_CONCURRENCY"),

		SentryDSN:     getEnvStr("SENTRY_DSN"),
		SentryMachine: getEnvStr("SENTRY_MACHINE"),

		AblyAPIKey: getEnvStr("ABLY_API_KEY"),

		EndpointEnglishHub:     getEnvStr("ENDPOINT_ENGLISH_HUB"),
		EndpointExerciseGrpc:   getEnvStr("ENDPOINT_EXERCISE_GRPC"),
		EndpointVocabularyGrpc: getEnvStr("ENDPOINT_VOCABULARY_GRPC"),
	}
	cfg.IsEnvRelease = cfg.Environment == "release"

	// validation
	if cfg.Environment == "" {
		panic(errors.New("missing ENVIRONMENT"))
	}

	if cfg.FirebaseServiceAccount == "" {
		panic(errors.New("missing FIREBASE_SERVICE_ACCOUNT"))
	}

	if cfg.MongoURL == "" {
		panic(errors.New("missing MONGO_URL"))
	}
	if cfg.MongoDBName == "" {
		panic(errors.New("missing MONGO_DB_NAME"))
	}
	if cfg.MongoDBName == "" {
		panic(errors.New("missing MONGO_DB_NAME"))
	}

	if cfg.CachingRedisURL == "" {
		panic(errors.New("missing CACHING_REDIS_URL"))
	}

	if cfg.QueueRedisURL == "" {
		panic(errors.New("missing QUEUE_REDIS_URL"))
	}

	if cfg.AccessTokenSecret == "" {
		panic(errors.New("missing ACCESS_TOKEN_SECRET"))
	}

	if cfg.AblyAPIKey == "" {
		panic(errors.New("missing ABLY_API_KEY"))
	}

	if cfg.EndpointEnglishHub == "" {
		panic(errors.New("missing ENDPOINT_ENGLISH_HUB"))
	}
	if cfg.EndpointExerciseGrpc == "" {
		panic(errors.New("missing ENDPOINT_EXERCISE_GRPC"))
	}
	if cfg.EndpointVocabularyGrpc == "" {
		panic(errors.New("missing ENDPOINT_VOCABULARY_GRPC"))
	}

	return cfg
}
