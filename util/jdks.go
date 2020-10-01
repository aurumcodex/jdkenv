package util

type Oracle struct {
	Base     string `toml:base_url`
	Ext      string `toml:ext_url`
	Filename string `toml:filename`
}

type Corretto struct {
	Base     string `toml:base_url`
	Filename string `toml:filename`
	File     string `toml:file`
}
