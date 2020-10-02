package util

type Oracle struct {
	Base string `toml:base_url`
	Data OracleInfo
}

type OracleInfo struct {
	Ext      string `toml:ext_url`
	Filename string `toml:filename`
}

type Corretto struct {
	Base string `toml:base_url`
	Data CorrettoInfo
}

type CorrettoInfo struct {
	Filename string `toml:filename`
	File     string `toml:file`
}

type Liberica struct {
	Base string `toml:base_url`
	Data LibericaInfo
}

type LibericaInfo struct {
	Ext      string `toml:ext_url`
	Filename string `toml:filename`
}

type OpenJDK struct {
}

type OpenJDKInfo struct {
}
