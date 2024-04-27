pub struct Solution;

impl Solution {
    pub fn find_column_width(grid: Vec<Vec<i32>>) -> Vec<i32> {
        let mut res: Vec<i32> = vec![1; grid[0].len()];
        for list in &grid {
            for (i, v) in list.iter().enumerate() {
                res[i] = res[i].max(v.to_string().len() as i32);
            }
        }
        res
    }
}
