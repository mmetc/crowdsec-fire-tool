package main

import (
	"os"
	"path"

	"github.com/crowdsecurity/crowdsec/pkg/cticlient"
)

func intPtr(i int) *int {
	return &i
}

func main() {
	client := cticlient.NewCrowdsecCTIClient(cticlient.WithAPIKey(os.Getenv("CTI_API_KEY")))
	paginator := cticlient.NewFirePaginator(client, cticlient.FireParams{
		Limit: intPtr(1000),
	})

	outputDir := os.Getenv("OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "./"
	}

	file, err := os.Create(path.Join(outputDir, "fire.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {
		items, err := paginator.Next()
		if err != nil {
			panic(err)
		}
		if items == nil {
			break
		}

		for _, item := range items {
			file.WriteString(item.Ip + "\n")
		}
	}
}
