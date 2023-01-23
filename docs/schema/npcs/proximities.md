# proximities

!!! info
	This page was last generated 2023.01.22

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcHJveGltaXRpZXMge1xuICAgICAgICB2YXJjaGFyIHpvbmVpZFxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB2YXJjaGFyIHNob3J0X25hbWVcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgIH1cbiAgICBwcm94aW1pdGllcyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcHJveGltaXRpZXMge1xuICAgICAgICB2YXJjaGFyIHpvbmVpZFxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB2YXJjaGFyIHNob3J0X25hbWVcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgIH1cbiAgICBwcm94aW1pdGllcyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgcHJveGltaXRpZXMge1xuICAgICAgICB2YXJjaGFyIHpvbmVpZFxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB2YXJjaGFyIHNob3J0X25hbWVcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgIH1cbiAgICBwcm94aW1pdGllcyB8fC0tb3sgem9uZSA6IE9uZS10by1PbmVcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | zoneid | [zone](../../schema/zone/zone.md) | zoneidnumber |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| zoneid | int | [Zone Identifier](../../../../server/zones/zone-list) |
| exploreid | int | [Explore Identifier](../../schema/tasks/goallists.md) |
| minx | float | Minimum X Coordinate |
| maxx | float | Maximum X Coordinate |
| miny | float | Minimum Y Coordinate |
| maxy | float | Maximum Y Coordinate |
| minz | float | Minimum Z Coordinate |
| maxz | float | Maximum Z Coordinate |

