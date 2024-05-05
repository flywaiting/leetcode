use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        if s.is_empty() {
            return 0;
        }

        let mut ans = 0;
        let mut map = HashMap::new();
        let mut left = 0;
        for (right, e) in s.as_bytes().iter().enumerate() {
            // 窗口外的数据无效
            if map.contains_key(e) && map.get(e).unwrap() >= &left {
                ans = ans.max(right - left);
                left = *map.get(e).unwrap() + 1;
            }
            map.insert(e, right);
        }
        ans = ans.max(s.len() - left); // 没有每次更新 最后需要在确认一次
        ans as i32
    }
}
