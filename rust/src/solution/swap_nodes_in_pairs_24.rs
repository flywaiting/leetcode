use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = Box::new(ListNode { val: 0, next: head });
        let mut cur = head.as_mut();
        while let Some(mut node) = cur.next.take() {
            if let Some(mut next) = node.next.take() {
                node.next = next.next;
                next.next = Some(node);
                cur.next = Some(next);
                cur = cur.next.as_mut().unwrap().next.as_mut().unwrap();
            } else {
                cur.next = Some(node);
                break;
            }
        }
        head.next
    }
}
