package render

import (
	"html/template"
	"path/filepath"
)

// Load_Template_From_File() looks for the Tempalte at the Position and then returns this one.
// TODO This might also load a template from Cache instead of the filestruct. NO instead create a new fucntion arround this, in order to load this from Cache!
// TODO Implement Csachre.
// TODO missleading name
func Load_Template_From_File(paths Render_Filepaths) (template *template.Template) {
	ts, err := template.New(filepath.Base(paths.Filepath_Html)).ParseFiles(paths.Filepath_Html)
	if err != nil {
		panic(err)
	}

	if Check_If_Layout_Exists(paths.Filepath_Layout) {
		ts, err = ts.ParseGlob(paths.Filepath_Layout)
		if err != nil {
			panic(err)
		}
	}
	return ts
}

// Check_If_Layout_Exists() returns true if there is at least one Layout. Otherwise returns false.
func Check_If_Layout_Exists(Layout_Path string) bool {
	if Layout_Path == "" {
		return false
	}
	layouts, err := filepath.Glob(Layout_Path)
	if err != nil {
		panic(err)
	}
	return len(layouts) > 0
}
