package robot

func (r *robot) Forward() {
	r.mover.Forward()
}

func (r *robot) Backward() {
	r.mover.Backward()
}

func (r *robot) Left() {
	r.mover.Left()
}

func (r *robot) Right() {
	r.mover.Right()
}
