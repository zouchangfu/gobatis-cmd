/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package main

var sqlType2GoMap = map[string]string{
    "int":                "int",
    "integer":            "int",
    "tinyint":            "int",
    "smallint":           "int",
    "mediumint":          "int",
    "bigint":             "int64",
    "int unsigned":       "uint",
    "integer unsigned":   "uint",
    "tinyint unsigned":   "uint",
    "smallint unsigned":  "uint",
    "mediumint unsigned": "uint",
    "bigint unsigned":    "uint64",
    "bit":                "int",
    "bool":               "bool",
    "enum":               "string",
    "set":                "string",
    "varchar":            "string",
    "char":               "string",
    "tinytext":           "string",
    "mediumtext":         "string",
    "text":               "string",
    "longtext":           "string",
    "blob":               "string",
    "tinyblob":           "string",
    "mediumblob":         "string",
    "longblob":           "string",
    "date":               "time.Time",
    "datetime":           "time.Time",
    "timestamp":          "time.Time",
    "time":               "time.Time",
    "float":              "float64",
    "double":             "float64",
    "decimal":            "float64",
    "binary":             "string",
    "varbinary":          "string",
}

var sqlType2IfFormatMap = map[string]string{
    "int":                "%s != nil and %s != 0",
    "integer":            "%s != nil and %s != 0",
    "tinyint":            "%s != nil and %s != 0",
    "smallint":           "%s != nil and %s != 0",
    "mediumint":          "%s != nil and %s != 0",
    "bigint":             "%s != nil and %s != 0",
    "int unsigned":       "%s != nil and %s != 0",
    "integer unsigned":   "%s != nil and %s != 0",
    "tinyint unsigned":   "%s != nil and %s != 0",
    "smallint unsigned":  "%s != nil and %s != 0",
    "mediumint unsigned": "%s != nil and %s != 0",
    "bigint unsigned":    "%s != nil and %s != 0",
    "bit":                "%s != nil and %s != 0",
    "bool":               "%s != nil and %s != false",
    "enum":               "%s != nil",
    "set":                "%s != nil",
    "varchar":            "%s != nil",
    "char":               "%s != nil",
    "tinytext":           "%s != nil",
    "mediumtext":         "%s != nil",
    "text":               "%s != nil",
    "longtext":           "%s != nil",
    "blob":               "%s != nil",
    "tinyblob":           "%s != nil",
    "mediumblob":         "%s != nil",
    "longblob":           "%s != nil",
    "date":               "%s != nil",
    "datetime":           "%s != nil",
    "timestamp":          "%s != nil",
    "time":               "%s != nil",
    "float":              "%s != nil and %s != 0",
    "double":             "%s != nil and %s != 0",
    "decimal":            "%s != nil and %s != 0",
    "binary":             "%s != nil",
    "varbinary":          "%s != nil",
}