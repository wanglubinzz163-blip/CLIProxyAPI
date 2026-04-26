package config

import (
	"encoding/json"
	"strings"
)

// APIKeyEntry represents one client API key plus an optional remark/owner label.
type APIKeyEntry struct {
	Key    string `yaml:"key,omitempty" json:"key,omitempty"`
	Remark string `yaml:"remark,omitempty" json:"remark,omitempty"`
}

// APIKeyList supports both legacy string arrays and object arrays.
type APIKeyList []APIKeyEntry

type apiKeyEntryAlias APIKeyEntry

func normalizeAPIKeyEntry(entry APIKeyEntry) APIKeyEntry {
	entry.Key = strings.TrimSpace(entry.Key)
	entry.Remark = strings.TrimSpace(entry.Remark)
	return entry
}

func (l APIKeyList) Normalize() APIKeyList {
	if len(l) == 0 {
		return nil
	}
	out := make(APIKeyList, 0, len(l))
	seen := make(map[string]struct{}, len(l))
	for _, raw := range l {
		entry := normalizeAPIKeyEntry(raw)
		if entry.Key == "" {
			continue
		}
		if _, ok := seen[entry.Key]; ok {
			continue
		}
		seen[entry.Key] = struct{}{}
		out = append(out, entry)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func (l APIKeyList) Keys() []string {
	if len(l) == 0 {
		return nil
	}
	out := make([]string, 0, len(l))
	for _, entry := range l.Normalize() {
		if entry.Key != "" {
			out = append(out, entry.Key)
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func (l APIKeyList) MarshalYAML() (any, error) {
	normalized := l.Normalize()
	if len(normalized) == 0 {
		return []APIKeyEntry{}, nil
	}
	return []APIKeyEntry(normalized), nil
}

func (l *APIKeyList) UnmarshalYAML(unmarshal func(any) error) error {
	var arr []any
	if err := unmarshal(&arr); err != nil {
		var single string
		if err2 := unmarshal(&single); err2 == nil {
			*l = APIKeyList{{Key: strings.TrimSpace(single)}}.Normalize()
			return nil
		}
		return err
	}

	out := make(APIKeyList, 0, len(arr))
	for _, item := range arr {
		switch v := item.(type) {
		case string:
			out = append(out, APIKeyEntry{Key: v})
		case map[string]any:
			entry := APIKeyEntry{}
			for _, candidate := range []string{"key", "api-key", "apiKey", "Key"} {
				if raw, ok := v[candidate]; ok {
					if s, ok := raw.(string); ok {
						entry.Key = s
						break
					}
				}
			}
			if raw, ok := v["remark"]; ok {
				if s, ok := raw.(string); ok {
					entry.Remark = s
				}
			}
			if raw, ok := v["owner"]; ok && entry.Remark == "" {
				if s, ok := raw.(string); ok {
					entry.Remark = s
				}
			}
			out = append(out, entry)
		}
	}
	*l = out.Normalize()
	return nil
}

func (l APIKeyList) MarshalJSON() ([]byte, error) {
	return json.Marshal([]APIKeyEntry(l.Normalize()))
}

func (l *APIKeyList) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	if err := json.Unmarshal(data, &arr); err == nil {
		out := make(APIKeyList, 0, len(arr))
		for _, item := range arr {
			var s string
			if err := json.Unmarshal(item, &s); err == nil {
				out = append(out, APIKeyEntry{Key: s})
				continue
			}
			var entry APIKeyEntry
			if err := json.Unmarshal(item, &entry); err == nil {
				if entry.Key == "" {
					var record map[string]any
					if err2 := json.Unmarshal(item, &record); err2 == nil {
						for _, candidate := range []string{"api-key", "apiKey", "Key"} {
							if raw, ok := record[candidate]; ok {
								if s, ok := raw.(string); ok {
									entry.Key = s
									break
								}
							}
						}
						if entry.Remark == "" {
							if raw, ok := record["owner"]; ok {
								if s, ok := raw.(string); ok {
									entry.Remark = s
								}
							}
						}
					}
				}
				out = append(out, entry)
			}
		}
		*l = out.Normalize()
		return nil
	}

	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		*l = APIKeyList{{Key: single}}.Normalize()
		return nil
	}

	var aliasItems []apiKeyEntryAlias
	if err := json.Unmarshal(data, &aliasItems); err != nil {
		return err
	}
	out := make(APIKeyList, 0, len(aliasItems))
	for _, item := range aliasItems {
		out = append(out, APIKeyEntry(item))
	}
	*l = out.Normalize()
	return nil
}
