# merchantlist

!!! info
	This page was last generated 2022.06.17

## Relationship Diagram

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY2hhbnRsaXN0IHtcbiAgICAgICAgaW50IGl0ZW1cbiAgICAgICAgaW50IG1lcmNoYW50aWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbWVyY2hhbnRsaXN0IHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY2hhbnRsaXN0IHtcbiAgICAgICAgaW50IGl0ZW1cbiAgICAgICAgaW50IG1lcmNoYW50aWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbWVyY2hhbnRsaXN0IHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbWVyY2hhbnRsaXN0IHtcbiAgICAgICAgaW50IGl0ZW1cbiAgICAgICAgaW50IG1lcmNoYW50aWRcbiAgICB9XG4gICAgaXRlbXMge1xuICAgICAgICBpbnQgaWRcbiAgICB9XG4gICAgbWVyY2hhbnRsaXN0IHx8LS1veyBpdGVtcyA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagram}

## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | item | [items](../../schema/items/items.md) | id |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| merchantid | int | Merchant Identifier |
| slot | int | Slot |
| item | int | [Item Identifier](items.md) |
| faction_required | smallint | Faction Required |
| level_required | tinyint | Level Required |
| alt_currency_cost | smallint | [Alternate Currency Cost](alternate_currency.md) |
| classes_required | int | [Classes Required](../../../../server/player/class-list) |
| probability | int | Probability: 0 = Never, 100 = Always |
| bucket_name | varchar | Bucket Name |
| bucket_value | varchar | Bucket Value |
| bucket_comparison | tinyint | [Bucket Comparison Type](../../../../server/scripting/merchant-data-buckets) |
| min_expansion | tinyint | [Minimum Expansion](../../../../server/operation/expansion-list) |
| max_expansion | tinyint | [Maximum Expansion](../../../../server/operation/expansion-list) |
| content_flags | varchar | Content Flags Required to be Enabled |
| content_flags_disabled | varchar | Content Flags Required to be Disabled |

