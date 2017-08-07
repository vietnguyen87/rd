package main

func BuildInnerQueriesForSearchProductNewV8(matchQuerySlice []interface{}, isSearchInsideCategory bool, keywordFormat string, isAccentMarks bool, firstTerm string) []interface{} {
	var field1 []string
	if isSearchInsideCategory {
		field1 = []string{"Sku_user"}
	} else {
		field1 = []string{"Category_name_suggest", "Sku_user"}
	}
	var field2 = []string{"Name"}
	var keyWordFinal string
	keyWordFinal = keywordFormat

	for i := 0; i < len(field1); i++ {
		if isAccentMarks == true{
			matchQuerySlice = append(matchQuerySlice, map[string]interface{}{
				"match": map[string]interface{}{
					field1[i]: keyWordFinal,
				},
			})
		}
		if isAccentMarks == false{
			if field1[i] == "Sku_user" {
				matchQuerySlice = append(matchQuerySlice, map[string]interface{}{
					"match": map[string]interface{}{
						field1[i]: keyWordFinal,
					},
				})
			} else {
				matchQuerySlice = append(matchQuerySlice, map[string]interface{}{
					"dis_max": map[string]interface{}{
						"queries": []interface{}{
							map[string]interface{}{
								"match": map[string]interface{}{
									field1[i]: keyWordFinal,
								},
							},
							map[string]interface{}{
								"match": map[string]interface{}{
									field1[i] + ".synonyms_v2": keyWordFinal,
								},
							},
						},
					},
				})
			}
		}
	}
	for i := 0; i < len(field2); i++ {
		if isAccentMarks == true{
			tmpQuery1 := map[string]interface{}{
				"dis_max": map[string]interface{}{
					"queries": []interface{}{
						map[string]interface{}{
							"match": map[string]interface{}{
								field2[i]: map[string]interface{}{
									"query": keyWordFinal,
									"boost": 6,
								},
							},
						},
						map[string]interface{}{
							"match": map[string]interface{}{
								field2[i] + ".with_accent_marks_v2": map[string]interface{}{
									"boost": 12,
									"query": keyWordFinal,
								},
							},
						},
					},
				},
			}
			matchQuerySlice = append(matchQuerySlice, map[string]interface{}{
				"bool": map[string]interface{}{
					"should": []interface{}{
						tmpQuery1,
						map[string]interface{}{
							"match_phrase": map[string]interface{}{
								field2[i] + ".with_accent_marks_v2": map[string]interface{}{
									"query": firstTerm,
									"boost": 2,
								},
							},
						},
					},
				},
			})
		} else {
			matchQuerySlice = append(matchQuerySlice, map[string]interface{}{
				"dis_max": map[string]interface{}{
					"queries": []interface{}{
						map[string]interface{}{
							"match": map[string]interface{}{
								field2[i]: map[string]interface{}{
									"query": keyWordFinal,
									"boost": 6,
								},
							},
						},
						map[string]interface{}{
							"match": map[string]interface{}{
								field2[i] + ".synonyms_v2": map[string]interface{}{
									"query": keyWordFinal,
									"boost": 6,
								},
							},
						},
					},
				},
			})
		}
	}
	return matchQuerySlice
}
