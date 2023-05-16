package cmdutil

const (
	DefaultConfigFileName = ".reman"
	VersionKey            = "$version" // Key for the version info in config file, current version is defined in `cmdutil.Version`

	AuthorNameKey     = "author.name" // Name of the author, used for writing filenames
	AuthorNameLimit   = 19            // Character limit for the name of the author
	DefaultAuthorName = "Reman"       // Default value for author name, defaults are set in config at init
)
