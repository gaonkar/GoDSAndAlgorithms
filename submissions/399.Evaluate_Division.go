/*
Equations are given in the format A / B = k, where A and B are variables represented as strings, and k is a real
number (floating point number). Given some queries, return the answers. If the answer does not exist, return -1.0.

Example:
Given a / b = 2.0, b / c = 3.0.
queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ? .
return [6.0, 0.5, -1.0, 1.0, -1.0 ].

The input is: vector<pair<string, string>> equations, vector<double>& values, vector<pair<string, string>> queries ,
where equations.size() == values.size(), and the values are positive. This represents the equations. Return vector<double>.

According to the example above:

equations = [ ["a", "b"], ["b", "c"] ],
values = [2.0, 3.0],
queries = [ ["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"] ].


The input is always valid. You may assume that evaluating the queries will result in no division by zero and there is
no contradiction.
*/


func Query(N map[string] bool,A map[string] []string,V map[string] float64, i, j string, v map[string] bool) (r float64) {
    if !N[i] || !N[j] {return -1.0}
    //fmt.Println("Q", i, j)
    if V[i + "_" + j] > 0 {
        return V[i + "_" + j]
    }
    if V[j + "_" + i] > 0 {
        return V[j + "_" + i]
    }
    r = -1
    v[i]=true
    for _,x := range(A[i]) {
        if v[x] {continue}
        //fmt.Println(i,x,j,V[i + "_" + x])
        val :=  Query(N,A,V, x, j, v)
        if val > 0 {
            r = V[i + "_" + x] * val
        }
        if r > 0 {break}
    }
    v[i]=false
    return r
}

func calcEquation(equations [][]string, values []float64, queries [][]string) (r []float64) {
    N := make(map[string] bool)
    A := make(map[string] []string)
    V := make(map[string] float64)
    for i,x := range(equations) {
        N[x[0]] = true
        N[x[1]] = true
        A[x[0]] = append(A[x[0]], x[1])
        A[x[1]] = append(A[x[1]], x[0])
        V[x[0] + "_" + x[1]] = values[i]
        V[x[1] + "_" + x[0]] = 1/values[i]
    }
    //fmt.Println(N, A, V)
    for _,q := range(queries) {
        r = append(r, Query(N,A,V, q[0], q[1], make(map[string] bool)))
    }
    return r
}
