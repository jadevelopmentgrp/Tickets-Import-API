module github.com/jadevelopmentgrp/Tickets-Import-API

go 1.22.6

// replace github.com/TicketsBot-cloud/database => ../database

require (
	github.com/BurntSushi/toml v1.2.1
	github.com/apex/log v1.1.2
	github.com/caarlos0/env/v11 v11.2.2
	github.com/gin-gonic/contrib v0.0.0-20191209060500-d6e26eeaa607
	github.com/gin-gonic/gin v1.9.1
	github.com/go-playground/validator/v10 v10.14.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-redis/redis_rate/v9 v9.1.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.6.0
	github.com/jackc/pgconn v1.14.3
	github.com/jackc/pgtype v1.14.4
	github.com/jackc/pgx/v4 v4.18.3
	github.com/jadevelopmentgrp/Tickets-Archiver-Client v1.0.1
	github.com/jadevelopmentgrp/Tickets-Database v1.0.8
	github.com/jadevelopmentgrp/Tickets-Utilities v1.0.2
	github.com/jadevelopmentgrp/Tickets-Worker v1.5.1
	github.com/joho/godotenv v1.5.1
	github.com/minio/minio-go/v7 v7.0.85
	github.com/penglongli/gin-metrics v0.1.10
	github.com/rxdn/gdl v0.0.0-20241201120412-8fd61c53dd96
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.9.0
	github.com/weppos/publicsuffix-go v0.20.0
	go.uber.org/zap v1.24.0
	golang.org/x/sync v0.11.0
)

require (
	github.com/ClickHouse/ch-go v0.52.1 // indirect
	github.com/ClickHouse/clickhouse-go/v2 v2.10.0 // indirect
	github.com/TicketsBot/ttlcache v1.6.1-0.20200405150101-acc18e37b261 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.2.0 // indirect
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/caarlos0/env v3.5.0+incompatible // indirect
	github.com/caarlos0/env/v10 v10.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-faster/city v1.0.1 // indirect
	github.com/go-faster/errors v0.6.1 // indirect
	github.com/go-ini/ini v1.67.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/goccy/go-json v0.10.4 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx v3.6.2+incompatible // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jadevelopmentgrp/Tickets-Analytics v1.0.0 // indirect
	github.com/jadevelopmentgrp/Tickets-Archiver v1.0.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/juju/ratelimit v1.0.2 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pasztorpisti/qs v0.0.0-20171216220353-8d6c33ee906c // indirect
	github.com/paulmach/orb v0.9.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.20.5 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/rs/xid v1.6.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/tatsuworks/czlib v0.0.0-20190916144400-8a51758ea0d9 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	nhooyr.io/websocket v1.8.17 // indirect
)
