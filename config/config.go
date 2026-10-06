package config

import "errors"

// GetPort читает обязательный параметр PORT из карты конфигурации
func GetPort(cfg map[string]string) (string, error) {
	port, ok := cfg["PORT"]
	if !ok {
		return "", errors.New("ключ PORT отсутствует в конфигурации")
	}
	return port, nil
}

// validateKey — внутренняя проверка
func validateKey(key string) bool {
	return key != ""
}
