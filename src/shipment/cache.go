package shipment

var CacheDB map[string]string

// Initialize the CacheDB map
func InitCache() {
	CacheDB = make(map[string]string)
}

// Set a value in the CacheDB
func SetCache(key, value string) {
	CacheDB[key] = value
}

// Get a value from the CacheDB
func GetCache() map[string]string {
	return CacheDB
}
