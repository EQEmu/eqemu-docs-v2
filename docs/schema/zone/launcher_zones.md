# launcher_zones

!!! info
	This page was last generated 2023.01.30

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbGF1bmNoZXJfem9uZXMge1xuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICAgICAgdmFyY2hhciBsYXVuY2hlclxuICAgIH1cbiAgICBsYXVuY2hlciB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBsYXVuY2hlcl96b25lcyB8fC0tb3sgbGF1bmNoZXIgOiBcIk9uZS10by1PbmVcIlxuICAgIGxhdW5jaGVyX3pvbmVzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbGF1bmNoZXJfem9uZXMge1xuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICAgICAgdmFyY2hhciBsYXVuY2hlclxuICAgIH1cbiAgICBsYXVuY2hlciB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBsYXVuY2hlcl96b25lcyB8fC0tb3sgbGF1bmNoZXIgOiBcIk9uZS10by1PbmVcIlxuICAgIGxhdW5jaGVyX3pvbmVzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbGF1bmNoZXJfem9uZXMge1xuICAgICAgICB2YXJjaGFyIHpvbmVcbiAgICAgICAgdmFyY2hhciBsYXVuY2hlclxuICAgIH1cbiAgICBsYXVuY2hlciB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgaW50IHpvbmVpZG51bWJlclxuICAgICAgICB0aW55aW50dW5zaWduZWQgdmVyc2lvblxuICAgICAgICB2YXJjaGFyIGNvbnRlbnRfZmxhZ3NcbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzX2Rpc2FibGVkXG4gICAgICAgIHZhcmNoYXIgc2hvcnRfbmFtZVxuICAgIH1cbiAgICBsYXVuY2hlcl96b25lcyB8fC0tb3sgbGF1bmNoZXIgOiBcIk9uZS10by1PbmVcIlxuICAgIGxhdW5jaGVyX3pvbmVzIHx8LS1veyB6b25lIDogXCJPbmUtdG8tT25lXCJcblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjp0cnVlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6dHJ1ZX0=){target=diagrams}


## Relationships

| Relationship Type | Local Key | Relates to Table | Foreign Key |
| :--- | :--- | :--- | :--- |
| One-to-One | launcher | [launcher](../../schema/zone/launcher.md) | name |
| One-to-One | zone | [zone](../../schema/zone/zone.md) | short_name |


## Schema

| Column | Data Type | Description |
| :--- | :--- | :--- |
| launcher | varchar | Launcher |
| zone | varchar | [Zone Short Name](../../../../server/zones/zone-list) |
| port | mediumint | Port |

