package maps

// Find returns the value V associated to the key K in the given map. It returns
// the provided default value if the key is not in the map.
func Find[K comparable, V any](col map[K]V, key K, defaultValue V) V {
	v, ok := col[key]

	if !ok {
		return defaultValue
	}

	return v
}
