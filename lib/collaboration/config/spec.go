package config

type Spec interface{}

type TransformationGroupSpec struct {
	CollaboratorRef string               `yaml:"collaborator"`
	Transformations []TransformationSpec `yaml:"transformations"`
}

// For now only supporting Count Query.
// TODO - Make this Spec such that any type of transformation can be parsed using this spec.
type TransformationSpec struct {
	Name                     string     `yaml:"name"`
	Type                     string     `yaml:"type"`
	UniqueId                 string     `yaml:"unique_id,omitempty"`
	AppLocation              string     `yaml:"app_location"`
	From                     []FromSpec `yaml:"from"`
	JoinKey                  string     `yaml:"join_key,omitempty"`
	NoiseParams              []string   `yaml:"noise_parameters,omitempty"`
	Template                 string     `yaml:"template,omitempty"`
	DestinationOwnersAllowed []string   `yaml:"destination_owners_allowed"`
	// TODO - do we need access control at destination level granularity?
	DestinationsAllowed []string `yaml:"destination_allowed,omitempty"`
}

type SourceGroupSpec struct {
	CollaboratorRef string       `yaml:"collaborator"`
	Sources         []SourceSpec `yaml:"sources"`
}

type SourceSpec struct {
	Name        string       `yaml:"name"`
	CSVLocation string       `yaml:"csv_location"`
	Description string       `yaml:"description"`
	Columns     []ColumnSpec `yaml:"columns"`
	// TODO- Do we need to add addressRef here?
	TransformationOwnersAllowed []string                       `yaml:"transformation_owners_allowed"`
	DestinationsAllowed         []SourceDestinationAllowedSpec `yaml:"destinations_allowed"`
}

type ColumnSpec struct {
	Name              string   `yaml:"name"`
	Type              string   `yaml:"type"`
	MaskingType       string   `yaml:"masking_type"`
	Selectable        bool     `yaml:"selectable,omitempty"`
	AggregatesAllowed []string `yaml:"aggregates_allowed,omitempty"`
	JoinKey           bool     `yaml:"join_key,omitempty"`
}

type DestinationGroupSpec struct {
	CollaboratorRef string            `yaml:"collaborator"`
	Destinations    []DestinationSpec `yaml:"destinations"`
}

type DestinationSpec struct {
	Name string `yaml:"name"`
	Ref  string `yaml:"ref"`
}

type SourceDestinationAllowedSpec struct {
	Ref         string              `yaml:"ref"`
	NoiseParams []map[string]string `yaml:"noise_parameters"`
}

type FromSpec struct {
	Name        string `yaml:"name"`
	Ref         string `yaml:"ref"`
	LocationTag string `yaml:"location_tag"`
}
