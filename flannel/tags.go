package flannel

const (
	TAG = `flannel`
)

func (f flannel) getValueOfTag(tag string) (name string, value string, ok bool) {
	s := f.modelStruct
	for _, v := range s.Names() {
		filed := s.Field(v)
		if filed.Tag(TAG) == tag {
			if value, ok = filed.Value().(string); ok {
				return v, value, true
			} else {
				return v, "", false
			}

		}
	}
	return "", "", false
}
