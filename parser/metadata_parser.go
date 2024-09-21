// metadata_parser.go
package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Markdownファイルには、通常コンテンツとともにメタデータを埋め込むことができます。このメタデータは、Front Matter と呼ばれる形式で、主にYAMLやTOMLのフォーマットで記述されます。このメタデータ部分をパースして、記事のタイトル、タグ、日付などを取得するのがこのコードの目的です。
// •	gopkg.in/yaml.v2: YAML形式のデータをGoの構造体にパースするために使われるライブラリです。Markdownファイルに含まれるメタデータがYAML形式の場合、このライブラリを使って簡単にメタデータをパースできます。
type ArticleMeta struct {
	Title string   `yaml:"title"`
	Tags  []string `yaml:"tags"`
	Date  string   `yaml:"date"`
}

// •	ArticleMeta: Markdownファイルのメタデータを格納する構造体。この記事では、タイトル (Title)、タグ (Tags)、日付 (Date) をフィールドとして定義しています。フィールドに付いているタグ（yaml:"title" など）は、YAML形式のキーとGoの構造体のフィールドを対応させるために使用します。
func parseFrontMatter(content []byte) (ArticleMeta, error) {
	var meta ArticleMeta
	err := yaml.Unmarshal(content, &meta)
	return meta, err
}

// parseFrontMatter: この関数は、YAML形式のメタデータを受け取り、それをGoの構造体にパースします。
//   - yaml.Unmarshal: YAMLデータをGoの構造体（ここでは ArticleMeta）に変換するための関数です。この関数は、YAMLデータをパースして、構造体に対応するフィールドに値を設定します。
//   - 引数 content: YAML形式のメタデータ（[]byte）です。
//
// 　　　　- 戻り値: パースされた ArticleMeta 構造体とエラーオブジェクトです。
func main() {
	frontMatter := []byte(`
    ---
    title: "サンプル記事"
    tags: ["Go", "API"]
    date: "2024-09-20"
    ---
    `)

	meta, err := parseFrontMatter(frontMatter)
	if err != nil {
		fmt.Println("Error parsing front matter:", err)
		return
	}
	fmt.Println("Parsed meta:", meta)
}
