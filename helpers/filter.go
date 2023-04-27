package helpers

import "net/url"

func ConvertToQueryString(filters map[string]interface{}) string {
	values := url.Values{}

	for key, value := range filters {
		if s, ok := value.(string); ok {
			values.Add(key, s)
		}
	}

	return values.Encode()
}
