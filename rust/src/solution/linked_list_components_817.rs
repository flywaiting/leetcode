use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn num_components(head: Option<Box<ListNode>>, nums: Vec<i32>) -> i32 {
        // let mut res = 0;
        // let mut flag = false;
        // let mut p = head.as_ref();
        // while let Some(node) = p {
        //     if nums.contains(&node.val) {
        //         flag = true;
        //     } else {
        //         if flag {
        //             res += 1;
        //         }
        //         flag = false;
        //     }

        //     p = node.next.as_ref();
        // }
        // if flag {
        //     res += 1;
        // }
        // res

        let mut cnt = 0;
        let mut flag = false;
        let nums = nums.iter().collect::<std::collections::HashSet<_>>();
        let mut p = head.as_ref();
        while let Some(node) = p {
            p = node.next.as_ref();
            if nums.contains(&node.val) {
                if !flag {
                    flag = true;
                    cnt += 1;
                }
            } else {
                flag = false;
            }
        }
        cnt
    }
}

// note
// 第一次题意理解有误 成了记录最大长度的逻辑
// 第二次 正常解题 数组包含元素的判断方法性能好差 有点出乎意料...
// 学习提高 用 `HashSet` 可以进行高效判定
