pub struct Solution;

impl Solution {
    pub fn distribute_candies(candies: i32, num_people: i32) -> Vec<i32> {
        let mut c = 0;
        while (c + 1) * (c + 2) / 2 < candies {
            c += 1;
        }
        let mid = candies - c * (c + 1) / 2;
        let m = c % num_people;
        let n = c / num_people;
        let mut res = vec![0; num_people as usize];
        for i in 0..num_people {
            res[i as usize] = (i + 1) * n + n * (n - 1) / 2 * num_people;
            if i < m {
                res[i as usize] += i + 1 + n * num_people;
            }
        }
        res[m as usize] += mid;
        res
    }
}

// 总结
// 不要特殊处理 最后一个数据统一处理 一定要统一处理！！！
// 特殊处理 带来一系列情况判定 将问题复杂化了
