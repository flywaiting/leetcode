use std::cmp::min;

impl Solution {
    pub fn minimum_sum(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut ans = i32::MAX;
        for i in 0..n {
            for j in i + 1..n {
                for k in j + 1..n {
                    if nums[i] < nums[j] && nums[k] < nums[j] {
                        ans = min(ans, nums[i] + nums[j] + nums[k])
                    }
                }
            }
        }
        if ans == i32::MAX {
            -1
        } else {
            ans
        }
    }

    pub fn minimum_sum_2(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut left = vec![0; n];
        for i in 0..n {
            if i == 0 {
                left[i] = nums[i];
            } else {
                left[i] = nums[i].min(left[i - 1]);
            }
        }

        let mut right = nums[n - 1];
        let mut ans = i32::MAX;
        for i in (0..n - 1).rev() {
            if nums[i] > left[i] && nums[i] > right {
                ans = ans.min(nums[i] + left[i] + right);
            }
            right = right.min(nums[i]);
        }

        if ans == i32::MAX {
            -1
        } else {
            ans
        }
    }
}

pub struct Solution;

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test_2908() {
        assert_eq!(Solution::minimum_sum(vec![8, 6, 1, 5, 3]), 9);
        assert_eq!(Solution::minimum_sum(vec![5, 4, 8, 7, 10, 2]), 13);
        assert_eq!(Solution::minimum_sum(vec![6, 5, 4, 3, 4, 5]), -1);
    }
    #[test]
    fn test_2908_2() {
        assert_eq!(Solution::minimum_sum_2(vec![8, 6, 1, 5, 3]), 9);
        assert_eq!(Solution::minimum_sum_2(vec![5, 4, 8, 7, 10, 2]), 13);
        assert_eq!(Solution::minimum_sum_2(vec![6, 5, 4, 3, 4, 5]), -1);
    }
}
