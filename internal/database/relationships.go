package database

import (
	"github.com/EQEmu/eqemu-docs-v2/config"
	"log"
	"strings"
)

type RelationshipService struct {
}

func NewRelationshipService() *RelationshipService {
	return &RelationshipService{}
}

// TableRelationships - represents tables from Relationships
type TableRelationships struct {
	Name          string
	Keys          []string
	Relationships []Relationship
}

// Relationship single key Relationships
type Relationship struct {
	RelationshipType string
	LocalTable       string
	LocalKey         string
	RemoteTable      string
	RemoteKey        string
}

// GetTables returns a slice of table objects
// @example
//   database.TableRelationships{
//    Name: "task_activities",
//    Keys: []string{
//      "delivertonpc",
//      "goalid",
//    },
//    Relationships: []database.Relationship{
//      database.Relationship{
//        RelationshipType: "1-1",
//        LocalTable:       "task_activities",
//        LocalKey:         "delivertonpc",
//        RemoteTable:      "npc_types",
//        RemoteKey:        "id",
//      },
//      database.Relationship{
//        RelationshipType: "1-*",
//        LocalTable:       "task_activities",
//        LocalKey:         "goalid",
//        RemoteTable:      "goallists",
//        RemoteKey:        "listid",
//      },
//    },
//  },
//
func (r *RelationshipService) GetTables() []TableRelationships {
	cfg, err := config.GetDbRelationShipsConfig()
	if err != nil {
		log.Fatal(err)
	}

	tables := []TableRelationships{}
	for table, relations := range cfg {

		relationships := []Relationship{}
		keys := []string{}
		for _, relationship := range relations {
			//fmt.Printf("[%v] [%v]\n", table, relationship)

			// example notation
			// [relationship-type] [local-key]->[remote-table]:[remote-key]
			// 1-* lootdrop_id->lootdrop_entries:lootdrop_id
			spaceSplit := strings.Split(relationship, " ")
			if len(spaceSplit) > 0 {
				relationshipType := strings.TrimSpace(spaceSplit[0])
				localRemoteSplit := strings.Split(spaceSplit[1], "->")
				if len(localRemoteSplit) > 0 {
					localKey := strings.TrimSpace(localRemoteSplit[0])
					remoteTableKeySplit := strings.Split(localRemoteSplit[1], ":")

					if len(remoteTableKeySplit) > 0 {
						remoteTable := strings.TrimSpace(remoteTableKeySplit[0])
						remoteKey := strings.TrimSpace(remoteTableKeySplit[1])

						//fmt.Printf("type [%v] local key [%v] remote table [%v] remote key [%v]\n",
						//	RelationshipType,
						//	LocalKey,
						//	RemoteTable,
						//	RemoteKey,
						//)

						// only add key if not already added
						if !containsStringSlice(keys, localKey) {
							keys = append(keys, localKey)
						}

						relationships = append(relationships, Relationship{
							RelationshipType: relationshipType,
							LocalTable:       table,
							LocalKey:         localKey,
							RemoteTable:      remoteTable,
							RemoteKey:        remoteKey,
						})
					}
				}
			}
		}

		tables = append(tables, TableRelationships{
			Name:          table,
			Keys:          keys,
			Relationships: relationships,
		})
	}

	return tables
}
