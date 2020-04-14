package main

func entityParser(text string) string {
	if len(text) < 0 {
		return ""
	}
	if len(text) > 100000 {
		return ""
	}

	//text = strings.Replace(text, "&quot;", "\"", -1)
	//text = strings.Replace(text, "&apos;", "'", -1)
	//text = strings.Replace(text, "&gt;", ">", -1)
	//text = strings.Replace(text, "&amp;", "&", -1)
	//text = strings.Replace(text, "&lt;", "<", -1)
	//text = strings.Replace(text, "&frasl;", "/", -1)

	//replacer := strings.NewReplacer("&quot;", "\"", "&apos;", "'", "&amp;", "&", "&gt;", ">", "&lt;", "<", "&frasl;", "/")
	replacer := NewReplacer("&quot;", "\"", "&apos;", "'", "&amp;", "&", "&gt;", ">", "&lt;", "<", "&frasl;", "/")
	return replacer.Replace(text)
}
