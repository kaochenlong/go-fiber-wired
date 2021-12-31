package controllers

import "github.com/joho/godotenv"

var ENV map[string]string

func init() {
	ENV, _ = godotenv.Read()
}
