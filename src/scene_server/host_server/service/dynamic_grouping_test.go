package service

import (
	"configcenter/src/common/metadata"
	"encoding/json"
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	results := []metadata.DynamicGroupClassification{
		{
			ID:       "1",
			Name:     "1",
			ParentID: "-1",
		},
		{
			ID:       "2",
			Name:     "2",
			ParentID: "1",
		},
		{
			ID:       "3",
			Name:     "3",
			ParentID: "2",
		},
	}

	tree := &ClassificationTree{
		ID:       "-1",
		Name:     "root",
		ParentID: "-1",
		Children: make([]*ClassificationTree, 0),
	}

	for _, item := range results {
		if item.ParentID == "-1" {
			tree.Children = append(tree.Children, &ClassificationTree{
				ID:       item.ID,
				Name:     item.Name,
				ParentID: item.ParentID,
				Children: make([]*ClassificationTree, 0),
			})
		} else {
			buildTree(&item, tree)
		}
	}

	rawData, err := json.Marshal(tree)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", string(rawData))
}
