/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description:
 */

package test

import (
	"github.com/xfali/gobatis-cmd/pkg/common"
	"github.com/xfali/gobatis-cmd/pkg/config"
	"github.com/xfali/gobatis-cmd/pkg/generator"
	"github.com/xfali/gobatis-cmd/pkg/plugin"
	"testing"
)

func createModeInfo() *[]common.ModelInfo {
	model := []common.ModelInfo{}
	info := common.ModelInfo{
		ColumnName: "id",
		DataType:   "bigint",
		Nullable:   "NO",
		Comment:    "id",
		Tag:        "id",
		ColumnKey:  "PRI",
	}
	model = append(model, info)

	info = common.ModelInfo{
		ColumnName: "username",
		DataType:   "varchar",
		Nullable:   "NO",
		Comment:    "username",
		Tag:        "username",
	}
	model = append(model, info)

	info = common.ModelInfo{
		ColumnName: "password",
		DataType:   "varchar",
		Nullable:   "NO",
		Comment:    "password",
		Tag:        "password",
	}
	model = append(model, info)

	info = common.ModelInfo{
		ColumnName: "update_time",
		DataType:   "timestamp",
		Nullable:   "YES",
		Comment:    "update_time",
		Tag:        "update_time",
	}
	model = append(model, info)
	return &model
}

func TestMode(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "",
			MapperFile:  "xml",
			ModelFile:   "model.go",
		}

		generator.GenModel(config, "test_table", *createModeInfo())
	})

	t.Run("1", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield",
			MapperFile:  "xml",
			ModelFile:   "model.go",
		}

		generator.GenModel(config, "test_table", *createModeInfo())
	})

	t.Run("2", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield,json",
			MapperFile:  "xml",
			ModelFile:   "model.go",
		}

		generator.GenModel(config, "test_table", *createModeInfo())
	})

	t.Run("2.1", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield,json,",
			MapperFile:  "xml",
			ModelFile:   "model.go",
		}

		generator.GenModel(config, "test_table", *createModeInfo())
	})

	t.Run("2.2", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xml,json,",
			MapperFile:  "xml",
			ModelFile:   "model.go",
		}

		generator.GenModel(config, "test_table", *createModeInfo())
	})

	t.Run("3", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield,json,xml",
			MapperFile:  "xml",
			ModelFile:   "model.go",
		}

		generator.GenModel(config, "test_table", *createModeInfo())
	})
}

func TestXml(t *testing.T) {
	t.Run("no keyword", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield",
			MapperFile:  "xml",
		}
		generator.GenXml(config, "test_table", *createModeInfo())
	})

	t.Run("mysql", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield",
			MapperFile:  "xml",
			Driver:      "mysql",
			Keyword:     true,
		}
		generator.GenXml(config, "test_table", *createModeInfo())
	})

	t.Run("postgres", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield",
			MapperFile:  "xml",
			Driver:      "postgres",
			Keyword:     true,
		}
		generator.GenXml(config, "test_table", *createModeInfo())
	})

	t.Run("sqlserver", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield",
			MapperFile:  "xml",
			Driver:      "mssql",
			Keyword:     true,
		}
		generator.GenXml(config, "test_table", *createModeInfo())
	})
}

func TestTemplate(t *testing.T) {
	t.Run("no keyword", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield",
			MapperFile:  "template",
		}
		generator.GenTemplate(config, "test_table", *createModeInfo())
	})

	t.Run("mysql", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield",
			MapperFile:  "template",
			Driver:      "mysql",
			Keyword:     true,
		}
		generator.GenTemplate(config, "test_table", *createModeInfo())
	})

	t.Run("postgres", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield",
			MapperFile:  "template",
			Driver:      "postgres",
			Keyword:     true,
		}
		generator.GenTemplate(config, "test_table", *createModeInfo())
	})

	t.Run("sqlserver", func(t *testing.T) {
		config := config.Config{
			PackageName: "test_package",
			Path:        "c:/tmp/",
			TagName:     "xfield",
			MapperFile:  "template",
			Driver:      "mssql",
			Keyword:     true,
		}
		generator.GenTemplate(config, "test_table", *createModeInfo())
	})
}

func TestProxy(t *testing.T) {
	config := config.Config{
		PackageName: "test_package",
		Path:        "c:/tmp/",
		TagName:     "xfield",
		MapperFile:  "xml",
	}

	generator.GenProxy(config, "test_table", *createModeInfo())
}

func TestV2Proxy(t *testing.T) {
	config := config.Config{
		PackageName: "test_package",
		Path:        "c:/tmp/",
		TagName:     "xfield",
		MapperFile:  "xml",
	}

	generator.GenV2Proxy(config, "test_table", *createModeInfo())
}

func TestAll1(t *testing.T) {
	config := config.Config{
		PackageName: "test_package",
		Path:        "c:/tmp/",
		TagName:     "xfield",
		MapperFile:  "xml",
		//ModelFile:   "model.go",
	}
	info := *createModeInfo()
	generator.GenModel(config, "test_table", info)
	generator.GenXml(config, "test_table", info)
	generator.GenProxy(config, "test_table", info)
}

func TestAll2(t *testing.T) {
	config := config.Config{
		PackageName: "test_package",
		Path:        "c:/tmp/",
		TagName:     "xfield",
		MapperFile:  "go",
		//ModelFile:   "model.go",
	}
	info := *createModeInfo()
	generator.GenModel(config, "test_table", info)
	generator.GenXml(config, "test_table", info)
	generator.GenProxy(config, "test_table", info)
}

func TestAll3(t *testing.T) {
	config := config.Config{
		PackageName: "test_package",
		Path:        "c:/tmp/",
		TagName:     "xfield",
		MapperFile:  "xml",
		//ModelFile:   "model.go",
		Plugin: "c:/tmp/webplugin.exe",
	}
	info := *createModeInfo()
	generator.GenModel(config, "TEST_TABLE", info)
	generator.GenXml(config, "TEST_TABLE", info)
	generator.GenV2Proxy(config, "TEST_TABLE", info)
	plugin.RunPlugin(config, "TEST_TABLE", info)
}

func TestPlugin(t *testing.T) {
	config := config.Config{
		PackageName: "test_package",
		Path:        "c:/tmp/",
		TagName:     "xfield",
		MapperFile:  "xml",
		Plugin:      "c:/tmp/webplugin.exe",
		//ModelFile:   "model.go",
	}
	info := *createModeInfo()
	plugin.RunPlugin(config, "TEST_TABLE", info)
}
