package osvid

type OSVid struct {
	VPath        string
	VBasePath    string
	LastModified string
}

func NewOSVid() (*OSVid, error) {
	payload := new(OSVid)

	return payload, nil
}
