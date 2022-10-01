package models

type PackageJson struct {
	Name            string                 `json:"name"`
	Version         string                 `json:"version"`
	Private         bool                   `json:"private"`
	Homepage        string                 `json:"homepage"`
	Dependencies    map[string]string      `json:"dependencies"`
	EslintConfig    map[string]string      `json:"eslintConfig"`
	DevDependencies map[string]string      `json:"devDependencies"`
	Scripts         map[string]string      `json:"scripts"`
	Browserslist    map[string]interface{} `json:"browserslist"`
}
