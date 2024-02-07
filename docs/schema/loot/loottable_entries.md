# loottable_entries

!!! info
	This page was last generated 2024.02.07

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdHRhYmxlX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICB9XG4gICAgbG9vdHRhYmxlIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBsb290ZHJvcF9lbnRyaWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICAgICAgaW50IGl0ZW1faWRcbiAgICB9XG4gICAgbG9vdHRhYmxlX2VudHJpZXMgfHwtLW97IGxvb3R0YWJsZSA6IFwiT25lLXRvLU9uZVwiXG4gICAgbG9vdHRhYmxlX2VudHJpZXMgfHwtLW97IGxvb3Rkcm9wX2VudHJpZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdHRhYmxlX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICB9XG4gICAgbG9vdHRhYmxlIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBsb290ZHJvcF9lbnRyaWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICAgICAgaW50IGl0ZW1faWRcbiAgICB9XG4gICAgbG9vdHRhYmxlX2VudHJpZXMgfHwtLW97IGxvb3R0YWJsZSA6IFwiT25lLXRvLU9uZVwiXG4gICAgbG9vdHRhYmxlX2VudHJpZXMgfHwtLW97IGxvb3Rkcm9wX2VudHJpZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdHRhYmxlX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICB9XG4gICAgbG9vdHRhYmxlIHtcbiAgICAgICAgaW50dW5zaWduZWQgaWRcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBsb290ZHJvcF9lbnRyaWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICAgICAgaW50IGl0ZW1faWRcbiAgICB9XG4gICAgbG9vdHRhYmxlX2VudHJpZXMgfHwtLW97IGxvb3R0YWJsZSA6IFwiT25lLXRvLU9uZVwiXG4gICAgbG9vdHRhYmxlX2VudHJpZXMgfHwtLW97IGxvb3Rkcm9wX2VudHJpZXMgOiBcIkhhcy1NYW55XCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | loottable_id | [loottable](../../schema/loot/loottable.md) | id |
| Has-Many | lootdrop_id | [lootdrop_entries](../../schema/loot/lootdrop_entries.md) | lootdrop_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| loottable_id | int | [Loottable Identifier](loottable.md) |
| lootdrop_id | int | [Lootdrop Identifier](lootdrop.md) |
| multiplier | tinyint | Multiplier |
| droplimit | tinyint | Maximum Drops |
| mindrop | tinyint | Minimum Drops |
| probability | float | Probability: 0 = Never, 100 = Always |

