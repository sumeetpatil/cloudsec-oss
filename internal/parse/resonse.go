/*
Copyright © 2022 Sumeet Patil sumeet.patil@sap.com

*/

package parse

type Response struct {
	Vulns []struct {
		ID                string   `json:"id"`
		Summary           string   `json:"summary"`
		Details           string   `json:"first_name"`
		Aliases           []string `json:"aliases"`
		Modified          string   `json:"modified"`
		Published         string   `json:"published"`
		Database_specific struct {
			Cwe_ids         []string `json:"cwe_ids"`
			Severity        string   `json:"severity"`
			Github_reviewed bool     `json:"github_reviewed"`
		} `json:"database_specific"`
		References []struct {
			RefType         string `json:"type"`
			Url             string `json:"url"`
			Github_reviewed bool   `json:"github_reviewed"`
		} `json:"references"`
		Affected []struct {
			AffectedPackage struct {
				Name      string `json:"name"`
				Ecosystem string `json:"ecosystem"`
				Purl      string `json:"purl"`
			} `json:"package"`
			Versions          []string `json:"versions"`
			Database_specific struct {
				Last_known_affected_version_range string `json:"last_known_affected_version_range"`
				Source                            string `json:"source"`
			} `json:"database_specific"`
			Schema_version string `json:"schema_version"`
		} `json:"affected"`
		Severity []struct {
			Severitytype string `json:"type"`
			Score        string `json:"score"`
		} `json:"severity"`
	} `json:"vulns"`
}
