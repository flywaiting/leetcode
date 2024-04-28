use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn partition(head: Option<Box<ListNode>>, x: i32) -> Option<Box<ListNode>> {
        let mut res = Box::new(ListNode::new(0));
        let mut tail = None;
        let mut left = &mut res;
        let mut right = &mut tail;

        let mut head = head;
        while let Some(mut node) = head {
            head = node.next.take();
            if node.val < x {
                left = left.next.insert(node);
            } else {
                right = &mut right.insert(node).next;
            }
        }

        left.next = tail;
        res.next
    }
}

// 解题思路
// 居然一把过了 表示有点怀疑...
// 好吧 `rust`还是牛啊 能过测试用例 几乎就能通过
