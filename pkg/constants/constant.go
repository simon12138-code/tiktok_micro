package constant

// 存放配置信息，可以通过Viper插件进行优化
const (
	NoteTableName           = "note"
	UserTableName           = "user"
	SecretKey               = "secret key"
	IdentityKey             = "id"
	JWTKey                  = "EYsnfKMf5XWk87LASEs28Dj5ZqGkSerH"
	Total                   = "total"
	Notes                   = "notes"
	NoteID                  = "note_id"
	JWTSigningKey           = "EYsnfKMf5XWk87LASEs28Dj5ZqGkSerH"
	ApiServiceName          = "api"
	NoteServiceName         = "note"
	UserServiceName         = "user"
	MySQLDefaultDSN         = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress             = "127.0.0.1:2379"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10
)
