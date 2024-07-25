package valueobject

type ProjectID string

func NewProjectID(ID string) ProjectID {
    return ProjectID(ID)
}


func (p *ProjectID) String() string {
    return string(*p)
}

//validate that every project id is not empty and starts with "proj_" and 27 characters long
func (p *ProjectID) Validate() bool {
    if len(*p) != 31 {
        return false
    }
    if (*p)[0:5] != "proj_" {
        return false
    }
    return true
}
