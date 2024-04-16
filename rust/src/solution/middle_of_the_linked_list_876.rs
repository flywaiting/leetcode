use super::config::ListNode;

impl Solution {
    pub fn middle_node(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut slow = &head;
        let mut fast = &head.as_ref().unwrap().next;
        while let Some(node) = fast {
            if node.next.is_none() {
                break;
            } else {
                slow = &slow.as_ref().unwrap().next;
                fast = &node.next.as_ref().unwrap().next;
            }
        }
        if fast.is_none() {
            slow.clone()
        } else {
            slow.as_ref()?.next.clone()
        }
    }
}
pub struct Solution;
