/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package main

import (
    "flag"
    "github.com/xfali/gobatis-cmd/io"
    "log"
    "os"
    "strings"
)

type config struct {
    path        string
    packageName string
    modelFile   string
    tagName     string
}

func main() {
    driver := flag.String("driver", "mysql", "driver of db")
    packageName := flag.String("pkg", "xfali/gobatis/default", "Set the package name of .go file")
    dbName := flag.String("db", "", "the name of db instance used in model files")
    tableName := flag.String("table", "", "the name of table to be generated")
    host := flag.String("host", "localhost", "host of db")
    port := flag.Int("port", 3306, "port of db ")
    username := flag.String("user", "", "user name of db")
    pw := flag.String("pw", "", "password of db")
    path := flag.String("path", "", "root path to save files")
    modelfile := flag.String("model", "", "the name of model file")
    tagName := flag.String("tag", "xfield", "the name of field tag")
    flag.Parse()

    db, err := connect(*driver, *username, *pw, *host, *port)
    if err != nil {
        log.Print(err)
        os.Exit(-1)
    }
    defer close(db)

    root := formatPath(*path)

    config := config{
        path:        root,
        packageName: *packageName,
        modelFile:   *modelfile,
        tagName:     *tagName,
    }

    if *tableName == "" {
        err2 := generateAll(config, db, *dbName)
        if err2 != nil {
            log.Print(err2)
            os.Exit(-2)
        }
    } else {
        err2 := generate(config, db, *dbName, *tableName)
        if err2 != nil {
            log.Print(err2)
            os.Exit(-2)
        }
    }
    os.Exit(0)
}

func formatPath(path string) string {
    root := strings.TrimSpace(path)
    if root == "" {
        root = "./"
    } else {
        if !io.IsPathExists(path) {
            io.Mkdir(path)
        }
        if root[len(root)-1:] != "/" {
            root = root + "/"
        }
    }
    return root
}