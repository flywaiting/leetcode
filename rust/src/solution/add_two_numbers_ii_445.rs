use super::config::ListNode;

pub struct Solution;

impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut l1 = Self::reverse(l1);
        let mut l2 = Self::reverse(l2);

        let mut sum = None;
        let mut carrier = 0;
        while l1.is_some() || l2.is_some() {
            match (l1, l2) {
                (Some(mut node1), Some(mut node2)) => {
                    l1 = node1.next.take();
                    l2 = node2.next.take();
                    carrier += node1.val + node2.val;
                    node1.val = carrier % 10;
                    carrier /= 10;
                    node1.next = sum;
                    sum = Some(node1);
                }
                (Some(mut node), None) | (None, Some(mut node)) => {
                    l1 = node.next.take();
                    l2 = None;
                    carrier += node.val;
                    node.val = carrier % 10;
                    carrier /= 10;
                    node.next = sum;
                    sum = Some(node);
                }
                _ => break,
            }
        }

        if carrier > 0 {
            let mut node = Box::new(ListNode::new(carrier));
            node.next = sum;
            sum = Some(node);
        }
        sum
    }

    fn reverse(mut root: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut cur = None;
        while let Some(mut node) = root {
            root = node.next.take();
            node.next = cur;
            cur = Some(node);
        }
        cur
    }
}
