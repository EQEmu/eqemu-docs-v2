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

// TableRelationships - represents tables from relationships
type TableRelationships struct {
	name          string
	keys          []string
	relationships []Relationship
}

// Relationship single key relationships
type Relationship struct {
	relationshipType string
	localTable       string
	localKey         string
	remoteTable      string
	remoteKey        string
}

// GetTables returns a slice of table objects
// @example
//   database.TableRelationships{
//    name: "task_activities",
//    keys: []string{
//      "delivertonpc",
//      "goalid",
//    },
//    relationships: []database.Relationship{
//      database.Relationship{
//        relationshipType: "1-1",
//        localTable:       "task_activities",
//        localKey:         "delivertonpc",
//        remoteTable:      "npc_types",
//        remoteKey:        "id",
//      },
//      database.Relationship{
//        relationshipType: "1-*",
//        localTable:       "task_activities",
//        localKey:         "goalid",
//        remoteTable:      "goallists",
//        remoteKey:        "listid",
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
						//	relationshipType,
						//	localKey,
						//	remoteTable,
						//	remoteKey,
						//)

						// only add key if not already added
						if !containsStringSlice(keys, localKey) {
							keys = append(keys, localKey)
						}

						relationships = append(relationships, Relationship{
							relationshipType: relationshipType,
							localTable:       table,
							localKey:         localKey,
							remoteTable:      remoteTable,
							remoteKey:        remoteKey,
						})
					}
				}
			}
		}

		tables = append(tables, TableRelationships{
			name:          table,
			keys:          keys,
			relationships: relationships,
		})
	}

	return tables
}
