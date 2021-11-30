package flannel

import (
	"git.onespace.co.th/osgolib/fabricv2-cckit/router"
)

func Router(r *router.Group) {
	r.Group(`flannel`).
		Query(`console`, Console, nil).
		Query(`document`, Document, nil)
}
