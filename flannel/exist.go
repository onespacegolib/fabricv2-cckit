package flannel

func (f flannel) Exist(id string) (bool, Flannel) {
	state, _ := f.context.GetStateByKey(f.schema.ObjectType, id)
	if state == nil {
		return false, f
	}
	return true, f
}
