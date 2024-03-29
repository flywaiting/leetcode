// 将重叠的范围当作一个对象 然后就是个简单的排列组合
impl Solution {
    pub fn count_ways(ranges: Vec<Vec<i32>>) -> i32 {
        let mut ranges = ranges;
        // ranges.sort_by(|a, b| a[0].cmp(&b[0]));
        // let mut ans: i64 = 1;
        // let mut limit = -1;
        // for p in ranges {
        //     if limit < p[0] {
        //         ans = ans * 2 % 1_000_000_007;
        //     }
        //     limit = p[1].max(limit)
        // }
        // ans as i32

        ranges.sort_unstable_by(|a, b| a[0].cmp(&b[0]));
        ranges
            .iter()
            .fold((1, -1), |(ans, limit), e| {
                (
                    if e[0] > limit {
                        ans * 2 % 1_000_000_007
                    } else {
                        ans
                    },
                    limit.max(e[1]),
                )
            })
            .0
    }
}

pub struct Solution;

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test_2580() {
        assert_eq!(
            Solution::count_ways(vec![[6, 10].to_vec(), [5, 15].to_vec()]),
            2
        );
        assert_eq!(
            Solution::count_ways(vec![
                [1, 3].to_vec(),
                [10, 20].to_vec(),
                [2, 5].to_vec(),
                [4, 8].to_vec()
            ]),
            4
        );
    }
}
