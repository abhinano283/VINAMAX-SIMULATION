package vinamax

//Runs the simulation for a certain time
func Run(time float64) {
	testinput()
	for i := range universe.lijst {
		norm(universe.lijst[i].m)
	}
	write(averages(universe.lijst))
	for j := T; T < j+time; {
		if Demag {
			calculatedemag()
		}
		//TODO dit variabel maken tussen euler en heun
		heunstep(universe.lijst)
		write(averages(universe.lijst))
	}
}

//perform a timestep using euler forward method
func eulerstep(Lijst []*Particle) {
	for i := range Lijst {
		Lijst[i].m[0] += Lijst[i].tau()[0] * Dt
		Lijst[i].m[1] += Lijst[i].tau()[1] * Dt
		Lijst[i].m[2] += Lijst[i].tau()[2] * Dt
		Lijst[i].m = norm(Lijst[i].m)

	}
	T += Dt
}

//perform a timestep using heun method
//http://en.wikipedia.org/wiki/Heun_method
func heunstep(Lijst []*Particle) {
	for _, p := range Lijst {

		tau1 := p.tau()

		//tau van t+1, positie nadat met tau1 al is doorgevoerd
		p.m[0] += tau1[0] * Dt
		p.m[1] += tau1[1] * Dt
		p.m[2] += tau1[2] * Dt

		tau2 := p.tau()

		p.m[0] += ((-tau1[0] + tau2[0]) * 0.5 * Dt)
		p.m[1] += ((-tau1[1] + tau2[1]) * 0.5 * Dt)
		p.m[2] += ((-tau1[2] + tau2[2]) * 0.5 * Dt)

		p.m = norm(p.m)

	}
	T += Dt
}
