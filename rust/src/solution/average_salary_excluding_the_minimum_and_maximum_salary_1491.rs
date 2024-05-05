pub struct Solution;

impl Solution {
    pub fn average(salary: Vec<i32>) -> f64 {
        let mut min = i32::MAX;
        let mut max = i32::MIN;
        let mut sum = salary.iter().fold(0, |acc, &e| {
            if e < min {
                min = e;
            }
            if e > max {
                max = e;
            }
            acc + e
        });

        sum -= min + max;
        (sum as f64) / ((salary.len() - 2) as f64)
    }
}
