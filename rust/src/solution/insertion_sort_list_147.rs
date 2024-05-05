use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn insertion_sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = head;
        let mut dummy = Box::new(ListNode::new(0));

        while let Some(mut node) = head {
            head = node.next.take();

            let mut cur = &mut dummy;
            while cur.next.is_some() && cur.next.as_ref().unwrap().val < node.val {
                cur = cur.next.as_mut().unwrap();
            }
            node.next = cur.next.take();
            cur.next = Some(node)
        }

        dummy.next
    }
}
