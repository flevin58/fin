package cmd

const (
	OkGliph  string = "👍"
	ErrGliph string = "❌"
)

// Variables that will hold flag data
var (
	flagAdd     bool
	flagDebug   bool
	flagRemove  bool
	flagSync    bool
	flagExtract string
	flagList    string
	flagOutput  string
)
