package xconfig

// Config config
type Config struct {
	AppName   string `toml:"app_name"`
	AppSecret string `toml:"app_secret"`
	Web       struct {
		Enabled    bool   `toml:"enabled"`
		Name       string `toml:"name"`
		IsRecovery bool   `toml:"is_recovery"`
		IsSlash    bool   `toml:"is_slash"`
		SecretKey  string `toml:"secret_key"`
		EnableDocs bool   `toml:"enable_docs"`
		Gzip       struct {
			Enable bool `toml:"enable"`
			Level  int  `toml:"level"`
		} `toml:"gzip"`
		Auth struct {
			Enabled bool `toml:"enabled"`
			Email   bool `toml:"email"`
		} `toml:"auth"`
		Realtime struct {
			Enabled bool   `toml:"enabled"`
			Broker  string `toml:"broker"`
			Conn    string `toml:"conn"`
		} `toml:"realtime"`
		Pubsub struct {
			Enabled bool   `toml:"enabled"`
			Broker  string `toml:"broker"`
			Conn    string `toml:"conn"`
		} `toml:"pubsub"`
		FileStore struct {
			Enabled   bool   `toml:"enabled"`
			StoreType string `toml:"storeType"`
			Conn      string `toml:"conn"`
			Rules     []struct {
				Prefix string `toml:"prefix"`
				Create string `toml:"create"`
				Read   string `toml:"read"`
				Delete string `toml:"delete"`
			} `toml:"rules"`
		} `toml:"fileStore"`
		View struct {
			Enabled     bool          `toml:"enabled"`
			ViewPath    string        `toml:"view_path"`
			HTMLPattern string        `toml:"html_pattern"`
			HTMLFiles   []interface{} `toml:"html_files"`
			Delims      []string      `toml:"delims"`
		} `toml:"view"`
		Functions struct {
			Enabled bool   `toml:"enabled"`
			Broker  string `toml:"broker"`
			Conn    string `toml:"conn"`
			Rule    string `toml:"rule"`
		} `toml:"functions"`
		Static struct {
			Enabled          bool     `toml:"enabled"`
			Dir              []string `toml:"dir"`
			ExtensionsToGzip []string `toml:"ExtensionsToGzip"`
		} `toml:"static"`
		Logger struct {
			Enabled       bool     `toml:"enabled"`
			UTC           bool     `toml:"UTC"`
			SkipPath      []string `toml:"skip_path"`
			SkipPathRegex string   `toml:"skip_path_regex"`
		} `toml:"logger"`
		Forcessl struct {
			Enabled            bool   `toml:"enabled"`
			TrustXFPHeader     bool   `toml:"trustXFPHeader"`
			Enable301Redirects bool   `toml:"enable301Redirects"`
			Message            string `toml:"message"`
		} `toml:"forcessl"`
		XSRF struct {
			Enabled bool   `toml:"enabled"`
			Key     string `toml:"key"`
			Expire  int    `toml:"expire"`
		} `toml:"xsrf"`
		Upload struct {
			Enabled      bool     `toml:"enabled"`
			FileMaxSize  int      `toml:"file_max_size"`
			FileExt      []string `toml:"file_ext"`
			Prefix       string   `toml:"prefix"`
			StorageType  string   `toml:"storage_type"`
			OssName      string   `toml:"oss_name"`
			SaveLocation string   `toml:"save_location"`
		} `toml:"upload"`
		Server struct {
			Graceful   bool   `toml:"graceful"`
			Addr       string `toml:"addr"`
			UnixSocket string `toml:"unix_socket"`
		} `toml:"server"`
		Jwt struct {
			Enabled          bool   `toml:"enabled"`
			Realm            string `toml:"realm"`
			SigningAlgorithm string `toml:"signing_algorithm"`
			PubKeyFile       string `toml:"pub_key_file"`
			PrivKeyFile      string `toml:"priv_key_file"`
			Key              string `toml:"key"`
			Timeout          string `toml:"timeout"`
			MaxRefresh       string `toml:"max_refresh"`
			IdentityKey      string `toml:"identity_key"`
			TokenLookup      string `toml:"token_lookup"`
			TokenHeadName    string `toml:"token_head_name"`
		} `toml:"jwt"`
		Cors struct {
			Enabled           bool     `toml:"enabled"`
			AllowAllOrigins   bool     `toml:"allow_all_origins"`
			AllowOrigins      []string `toml:"allow_origins"`
			AllowRegexOrigins []string `toml:"allow_regex_origins"`
			AllowMethods      []string `toml:"allow_methods"`
			AllowHeaders      []string `toml:"allow_headers"`
			ExposeHeaders     []string `toml:"expose_headers"`
			AllowCredentials  bool     `toml:"allow_credentials"`
			AllowWildcard     bool     `toml:"allow_wildcard"`
			AllowWebSockets   bool     `toml:"allow_web_sockets"`
			AllowFiles        bool     `toml:"allow_files"`
			MaxAge            int      `toml:"max_age"`
		} `toml:"cors"`
		Cookie struct {
			Enabled        bool   `toml:"enabled"`
			FlashName      string `toml:"flash_name"`
			FlashSeparator string `toml:"flash_separator"`
			AutoSetCookie  bool   `toml:"auto_set_cookie"`
			Expire         int    `toml:"expire"`
		} `toml:"cookie"`
		Db struct {
			Driver string `toml:"driver"`
			Name   string `toml:"name"`
		} `toml:"db"`
		Session struct {
			Enabled    bool     `toml:"enabled"`
			Name       string   `toml:"name"`
			Driver     string   `toml:"driver"`
			DriverName string   `toml:"driver_name"`
			Expire     int      `toml:"expire"`
			KeyPairs   []string `toml:"key_pairs"`
			KeyPrefix  string   `toml:"key_prefix"`
			Path       string   `toml:"path"`
			MaxAge     int      `toml:"max_age"`
			Secure     bool     `toml:"secure"`
			HTTPOnly   bool     `toml:"http_only"`
			SameSite   int      `toml:"same_site"`
			Domain     string   `toml:"domain"`
		} `toml:"session"`
		Admin struct {
			Enabled bool   `toml:"enabled"`
			Addr    string `toml:"addr"`
			Port    int    `toml:"port"`
			User    string `toml:"user"`
			Pass    string `toml:"pass"`
			Role    string `toml:"role"`
			Secret  string `toml:"secret"`
		} `toml:"admin"`
	} `toml:"web"`
	HTTP struct {
		Enabled bool   `toml:"enabled"`
		Charset string `toml:"charset"`
		Client  struct {
		} `toml:"client"`
	} `toml:"http"`
	Mongodb struct {
		Enabled bool   `toml:"enabled"`
		Default string `toml:"default"`
		Cfg     []struct {
			Name                    string `toml:"name"`
			Database                string `toml:"database"`
			URL                     string `toml:"url"`
			Host                    string `toml:"host"`
			Port                    int    `toml:"port"`
			Db                      string `toml:"db"`
			Username                string `toml:"username"`
			Password                string `toml:"password"`
			RepositoriesEnabled     bool   `toml:"repositories_enabled"`
			AuthMechanism           string `toml:"auth_mechanism"`
			AuthMechanismProperties struct {
			} `toml:"auth_mechanism_properties"`
			AuthSource             string   `toml:"auth_source"`
			PasswordSet            bool     `toml:"password_set"`
			AppName                string   `toml:"app_name"`
			ConnectTimeout         int      `toml:"connect_timeout"`
			Compressors            []string `toml:"compressors"`
			HeartbeatInterval      int      `toml:"heartbeat_interval"`
			Hosts                  []string `toml:"hosts"`
			LocalThreshold         int      `toml:"local_threshold"`
			MaxConnIdleTime        int      `toml:"max_conn_idle_time"`
			MaxPoolSize            int      `toml:"max_pool_size"`
			ReplicaSet             int      `toml:"replica_set"`
			RetryWrites            bool     `toml:"retry_writes"`
			RetryReads             bool     `toml:"retry_reads"`
			ServerSelectionTimeout int      `toml:"server_selection_timeout"`
			Direct                 bool     `toml:"direct"`
			SocketTimeout          int      `toml:"socket_timeout"`
			ZlibLevel              int      `toml:"zlib_level"`
			AuthenticateToAnything bool     `toml:"authenticate_to_anything"`
		} `toml:"cfg"`
	} `toml:"mongodb"`
	Rds struct {
		Enabled bool   `toml:"enabled"`
		Default string `toml:"default"`
		Prefix  string `toml:"prefix"`
		Cfg     []struct {
			Name              string `toml:"name"`
			URL               string `toml:"url"`
			Driver            string `toml:"driver"`
			User              string `toml:"user"`
			Pass              string `toml:"pass"`
			Db                string `toml:"db"`
			Initialize        bool   `toml:"initialize"`
			Schema            string `toml:"schema"`
			Data              string `toml:"data"`
			SQLScriptEncoding string `toml:"sql_script_encoding"`
			Platform          string `toml:"platform"`
			Separator         string `toml:"separator"`
			Username          string `toml:"username"`
			Password          string `toml:"password"`
			MaxActive         int    `toml:"max_active"`
			MaxIdle           int    `toml:"max_idle"`
			MaxOpen           int    `toml:"max_open"`
			MinIdle           int    `toml:"min_idle"`
			MaxLfetime        int    `toml:"max_lfetime"`
			InitialSize       int    `toml:"initial_size"`
			ValidationQuery   string `toml:"validation_query"`
			MaxWait           string `toml:"max_wait"`
		} `toml:"cfg"`
	} `toml:"rds"`
	Redis struct {
		Enabled bool   `toml:"enabled"`
		Default string `toml:"default"`
		Cfg     []struct {
			Name               string `toml:"name"`
			URL                string `toml:"url"`
			User               string `toml:"user"`
			Db                 int    `toml:"db"`
			Network            string `toml:"network"`
			Addr               string `toml:"addr"`
			Password           string `toml:"password"`
			MaxRetries         int    `toml:"max_retries"`
			MaxRetry           int    `toml:"max_retry"`
			DialTimeout        int    `toml:"dial_timeout"`
			ReadTimeout        int    `toml:"read_timeout"`
			WriteTimeout       int    `toml:"write_timeout"`
			MaxConnAge         int    `toml:"max_conn_age"`
			PoolSize           int    `toml:"pool_size"`
			PoolTimeout        int    `toml:"pool_timeout"`
			IdleTimeout        int    `toml:"idle_timeout"`
			IdleCheckFrequency int    `toml:"idle_check_frequency"`
		} `toml:"cfg"`
	} `toml:"redis"`
	Search struct {
		Enabled bool `toml:"enabled"`
		Es      struct {
			Enabled bool   `toml:"enabled"`
			Default string `toml:"default"`
			Cfg     []struct {
				Name string `toml:"name"`
			} `toml:"cfg"`
		} `toml:"es"`
	} `toml:"search"`
	Mq struct {
		Enabled  bool `toml:"enabled"`
		RabbitMQ struct {
			Enabled             bool   `toml:"enabled"`
			Default             string `toml:"default"`
			DefaultExchange     string `toml:"default_exchange"`
			PublishRetryTime    int    `toml:"publish_retry_time"`
			DefaultPriority     string `toml:"default_priority"`
			DefaultXMaxPriority int    `toml:"default_x_max_priority"`
			DefaultXQueueMode   string `toml:"default_x_queue_mode"`
			DefaultRoutingKey   string `toml:"default_routing_key"`
			Cfg                 []struct {
				Name        string `toml:"name"`
				URL         string `toml:"url"`
				Username    string `toml:"username"`
				Password    string `toml:"password"`
				Port        int    `toml:"port"`
				Host        string `toml:"host"`
				VirtualHost string `toml:"virtual_host"`
				Dynamic     string `toml:"dynamic"`
				Channel     []struct {
					RoutingKey    string `toml:"routing_key"`
					ExchangeName  string `toml:"exchange_name"`
					ExchangeType  string `toml:"exchange_type"`
					QueueName     string `toml:"queue_name"`
					Consumer      string `toml:"consumer"`
					XMaxPriority  int    `toml:"x-max-priority"`
					XQueueMode    string `toml:"x-queue-mode"`
					Durable       bool   `toml:"durable"`
					AutoDelete    bool   `toml:"auto_delete"`
					Exclusive     bool   `toml:"exclusive"`
					NoWait        bool   `toml:"no_wait"`
					NoLocal       bool   `toml:"no_local"`
					AutoAck       bool   `toml:"auto_ack"`
					PrefetchCount int    `toml:"prefetch_count"`
					PrefetchSize  int    `toml:"prefetch_size"`
					Global        bool   `toml:"global"`
				} `toml:"channel"`
			} `toml:"cfg"`
		} `toml:"rabbitMQ"`
	} `toml:"mq"`
	Storage struct {
		Enabled bool `toml:"enabled"`
		Oss     struct {
			Enabled         bool   `toml:"enabled"`
			Default         string `toml:"default"`
			Endpoint        string `toml:"endpoint"`
			OutEndpoint     string `toml:"out_endpoint"`
			AccessKeyID     string `toml:"access_key_id"`
			AccessKeySecret string `toml:"access_key_secret"`
			RetryTimes      int    `toml:"retry_times"`
			UserAgent       string `toml:"user_agent"`
			IsDebug         bool   `toml:"is_debug"`
			Timeout         int    `toml:"timeout"`
			SecurityToken   string `toml:"security_token"`
			HTTPTimeout     int    `toml:"http_timeout"`
			HTTPMaxConns    int    `toml:"http_maxConns"`
			IsUseProxy      bool   `toml:"is_use_proxy"`
			ProxyHost       string `toml:"proxy_host"`
			IsAuthProxy     bool   `toml:"is_auth_proxy"`
			ProxyUser       string `toml:"proxy_user"`
			ProxyPassword   string `toml:"proxy_password"`
			IsEnableMd5     bool   `toml:"is_enable_md5"`
			Cfg             []struct {
				Name            string `toml:"name"`
				URL             string `toml:"url"`
				SaveOssPath     string `toml:"save_oss_path"`
				Bucket          string `toml:"bucket"`
				OssPath         string `toml:"oss_path"`
				Endpoint        string `toml:"endpoint"`
				OutEndpoint     string `toml:"out_endpoint"`
				AccessKeyID     string `toml:"access_key_id"`
				AccessKeySecret string `toml:"access_key_secret"`
				RetryTimes      int    `toml:"retry_times"`
				UserAgent       string `toml:"user_agent"`
				IsDebug         bool   `toml:"is_debug"`
				Timeout         int    `toml:"timeout"`
				SecurityToken   string `toml:"security_token"`
				HTTPTimeout     int    `toml:"http_timeout"`
				HTTPMaxConns    int    `toml:"http_maxConns"`
				IsUseProxy      bool   `toml:"is_use_proxy"`
				ProxyHost       string `toml:"proxy_host"`
				IsAuthProxy     bool   `toml:"is_auth_proxy"`
				ProxyUser       string `toml:"proxy_user"`
				ProxyPassword   string `toml:"proxy_password"`
				IsEnableMd5     bool   `toml:"is_enable_md5"`
			} `toml:"cfg"`
		} `toml:"oss"`
		File struct {
			Enabled bool   `toml:"enabled"`
			Default string `toml:"default"`
			Cfg     []struct {
				Name string `toml:"name"`
			} `toml:"cfg"`
		} `toml:"file"`
	} `toml:"storage"`
	Cache struct {
		Enabled bool   `toml:"enabled"`
		Default string `toml:"default"`
		Cfg     []struct {
			Driver    string `toml:"driver"`
			RedisName string `toml:"redis_name"`
		} `toml:"cfg"`
	} `toml:"cache"`
	SendCloud struct {
		Enabled bool   `toml:"enabled"`
		Default string `toml:"default"`
		Cfg     []struct {
			APIUser  string `toml:"api_user"`
			APIKey   string `toml:"api_key"`
			Encoding string `toml:"encoding"`
			Email    []struct {
				APIUser  string `toml:"api_user"`
				APIKey   string `toml:"api_key"`
				Encoding string `toml:"encoding"`
			} `toml:"email"`
			Sms []struct {
			} `toml:"sms"`
		} `toml:"cfg"`
		Email struct {
			Default string `toml:"default"`
			Enabled bool   `toml:"enabled"`
			Cfg     []struct {
				Name     string `toml:"name"`
				URL      string `toml:"url"`
				APIUser  string `toml:"api_user"`
				APIKey   string `toml:"api_key"`
				From     string `toml:"from"`
				FromName string `toml:"from_name"`
				Encoding string `toml:"encoding"`
			} `toml:"cfg"`
		} `toml:"email"`
		Sms struct {
			Cfg []struct {
				Name       string `toml:"name"`
				URL        string `toml:"url"`
				APIUser    string `toml:"api_user"`
				APIKey     string `toml:"api_key"`
				MsgType    string `toml:"msg_type"`
				TemplateID string `toml:"template_id"`
				From       string `toml:"from"`
				FromName   string `toml:"from_name"`
			} `toml:"cfg"`
		} `toml:"sms"`
	} `toml:"send_cloud"`
	Email struct {
		Enabled bool   `toml:"enabled"`
		Default string `toml:"default"`
		Cfg     []struct {
			Name       string `toml:"name"`
			Driver     string `toml:"driver"`
			DriverName string `toml:"driver_name"`
		} `toml:"cfg"`
	} `toml:"email"`
	Sms struct {
		Enabled bool   `toml:"enabled"`
		Default string `toml:"default"`
		Cfg     []struct {
			Name       string `toml:"name"`
			Driver     string `toml:"driver"`
			DriverName string `toml:"driver_name"`
		} `toml:"cfg"`
	} `toml:"sms"`
	Social struct {
		Enabled  bool `toml:"enabled"`
		Facebook struct {
			Enabled bool   `toml:"enabled"`
			Default string `toml:"default"`
			Cfg     []struct {
				AppID     string `toml:"app_id"`
				AppSecret string `toml:"app_secret"`
			} `toml:"cfg"`
		} `toml:"facebook"`
		Linkedin struct {
			Enabled bool   `toml:"enabled"`
			Default string `toml:"default"`
			Cfg     []struct {
				AppID     string `toml:"app_id"`
				AppSecret string `toml:"app_secret"`
			} `toml:"cfg"`
		} `toml:"linkedin"`
		Twitter struct {
			Enabled bool   `toml:"enabled"`
			Default string `toml:"default"`
			Cfg     []struct {
				AppID     string `toml:"app_id"`
				AppSecret string `toml:"app_secret"`
			} `toml:"cfg"`
		} `toml:"twitter"`
	} `toml:"social"`
	Services struct {
		Default string `toml:"default"`
		Cfg     []struct {
			Captcha string `toml:"captcha"`
		} `toml:"cfg"`
	} `toml:"services"`
	Log struct {
		Enabled             bool   `toml:"enabled"`
		TimeFormat          string `toml:"time_format"`
		LogLevel            string `toml:"log_level"`
		TimestampFieldName  string `toml:"timestamp_field_name"`
		LevelFieldName      string `toml:"level_field_name"`
		MessageFieldName    string `toml:"message_field_name"`
		ErrorFieldName      string `toml:"error_field_name"`
		CallerFieldName     string `toml:"caller_field_name"`
		ErrorStackFieldName string `toml:"error_stack_field_name"`
		OutputType          string `toml:"output_type"`
		Filename            string `toml:"filename"`
		MaxSize             int    `toml:"max_size"`
		Rotate              bool   `toml:"rotate"`
	} `toml:"log"`
	Pay struct {
		Enabled bool `toml:"enabled"`
		AliPay  struct {
			Enabled bool   `toml:"enabled"`
			Default string `toml:"default"`
			Cfg     []struct {
				Name       string `toml:"name"`
				PublicKey  string `toml:"public_key"`
				PrivateKey string `toml:"private_key"`
				AppID      string `toml:"app_id"`
				GatewayURL string `toml:"gateway_url"`
				SignType   string `toml:"sign_type"`
			} `toml:"cfg"`
		} `toml:"aliPay"`
	} `toml:"pay"`
	Memcached struct {
		Enabled bool   `toml:"enabled"`
		Default string `toml:"default"`
		Cfg     []struct {
			Name         string   `toml:"name"`
			Servers      []string `toml:"servers"`
			Timeout      int      `toml:"timeout"`
			MaxIdleConns int      `toml:"max_idle_conns"`
		} `toml:"cfg"`
	} `toml:"memcached"`
	Grpc struct {
		Enabled bool   `toml:"enabled"`
		Default string `toml:"default"`
		Cfg     []struct {
			Port int    `toml:"port"`
			Name string `toml:"name"`
		} `toml:"cfg"`
	} `toml:"grpc"`
	Ext struct {
	} `toml:"ext"`
}
