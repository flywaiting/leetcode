impl Solution {
    pub fn maximum_count(nums: Vec<i32>) -> i32 {
        let mut n = 0;
        let mut z = 0;
        let size = nums.len();
        for i in nums {
            if i < 0 {
                n += 1;
            } else if i == 0 {
                z += 1;
            } else {
                break;
            }
        }
        n.max(size - z - n) as i32
    }
}

pub struct Solution;

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test_2529() {
        assert_eq!(Solution::maximum_count(vec![-2, -1, -1, 1, 2, 3]), 3);
        assert_eq!(Solution::maximum_count(vec![-3, -2, -1, 0, 0, 1, 2]), 3);
        assert_eq!(Solution::maximum_count(vec![5, 20, 66, 1314]), 4);
    }
}
