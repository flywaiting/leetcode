use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map = HashMap::new();
        // 关于枚举的引用问题
        for (i, &v) in nums.iter().enumerate() {
            // 关于 hashMap 的所有权问题
            if let Some(&j) = map.get(&(target - v)) {
                return vec![j as i32, i as i32];
            }
            map.insert(v, i);
        }
        unreachable!()
    }
}

pub struct Solution;

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test_1() {
        assert_eq!(Solution::two_sum(vec![2, 7, 11, 15], 9), vec![0, 1]);
        assert_eq!(Solution::two_sum(vec![3, 2, 4], 6), vec![1, 2]);
        assert_eq!(Solution::two_sum(vec![3, 3], 6), vec![0, 1]);
    }
}
