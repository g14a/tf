package google

import (
	"github.com/manifoldco/promptui"
	"strings"
)

func GetRegions() []string {
	return []string{"asia-east1-a", "asia-east1-b",
		"asia-east1-c", "asia-east2-a", "asia-east2-b", "asia-east2-c",
		"asia-northeast1-a", "asia-northeast1-b", "asia-northeast1-c", "asia-northeast2-a",
		"asia-northeast2-b", "asia-northeast2-c", "asia-northeast3-a",
		"asia-northeast3-b", "asia-northeast3-c", "asia-south1-a", "asia-south1-b",
		"asia-south1-c", "asia-southeast1-a", "asia-southeast1-b", "asia-southeast1-c",
		"asia-southeast2-a", "asia-southeast2-b", "asia-southeast2-c",
		"australia-southeast1-a", "australia-southeast1-b", "australia-southeast1-c", "europe-north1-a",
		"europe-north1-b", "europe-north1-c", "europe-west1-b", "europe-west1-c",
		"europe-west1-d", "europe-west2-a", "europe-west2-b", "europe-west2-c",
		"europe-west3-a", "europe-west3-b", "europe-west3-c", "europe-west4-a", "europe-west4-b",
		"europe-west4-c", "europe-west6-a", "europe-west6-b", "europe-west6-c", "northamerica-northeast1-a",
		"northamerica-northeast1-b", "northamerica-northeast1-c", "southamerica-east1-a",
		"southamerica-east1-b", "southamerica-east1-c", "us-central1-a", "us-central1-b", "us-central1-c",
		"us-central1-f", "us-east1-b", "us-east1-c", "us-east1-d", "us-east4-a",
		"us-east4-b", "us-east4-c", "us-west1-a", "us-west1-b", "us-west1-c",
		"us-west2-a", "us-west2-b", "us-west2-c", "us-west3-a", "us-west3-b", "us-west3-c",
		"us-west4-a", "us-west4-b", "us-west4-c",
	}
}

func GCPRegionPrompt() *promptui.Select {
	return &promptui.Select{
		Label:             "GCP regions",
		Size:              20,
		Items:             GetRegions(),
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			provider := GetRegions()[index]
			name := strings.Replace(strings.ToLower(provider), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		},
	}
}
