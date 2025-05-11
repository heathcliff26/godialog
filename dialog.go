package godialog

type DialogCallback func(string, error)

type FallbackDialog interface {
	// Shows the open file dialog and calls the callback asynchronously.
	Open(title string, initialDirectory string, filters FileFilters, cb DialogCallback)
	// Shows the save file dialog and calls the callback asynchronously.
	Save(title string, initialDirectory string, filters FileFilters, cb DialogCallback)
}

// OS native file dialog. Allows to define a fallback implementation in case it does not work.
// File dialogs are always opened asynchronously.
type FileDialog struct {
	// The directory that the file dialog should open in.
	InitialDirectory string
	filters          FileFilters
	fallback         FallbackDialog
}

// Return the current file filters.
// Returns nil if no filters are set.
func (fd *FileDialog) Filters() FileFilters {
	return fd.filters
}

// Add a new filter to the list of filters.
func (fd *FileDialog) AddFilter(filter FileFilter) {
	fd.filters = append(fd.filters, filter)
}

// Set the file filters.
func (fd *FileDialog) SetFilters(filters FileFilters) {
	fd.filters = filters
}

// The current fallback dialog.
// Returns nil if no fallback is set.
func (fd *FileDialog) Fallback() FallbackDialog {
	return fd.fallback
}

// Set the fallback dialog in case the native dialog does not work.
func (fd *FileDialog) SetFallback(fallback FallbackDialog) {
	fd.fallback = fallback
}
