package styles

// Merge combines multiple styles.Props maps into one, with later styles overriding earlier ones.
func Merge(styleMaps ...Props) Props {
	mergedStyles := Props{}
	for _, styleMap := range styleMaps {
		for key, value := range styleMap {
			mergedStyles[key] = value
		}
	}
	return mergedStyles
}
