package main

import (
	"CustomHashMap/internal/hash_table"
	"CustomHashMap/internal/hash_table_group"
	"fmt"
)

func putSomeValuesToHashTable(ht hash_table.HashTable) {
	ht.Put("someKey1", "someValue1")
	ht.Put("someKey2", "someValue2")
	ht.Put("someKey3", "someValue3")
}

func main() {
	ht1 := hash_table.NewHashTableRMBasic(10)
	putSomeValuesToHashTable(ht1)

	ht2 := hash_table.NewHashTableRMBasic(10)
	putSomeValuesToHashTable(ht2)

	ht3 := hash_table.NewHashTableRMBasic(10)
	putSomeValuesToHashTable(ht3)

	htg1 := hash_table_group.NewHashTableGroupRMBasic(10)
	htg1.PutTable("ht1", ht1)

	htg2 := hash_table_group.NewHashTableGroupRMBasic(10)
	htg2.PutTable("ht2", ht2)

	htFromHtg1, _ := htg1.GetTable("ht1")
	fmt.Println(htFromHtg1)

	htFromHtg2, _ := htg2.GetTable("ht2")
	fmt.Println(htFromHtg2)
}
