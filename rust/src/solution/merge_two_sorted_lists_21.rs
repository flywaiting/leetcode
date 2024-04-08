use super::config::ListNode;

impl Solution {
    pub fn merge_two_lists(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (list1, list2) {
            (Some(mut a), Some(mut b)) => {
                if a.val < b.val {
                    a.next = Self::merge_two_lists(a.next, Some(b));
                    Some(a)
                } else {
                    b.next = Self::merge_two_lists(Some(a), b.next);
                    Some(b)
                }
            }
            (a, b) => {
                if a.is_some() {
                    a
                } else {
                    b
                }
            }
        }
    }
}

pub struct Solution;

// 总结
// 这道简单的题目 被所有权拦了两天.
// 在 `match` 语句中 输入参数的所有权已经发生转移 之前一直在这里徘徊 甚至看了他人的解题 也是一知半解
// 静下心来查看报错之后 开始结合相关内容理解 渐渐发现问题所在 或者说 `rust` 与其他语言的显著区别
// 以后类似的题目要尝试用迭代方式解题 就编程而言 递归太友好了
