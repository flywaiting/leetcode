use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = head;
        let mut list = Vec::new();

        let mut p = head.as_ref();
        while let Some(node) = p {
            list.push(node.val);
            p = node.next.as_ref();
        }

        list.sort_unstable();
        let mut i = 0;
        let mut cur = head.as_mut();
        while let Some(mut node) = cur {
            cur = node.next.as_mut();
            node.val = list[i];
            i += 1;
        }
        head
    }
}
