using JuMP
using GLPK
using DelimitedFiles
import LinearAlgebra
G = [
	 [0 12 12 0 12 0 0 0 1 0 0 0 0 0 0 0]
	 [0 0 0 6 0 6 0 0 0 5 0 0 0 0 0 0]
	 [0 0 0 1 0 0 6 0 0 0 7 0 0 0 0 0]
	 [0 0 0 0 0 0 0 4 0 0 0 8 0 0 0 0]
	 [0 0 0 0 0 4 3 0 0 0 0 0 2 0 0 0]
	 [0 0 0 0 0 0 0 4 0 0 0 0 0 2 0 0]
	 [0 0 0 0 0 0 0 8 0 0 0 0 0 0 8 0]
	 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 3]
	 [0 0 0 0 0 0 0 0 0 4 2 0 2 0 0 0]
	 [0 0 0 0 0 0 0 0 0 0 0 2 0 6 0 0]
	 [0 0 0 0 0 0 0 0 0 0 0 1 0 0 1 0]
	 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2]
	 [0 0 0 0 0 0 0 0 0 0 0 0 0 7 6 0]
	 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 5]
	 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2]
	 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
]
n = size(G)[1]
max_flow = Model(GLPK.Optimizer)
@variable(max_flow, f[1:n,1:n] >= 0)
@constraint(max_flow, [i = 1:n, j = 1:n], f[i, j] <= G[i, j])
@constraint(max_flow, [i = 1:n; i != 1 && i != n], sum(f[i, :]) == sum(f[:, i]))
@objective(max_flow, Max, sum(f[1, :]))
optimize!(max_flow)
objective_value(max_flow)
value.(f)
