# npc_types_tint

!!! info
	This page was last generated 2022.06.18

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3R5cGVzX3RpbnQge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBtZXJjX2FybW9yaW5mbyB7XG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG5wY190eXBlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IG5wY19mYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGFkdmVudHVyZV90ZW1wbGF0ZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBhbHRfY3VycmVuY3lfaWRcbiAgICAgICAgaW50dW5zaWduZWQgZW1vdGVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY2hhbnRfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbnBjX3NwZWxsc19pZFxuICAgICAgICBpbnR1bnNpZ25lZCBucGNfc3BlbGxzX2VmZmVjdHNfaWRcbiAgICAgICAgaW50dW5zaWduZWQgdHJhcF90ZW1wbGF0ZVxuICAgICAgICB0ZXh0IG5hbWVcbiAgICB9XG4gICAgbnBjX3R5cGVzX3RpbnQgfHwtLW97IG1lcmNfYXJtb3JpbmZvIDogSGFzLU1hbnlcbiAgICBucGNfdHlwZXNfdGludCB8fC0tb3sgbnBjX3R5cGVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3R5cGVzX3RpbnQge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBtZXJjX2FybW9yaW5mbyB7XG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG5wY190eXBlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IG5wY19mYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGFkdmVudHVyZV90ZW1wbGF0ZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBhbHRfY3VycmVuY3lfaWRcbiAgICAgICAgaW50dW5zaWduZWQgZW1vdGVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY2hhbnRfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbnBjX3NwZWxsc19pZFxuICAgICAgICBpbnR1bnNpZ25lZCBucGNfc3BlbGxzX2VmZmVjdHNfaWRcbiAgICAgICAgaW50dW5zaWduZWQgdHJhcF90ZW1wbGF0ZVxuICAgICAgICB0ZXh0IG5hbWVcbiAgICB9XG4gICAgbnBjX3R5cGVzX3RpbnQgfHwtLW97IG1lcmNfYXJtb3JpbmZvIDogSGFzLU1hbnlcbiAgICBucGNfdHlwZXNfdGludCB8fC0tb3sgbnBjX3R5cGVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbnBjX3R5cGVzX3RpbnQge1xuICAgICAgICBpbnR1bnNpZ25lZCBpZFxuICAgIH1cbiAgICBtZXJjX2FybW9yaW5mbyB7XG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBtZXJjX25wY190eXBlX2lkXG4gICAgfVxuICAgIG5wY190eXBlcyB7XG4gICAgICAgIGludHVuc2lnbmVkIGFybW9ydGludF9pZFxuICAgICAgICBpbnQgaWRcbiAgICAgICAgaW50IG5wY19mYWN0aW9uX2lkXG4gICAgICAgIGludHVuc2lnbmVkIGFkdmVudHVyZV90ZW1wbGF0ZV9pZFxuICAgICAgICBpbnR1bnNpZ25lZCBhbHRfY3VycmVuY3lfaWRcbiAgICAgICAgaW50dW5zaWduZWQgZW1vdGVpZFxuICAgICAgICBpbnR1bnNpZ25lZCBsb290dGFibGVfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbWVyY2hhbnRfaWRcbiAgICAgICAgaW50dW5zaWduZWQgbnBjX3NwZWxsc19pZFxuICAgICAgICBpbnR1bnNpZ25lZCBucGNfc3BlbGxzX2VmZmVjdHNfaWRcbiAgICAgICAgaW50dW5zaWduZWQgdHJhcF90ZW1wbGF0ZVxuICAgICAgICB0ZXh0IG5hbWVcbiAgICB9XG4gICAgbnBjX3R5cGVzX3RpbnQgfHwtLW97IG1lcmNfYXJtb3JpbmZvIDogSGFzLU1hbnlcbiAgICBucGNfdHlwZXNfdGludCB8fC0tb3sgbnBjX3R5cGVzIDogSGFzLU1hbnlcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| Has-Many | id | [merc_armorinfo](../../schema/mercenaries/merc_armorinfo.md) | armortint_id |
| Has-Many | id | [npc_types](../../schema/npcs/npc_types.md) | armortint_id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| id | int | Unique NPC Type Tint Identifier |
| tint_set_name | text | Tint Set Name |
| red1h | tinyint | Red for Helmet: 0 = None, 255 = Max |
| grn1h | tinyint | Green for Helmet: 0 = None, 255 = Max |
| blu1h | tinyint | Blue for Helmet: 0 = None, 255 = Max |
| red2c | tinyint | Red for Chest: 0 = None, 255 = Max |
| grn2c | tinyint | Green for Chest: 0 = None, 255 = Max |
| blu2c | tinyint | Blue for Chest: 0 = None, 255 = Max |
| red3a | tinyint | Red for Arms: 0 = None, 255 = Max |
| grn3a | tinyint | Green for Arms: 0 = None, 255 = Max |
| blu3a | tinyint | Blue for Arms: 0 = None, 255 = Max |
| red4b | tinyint | Red for Bracers: 0 = None, 255 = Max |
| grn4b | tinyint | Green for Bracers: 0 = None, 255 = Max |
| blu4b | tinyint | Blue for Bracers: 0 = None, 255 = Max |
| red5g | tinyint | Red for Hands: 0 = None, 255 = Max |
| grn5g | tinyint | Green for Hands: 0 = None, 255 = Max |
| blu5g | tinyint | Blue for Hands: 0 = None, 255 = Max |
| red6l | tinyint | Red for Legs: 0 = None, 255 = Max |
| grn6l | tinyint | Green for Legs: 0 = None, 255 = Max |
| blu6l | tinyint | Blue for Legs: 0 = None, 255 = Max |
| red7f | tinyint | Red for Feet: 0 = None, 255 = Max |
| grn7f | tinyint | Green for Feet: 0 = None, 255 = Max |
| blu7f | tinyint | Blue for Feet: 0 = None, 255 = Max |
| red8x | tinyint | Red for Unknown: 0 = None, 255 = Max |
| grn8x | tinyint | Green for Unknown: 0 = None, 255 = Max |
| blu8x | tinyint | Blue for Unknown: 0 = None, 255 = Max |
| red9x | tinyint | Red for Unknown: 0 = None, 255 = Max |
| grn9x | tinyint | Green for Unknown: 0 = None, 255 = Max |
| blu9x | tinyint | Blue for Unknown: 0 = None, 255 = Max |

