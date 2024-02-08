package auth

// Config contains configuration params for NATS authentication.
type Config struct {
	NKeyFile            string `json:"nkey_file" yaml:"nkey_file"`
	NKeyFileContent     string `json:"nkey_file_content" yaml:"nkey_file_content"`
	UserCredentialsFile string `json:"user_credentials_file" yaml:"user_credentials_file"`
	UserJWT             string `json:"user_jwt" yaml:"user_jwt"`
	UserNkeySeed        string `json:"user_nkey_seed" yaml:"user_nkey_seed"`
}

// New creates a new Config instance.
func New() Config {
	return Config{
		NKeyFile:            "",
		NKeyFileContent:     "",
		UserCredentialsFile: "",
		UserJWT:             "",
		UserNkeySeed:        "",
	}
}
