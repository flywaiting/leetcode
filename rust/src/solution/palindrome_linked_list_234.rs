use super::config::ListNode;

impl Solution {
    pub fn is_palindrome(head: Option<Box<ListNode>>) -> bool {
        let mut head = head;
        if head.is_none() {
            return true;
        }

        let mut nums = vec![];
        while let Some(node) = head {
            nums.push(node.val);
            head = node.next;
        }

        let (mut left, mut right) = (0, nums.len() - 1);
        while left <= right {
            if nums[left] != nums[right] {
                return false;
            } else {
                left += 1;
                right -= 1;
            }
        }
        true
    }
}

pub struct Solution;

// 解题思路
// 简单方式：用数组收集数据 然后进行简单的回文比较
// 高级方式：利用快慢指针分出两边链表 反转一边的链表进行比较 之后还原被反转链表
// 虽然但是。。。 没有想到修改链表这一点 也还好吧 思路限制了 当然 也更麻烦了
