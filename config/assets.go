package config

import (
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func BuildAssets() {
	browserEngines := []api.Engine{
		{Name: api.EngineChrome, Version: "58"},
		{Name: api.EngineFirefox, Version: "57"},
		{Name: api.EngineSafari, Version: "11"},
		{Name: api.EngineEdge, Version: "16"},
	}

	buildScripts := api.Build(api.BuildOptions{
		EntryPoints:       []string{"./app/assets/scripts/index.js"},
		Bundle:            true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Outfile:           "./public/assets/script.js",
		Sourcemap:         api.SourceMapLinked,
		Engines:           browserEngines,
		Write:             true,
	})

	if len(buildScripts.Errors) > 0 {
		os.Exit(1)
	}

	buildStyles := api.Build(api.BuildOptions{
		EntryPoints:       []string{"./app/assets/styles/index.css"},
		Bundle:            true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Outfile:           "./public/assets/style.css",
		Sourcemap:         api.SourceMapLinked,
		Engines:           browserEngines,
		Write:             true,
	})

	if len(buildStyles.Errors) > 0 {
		os.Exit(1)
	}
}
