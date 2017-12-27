package configuration

import "log"

var conf = "/opt/conf/configuration.properties"
var defaultConf = "C:/Users/itzik/Documents/work/test/src/main/resources/configuration.properties"

func GetPropertyValue(key string) string {
	if len(conf) == 0 {
		return getDefaultValue(key)
	}
	defer func() {
		recover()
	}()
	props, err := ReadPropertiesFile(conf)
	if err != nil {
		log.Panic("Error while reading properties file")
		return getDefaultValue(key)
	}
	if len(props[key]) == 0 {
		return getDefaultValue(key)
	}
	return props[key]
}

func getDefaultValue(key string) string {
	props, err := ReadPropertiesFile(defaultConf)
	if err != nil {
		log.Panic("Error while reading properties file")
		return ""
	}
	return props[key]
}
