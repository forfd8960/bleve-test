## Index Flow

### 1. New Index

```go
mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		panic(err)
	}
```

### 2. Index Data

```go
if err := index.Index(msg.Id, msg); err != nil {
			fmt.Printf("index err: %v\n", err)
			return
		}
```

### 2.1 indexImpl.Index

path: `github.com/blevesearch/bleve.(*indexImpl).Index()`

```go
func (i *indexImpl) Index(id string, data interface{}) (err error) {}
```

### 2.2 Map Doc MapDocument

```go
	doc := document.NewDocument(id)
   263:		err = i.m.MapDocument(doc, data)
   264:		if err != nil {
   265:			return
   266:		}
```

### 2.2.1 Map Doc

path: `github.com/blevesearch/bleve/mapping.(*IndexMappingImpl).MapDocument()`

```go
func (im *IndexMappingImpl) MapDocument(doc *document.Document, data interface{}) error {}
```

### 2.2.1.1 im.determineType

path: ``

```go
// _default
docType := im.determineType(data)
```

### 2.2.1.2 im.mappingForType

```go
/*
p docMapping
("*github.com/blevesearch/bleve/mapping.DocumentMapping")(0x14000118190)
*github.com/blevesearch/bleve/mapping.DocumentMapping {
	Enabled: true,
	Dynamic: true,
	Properties: map[string]*github.com/blevesearch/bleve/mapping.DocumentMapping nil,
	Fields: []*github.com/blevesearch/bleve/mapping.FieldMapping len: 0, cap: 0, nil,
	DefaultAnalyzer: "",
	StructTagKey: "",}
*/
docMapping := im.mappingForType(docType)
``
```

### 2.2.1.3 im.newWalkContext

path: `github.com/blevesearch/bleve/mapping.(*IndexMappingImpl).MapDocument()`

```go
walkContext := im.newWalkContext(doc, docMapping)
```

### 2.2.1.4 docMapping.walkDocument

path: ``

```go
/*
p doc
("*github.com/blevesearch/bleve/document.Document")(0x14000114740)
*github.com/blevesearch/bleve/document.Document {
	ID: "index-example1",
	Fields: []github.com/blevesearch/bleve/document.Field len: 3, cap: 4, [
		...,
		...,
		...,
	],
	CompositeFields: []*github.com/blevesearch/bleve/document.CompositeField len: 1, cap: 1, [
		*(*"github.com/blevesearch/bleve/document.CompositeField")(0x140001147c0),
	],}


 p doc.Fields[0]
github.com/blevesearch/bleve/document.Field(*github.com/blevesearch/bleve/document.TextField) *{
	name: "Id",
	arrayPositions: []uint64 len: 0, cap: 0, [],
	options: DefaultCompositeIndexingOptions|StoreField|IncludeTermVectors|DocValues (15),
	analyzer: *github.com/blevesearch/bleve/analysis.Analyzer {
		CharFilters: []github.com/blevesearch/bleve/analysis.CharFilter len: 0, cap: 0, nil,
		Tokenizer: github.com/blevesearch/bleve/analysis.Tokenizer(*github.com/blevesearch/bleve/analysis/tokenizer/unicode.UnicodeTokenizer) ...,
		TokenFilters: []github.com/blevesearch/bleve/analysis.TokenFilter len: 2, cap: 2, [
			...,
			...,
		],},
	value: []uint8 len: 14, cap: 16, [105,110,100,101,120,45,101,120,97,109,112,108,101,49],
	numPlainTextBytes: 14,}

p doc.Fields[0].analyzer
("*github.com/blevesearch/bleve/analysis.Analyzer")(0x14000114380)
*github.com/blevesearch/bleve/analysis.Analyzer {
	CharFilters: []github.com/blevesearch/bleve/analysis.CharFilter len: 0, cap: 0, nil,
	Tokenizer: github.com/blevesearch/bleve/analysis.Tokenizer(*github.com/blevesearch/bleve/analysis/tokenizer/unicode.UnicodeTokenizer) *{},
	TokenFilters: []github.com/blevesearch/bleve/analysis.TokenFilter len: 2, cap: 2, [
		...,
		...,
	],}

p doc.Fields[0].analyzer.TokenFilters[0]
github.com/blevesearch/bleve/analysis.TokenFilter(*github.com/blevesearch/bleve/analysis/token/lowercase.LowerCaseFilter) *{}

p doc.Fields[0].analyzer.TokenFilters[1]
github.com/blevesearch/bleve/analysis.TokenFilter(*github.com/blevesearch/bleve/analysis/token/stop.StopTokensFilter) *{
        stopTokens: github.com/blevesearch/bleve/analysis.TokenMap [
                "they": true,
                "theirs": true,
                "that": true,
                "won't": true,
                "cannot": true,
                "for": true,
                "with": true,
                "after": true,
                "when": true,
                "more": true,
                ...


p doc.Fields[0].analyzer.Tokenizer
github.com/blevesearch/bleve/analysis.Tokenizer(*github.com/blevesearch/bleve/analysis/tokenizer/unicode.UnicodeTokenizer) *{}
*/


(dlv) p doc.Fields[1]
github.com/blevesearch/bleve/document.Field(*github.com/blevesearch/bleve/document.TextField) *{
	name: "From",
	arrayPositions: []uint64 len: 0, cap: 0, [],
	options: DefaultCompositeIndexingOptions|StoreField|IncludeTermVectors|DocValues (15),
	analyzer: *github.com/blevesearch/bleve/analysis.Analyzer {
		CharFilters: []github.com/blevesearch/bleve/analysis.CharFilter len: 0, cap: 0, nil,
		Tokenizer: github.com/blevesearch/bleve/analysis.Tokenizer(*github.com/blevesearch/bleve/analysis/tokenizer/unicode.UnicodeTokenizer) ...,
		TokenFilters: []github.com/blevesearch/bleve/analysis.TokenFilter len: 2, cap: 2, [
			...,
			...,
		],},
	value: []uint8 len: 20, cap: 24, [102,111,114,102,100,56,57,54,48,64,103,105,116,104,117,98,46,99,111,109],
	numPlainTextBytes: 20,}
(dlv) p string(doc.Fields[1].value)
"forfd8960@github.com"

```

```go
docMapping.walkDocument(data, []string{}, []uint64{}, walkContext)
```

### 2.2.1.5 add `_all` field to doc

```go
allMapping := docMapping.documentMappingForPath("_all")
if allMapping == nil || allMapping.Enabled {
   330:				field := document.NewCompositeFieldWithIndexingOptions("_all", true, []string{}, walkContext.excludedFromAll, document.IndexField|document.IncludeTermVectors)
=> 331:				doc.AddField(field)
   332:			}
``
```

### 2.3

path: `github.com/blevesearch/bleve/index/upsidedown.(*UpsideDownCouch).Update()`

```go
func (udc *UpsideDownCouch) Update(doc *document.Document) (err error) {}
```
