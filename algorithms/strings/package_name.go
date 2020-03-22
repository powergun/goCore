package strings


type PackageName struct {
	baseName string
	versionString VersionString
}


func (pn PackageName) IsValid() bool {
	if pn.versionString.IsValid() && len(pn.baseName) > 0 {
		return true
	}
	return false
}


func (pn PackageName) BaseName() string {
	return pn.baseName
}


func (pn PackageName) VersionString() VersionString {
	return pn.versionString
}


func (pn PackageName) ToStr() string {
	if pn.IsValid() {
		return pn.baseName + "." + pn.versionString.ToStr()
	}
	return ""
}
