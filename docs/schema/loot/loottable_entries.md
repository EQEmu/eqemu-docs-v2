# loottable_entries

!!! info
	This page was last generated 2023.01.22

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdHRhYmxlX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICB9XG4gICAgbG9vdHRhYmxlIHtcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBsb290ZHJvcF9lbnRyaWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICAgICAgaW50IGl0ZW1faWRcbiAgICB9XG4gICAgbG9vdHRhYmxlX2VudHJpZXMgfHwtLW97IGxvb3R0YWJsZSA6IE9uZS10by1PbmVcbiAgICBsb290dGFibGVfZW50cmllcyB8fC0tb3sgbG9vdGRyb3BfZW50cmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdHRhYmxlX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICB9XG4gICAgbG9vdHRhYmxlIHtcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBsb290ZHJvcF9lbnRyaWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICAgICAgaW50IGl0ZW1faWRcbiAgICB9XG4gICAgbG9vdHRhYmxlX2VudHJpZXMgfHwtLW97IGxvb3R0YWJsZSA6IE9uZS10by1PbmVcbiAgICBsb290dGFibGVfZW50cmllcyB8fC0tb3sgbG9vdGRyb3BfZW50cmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbG9vdHRhYmxlX2VudHJpZXMge1xuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICB9XG4gICAgbG9vdHRhYmxlIHtcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBsb290ZHJvcF9lbnRyaWVzIHtcbiAgICAgICAgaW50dW5zaWduZWQgbG9vdGRyb3BfaWRcbiAgICAgICAgaW50IGl0ZW1faWRcbiAgICB9XG4gICAgbG9vdHRhYmxlX2VudHJpZXMgfHwtLW97IGxvb3R0YWJsZSA6IE9uZS10by1PbmVcbiAgICBsb290dGFibGVfZW50cmllcyB8fC0tb3sgbG9vdGRyb3BfZW50cmllcyA6IEhhcy1NYW55XG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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

