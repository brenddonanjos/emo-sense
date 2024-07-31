package main

import (
	"fmt"

	"github.com/brenddonanjos/emo-sense/microsservices/data-collect/pkg/collectors"
)

func main() {
	instagramCollector := collectors.NewInstagramCollector("520960923687994", "https://emo-sense-f9b3277725a5.herokuapp.com/")
	err := instagramCollector.Collect()
	if err != nil {
		fmt.Println(err)
	}
}
