package grammar

type Grammar struct {
	grammarMap map[string]string
}

func NewGrammar() Grammar {
	return Grammar{
		grammarMap: map[string]string{
			"present":            "現在形",
			"past":               "過去形",
			"past participle":    "過去分詞",
			"present participle": "現在分詞",
			"imperative":         "命令形",
			"infinitive":         "不定詞",
			"gerund":             "動名詞",
			"indicative":         "直説法",
			"subjunctive":        "仮定法",
			"passive":            "受動態",
			"active":             "能動態",
			"perfect":            "完了形",
			"progressive":        "進行形",
		},
	}
}

func List() []string {
	return []string{
		"present",
		"past",
		"past participle",
		"present participle",
		"imperative",
		"infinitive",
		"gerund",
		"indicative",
		"subjunctive",
		"passive",
		"active",
		"perfect",
		"progressive"}
}

func RemovedValue(grammars []string) []string {
	l := List()

	m := make(map[string]bool)
	for _, v := range grammars {
		m[v] = true
	}

	var res []string
	for _, v := range l {
		if !m[v] {
			res = append(res, v)
		}
	}

	return res
}
