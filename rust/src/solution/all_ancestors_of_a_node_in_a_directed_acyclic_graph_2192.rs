impl Solution {
    pub fn get_ancestors(n: i32, edges: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut g = vec![vec![]; n as usize];
        for e in edges.iter() {
            g[e[0] as usize].push(e[1] as usize);
        }

        let mut ans = vec![vec![]; n as usize];

        fn dfs(from: i32, i: usize, g: &Vec<Vec<usize>>, ans: &mut Vec<Vec<i32>>) {
            for &to in &g[i] {
                if ans[to].contains(&from) {
                    continue;
                }
                ans[to].push(from);
                dfs(from, to, g, ans);
            }
        }
        for i in 0..n {
            dfs(i, i as usize, &g, &mut ans);
        }
        ans
    }
}

pub struct Solution;
