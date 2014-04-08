package vinamax

//Calculates the torque working on the magnetisation of a particle
//using the Landau Lifshitz equation
func (p *particle) tau(temp vector) vector {
	mxB := p.m.cross(p.b_eff(temp))
	amxmxB := p.m.cross(mxB).times(Alpha)
	mxB = mxB.add(amxmxB)
	return mxB.times(-gamma0 / (1 + (Alpha * Alpha)))
}
