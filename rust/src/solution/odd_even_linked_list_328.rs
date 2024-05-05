use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn odd_even_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut dummy = Box::new(ListNode::new(0));
        let mut tail = Box::new(ListNode::new(0));

        let mut i = 1;
        let mut head = head;

        let mut odd = &mut dummy;
        let mut even = &mut tail;
        while let Some(mut node) = head {
            head = node.next.take();
            if i % 2 == 1 {
                odd = odd.next.insert(node);
            } else {
                even = even.next.insert(node);
            }

            i += 1;
        }

        odd.next = tail.next;

        dummy.next
    }
}
