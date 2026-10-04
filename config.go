package main

import "errors"

func GetPortFromConfig(config map[string]string) (string, error) {
	port, ok := config["PORT"]
	if !ok {
		return "", errors.New("ключ PORT отсутствует в конфигурации")
	}
	return port, nil
}
