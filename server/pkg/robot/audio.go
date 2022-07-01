package robot

func (r *robot) PlaySound(filename string) error {
	return r.audio.Play(filename)
}
