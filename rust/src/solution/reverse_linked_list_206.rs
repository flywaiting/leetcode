use super::config::ListNode;
impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut last = None;
        let mut cur = head;
        while let Some(mut node) = cur{
            cur = node.next.take();
            node.next = last;
            last = Some(node);
        }
        last
    }
}

pub struct Solution;
