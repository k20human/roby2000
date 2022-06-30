package robot

import "github.com/pkg/errors"

var ErrMustBackward = errors.New("must backward")

func (r *robot) MoveForward() error {
	d, err := r.distance.Dist()

	if d > 10 && err == nil {
		r.mover.Forward()
		return nil
	} else {
		return ErrMustBackward
	}
}

func (r *robot) MoveBackward() {
	r.mover.Backward()
}

func (r *robot) MoveLeft() {
	r.mover.Left()
}

func (r *robot) MoveRight() {
	r.mover.Right()
}
