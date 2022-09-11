# launcher_zones

!!! info
	This page was last generated 2022.09.10

## Relationship Diagram(s)

[Diagram Edit](https://mermaid.live/edit#eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbGF1bmNoZXJfem9uZXMge1xuICAgICAgICB2YXJjaGFyIGxhdW5jaGVyXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICBsYXVuY2hlciB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBsYXVuY2hlcl96b25lcyB8fC0tb3sgbGF1bmNoZXIgOiBPbmUtdG8tT25lXG4gICAgbGF1bmNoZXJfem9uZXMgfHwtLW97IHpvbmUgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams-edit}

[![](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbGF1bmNoZXJfem9uZXMge1xuICAgICAgICB2YXJjaGFyIGxhdW5jaGVyXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICBsYXVuY2hlciB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBsYXVuY2hlcl96b25lcyB8fC0tb3sgbGF1bmNoZXIgOiBPbmUtdG8tT25lXG4gICAgbGF1bmNoZXJfem9uZXMgfHwtLW97IHpvbmUgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9)](https://mermaid.ink/img/eyJjb2RlIjoiZXJEaWFncmFtXG4gICAgbGF1bmNoZXJfem9uZXMge1xuICAgICAgICB2YXJjaGFyIGxhdW5jaGVyXG4gICAgICAgIHZhcmNoYXIgem9uZVxuICAgIH1cbiAgICBsYXVuY2hlciB7XG4gICAgICAgIHZhcmNoYXIgbmFtZVxuICAgIH1cbiAgICB6b25lIHtcbiAgICAgICAgdmFyY2hhciBzaG9ydF9uYW1lXG4gICAgICAgIGludCB6b25laWRudW1iZXJcbiAgICAgICAgdGlueWludHVuc2lnbmVkIHZlcnNpb25cbiAgICAgICAgdmFyY2hhciBjb250ZW50X2ZsYWdzXG4gICAgICAgIHZhcmNoYXIgY29udGVudF9mbGFnc19kaXNhYmxlZFxuICAgIH1cbiAgICBsYXVuY2hlcl96b25lcyB8fC0tb3sgbGF1bmNoZXIgOiBPbmUtdG8tT25lXG4gICAgbGF1bmNoZXJfem9uZXMgfHwtLW97IHpvbmUgOiBPbmUtdG8tT25lXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In0sInVwZGF0ZUVkaXRvciI6dHJ1ZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOnRydWV9){target=diagrams}


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

